package api

import (
	"auto_post/app/internal/adapters/port"
	"auto_post/app/internal/domains/vk_machine/structs"
	"auto_post/app/pkg/dbo"
	"auto_post/app/pkg/ddo"
	"auto_post/app/pkg/errs"
	logger "auto_post/app/pkg/log"
	"auto_post/app/pkg/mapping"
	"auto_post/app/pkg/tools"
	status "auto_post/app/pkg/vars/statuses"
	"context"
	"fmt"
	"github.com/nuttech/bell/v2"
	"os"
)

// VKWallPostUseCase --
type VKWallPostUseCase struct {
	log             logger.Logger
	events          *bell.Events
	persister       port.Persister
	extractor       port.Extractor
	vkMachineDomain port.VkMachineDomain
}

// NewVKWallPostUseCase --
func NewVKWallPostUseCase(
	log logger.Logger,
	events *bell.Events,
	persister port.Persister,
	extractor port.Extractor,
	vkMachineDomain port.VkMachineDomain,
) port.VKWallPostUseCase {
	return VKWallPostUseCase{log: log, events: events, persister: persister, extractor: extractor, vkMachineDomain: vkMachineDomain}
}

// Execute _
func (ths VKWallPostUseCase) Execute(ctx context.Context) error {
	msg := fmt.Sprintf
	log := ths.log.WithMethod("usecase VKWallPost")

	// Выгрузка случайной записи из БД
	recordDBO := dbo.RecordDBO{}
	if err := ths.extractor.GetByActiveStatus(&recordDBO); err != nil {
		return extErr(errs.UnknownError, // TODO актуализировать константы ошибок
			msg("failed to get RND entity by active status with error: %s", err.Error()), log)
	}

	log.Debug("recordDBO: %v", recordDBO)

	// Поход в домен - подготовка вызова метода photos.getWallUploadServer
	getWallUploadServerDDO := ths.vkMachineDomain.GetWallUploadServer()

	// вызов метода
	responseGetWallUploadServer := structs.VkGetWallUploadServer{}
	_, err := tools.Request("v", getWallUploadServerDDO.MethodName, getWallUploadServerDDO.Params, &responseGetWallUploadServer)
	if err != nil {
		return extErr(errs.UnknownError,
			msg("getWallUploadServer failed with error: %s", err.Error()), log)
	}

	// Поход в домен - получение пути до файла
	vkMachineDDO := mapping.RecordDbOtoVkMachineDDO(&recordDBO)
	path := ths.vkMachineDomain.GetPath(vkMachineDDO)

	// подготовительный этап в формировании загрузки на сайт
	uploadPhotoWallResponse, err := tools.PhotoWall(responseGetWallUploadServer.Response.UploadURL, path)
	if err != nil {
		return extErr(errs.UnknownError,
			msg("failed PhotoWall with error %s", err.Error()), log)
	}

	reqSaveWallPhotoDDO := ddo.ReqSaveWallPhoto{
		Photo:  uploadPhotoWallResponse.Photo,
		Hash:   uploadPhotoWallResponse.Hash,
		Server: uploadPhotoWallResponse.Server,
	}
	// Поход в домен -
	saveWallPhotoDDO := ths.vkMachineDomain.SaveWallPhoto(&reqSaveWallPhotoDDO)
	// вызов метода
	responseSaveWallPhoto := structs.VkSaveWallPhoto{}
	_, err = tools.Request("v", saveWallPhotoDDO.MethodName, saveWallPhotoDDO.Params, &responseSaveWallPhoto)
	if err != nil {
		return extErr(errs.UnknownError,
			msg("saveWallPhoto failed with error: %s", err.Error()), log)
	}
	if len(responseSaveWallPhoto.Response) <= 0 {
		return extErr(errs.UnknownError,
			msg("method saveWallPhoto failed: Response index = 0; with error %s", err.Error()), log)
	}

	// Поход в домен - публикация поста на стене группы
	//vkMachineDDO = mapping.RecordDbOtoVkMachineDDO(&recordDBO)
	//ths.vkMachineDomain.UploadPhotoToServer(vkMachineDDO)

	// Поход в домен wall.post
	reqPostWallPhoto := ddo.ReqPostWallPhoto{
		URL:     recordDBO.URL,
		AuthURL: recordDBO.AuthURL,
		OwnerID: responseSaveWallPhoto.Response[0].OwnerID,
		ID:      responseSaveWallPhoto.Response[0].ID,
	}
	resPostWallPhoto := ths.vkMachineDomain.PostWallPhoto(&reqPostWallPhoto)

	_, err = tools.Request("v", resPostWallPhoto.MethodName, resPostWallPhoto.Params, nil)
	if err != nil {
		return extErr(errs.UnknownError,
			msg("ResPostWallPhoto failed with error: %s", err.Error()), log)
	}

	getUploadServer := ths.vkMachineDomain.GetUploadServer()
	resGetUploadServer := structs.VkGetWallUploadServer{}
	_, err = tools.Request("v", getUploadServer.MethodName, getUploadServer.Params, &resGetUploadServer)
	if err != nil {
		return extErr(errs.UnknownError,
			msg("GetUploadServer failed with error: %s", err.Error()), log)
	}

	// загрузить картинку по ссылке полученной ранее
	uploadPhoto, err := tools.PhotoGroup(resGetUploadServer.Response.UploadURL, path)
	if err != nil {
		ths.log.Error("uploaded to album failed: " + fmt.Sprint(err))
	}

	// сохранить картинку в альбом
	reqPhotoSave := ddo.ReqPhotosSave{
		Server:    uploadPhoto.Server,
		PhotoList: uploadPhoto.PhotoList,
		Hash:      uploadPhoto.Hash,
		URL:       recordDBO.URL,
		AlbumID:   uploadPhoto.AlbumID,
	}
	photoSave := ths.vkMachineDomain.PhotosSave(reqPhotoSave)
	_, err = tools.Request("v", photoSave.MethodName, photoSave.Params, nil)
	if err != nil {
		return extErr(errs.UnknownError,
			msg(" failed PhotoSave with error: %s", err.Error()), log)
	}

	// обновление статуса записи в БД на "used"
	recordDBO.Status = status.RecordUsedStatus
	if err := ths.persister.UpdateRecordStatus(&recordDBO); err != nil {
		return extErr(errs.UnknownError, // TODO прибраться в логике возврата ошибок
			msg("failed updateRecordStatus to used with error: %s", err.Error()), log)
	}
	// удалить файл
	if err = os.Remove(path); err != nil {
		ths.log.Error("error Remove file: " + fmt.Sprint(err))
	}

	return nil
}
