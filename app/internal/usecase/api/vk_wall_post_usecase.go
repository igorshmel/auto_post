package api

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/igorshmel/lic_auto_post/app/internal/adapters/port"
	"github.com/igorshmel/lic_auto_post/app/internal/domains/vk_machine/structs"
	"github.com/igorshmel/lic_auto_post/app/pkg/dbo"
	"github.com/igorshmel/lic_auto_post/app/pkg/ddo"
	"github.com/igorshmel/lic_auto_post/app/pkg/errs"
	"github.com/igorshmel/lic_auto_post/app/pkg/lib"
	logger "github.com/igorshmel/lic_auto_post/app/pkg/log"
	"github.com/igorshmel/lic_auto_post/app/pkg/mapping"
	status "github.com/igorshmel/lic_auto_post/app/pkg/vars/statuses"
	"github.com/igorshmel/lic_auto_post/app/pkg/vars/types"
	"github.com/nuttech/bell/v2"
	"os"
	"time"
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

	// Получить счетчик публикаций за текущий день
	publishCounterDBO := dbo.PublishCounterDBO{} // заполнение перенести в доменную область
	publishCounterDBO.UUID = uuid.New().String()
	publishCounterDBO.Date = lib.TruncateToDay(time.Now())
	publishCounterDBO.Count = 1
	publishCounterDBO.Type = types.ArtPublishType
	count, err := ths.extractor.GetArtPublishCountByDate(ctx, &publishCounterDBO)
	if err != nil {
		return err
	}
	// Если в этот день уже была публикация, то сценарий не выполняется
	if count > 0 {
		return nil
	}

	// Выгрузка случайной записи из БД
	recordDBO := dbo.RecordDBO{}
	if err := ths.extractor.GetByActiveStatus(&recordDBO); err != nil {
		if err.Error() == errs.MsgNotFound {
			return extErr(errs.MsgNotFound, msg("DB_ERR: no one record found"), log)
		}
		return extErr(errs.UnknownError, // TODO актуализировать константы ошибок
			msg("DB_ERR: failed to get RND entity by active status: %s", err.Error()), log)
	}

	log.Debug("recordDBO: %v", recordDBO)

	// Поход в домен - подготовка вызова метода photos.getWallUploadServer
	getWallUploadServerDDO := ths.vkMachineDomain.GetWallUploadServer()

	// Вызов метода
	responseGetWallUploadServer := structs.VkGetWallUploadServer{}
	_, err = lib.Request("v", getWallUploadServerDDO.MethodName, getWallUploadServerDDO.Params, &responseGetWallUploadServer)
	if err != nil {
		return extErr(errs.UnknownError,
			msg("getWallUploadServer failed with error: %s", err.Error()), log)
	}

	// Поход в домен - получение пути до файла
	vkMachineDDO := mapping.RecordDbOtoVkMachineDDO(&recordDBO)
	path := ths.vkMachineDomain.GetPath(vkMachineDDO)

	// Подготовительный этап в формировании загрузки на сайт
	uploadPhotoWallResponse, err := lib.PhotoWall(responseGetWallUploadServer.Response.UploadURL, path)
	if err != nil {

		// Обновление статуса записи в БД на "used"
		recordDBO.Status = status.RecordUsedStatus
		if err := ths.persister.UpdateRecordStatus(&recordDBO); err != nil {
			return extErr(errs.UnknownError,
				msg("failed updateRecordStatus to used with error: %s", err.Error()), log)
		}

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

	// Вызов метода
	responseSaveWallPhoto := structs.VkSaveWallPhoto{}
	_, err = lib.Request("v", saveWallPhotoDDO.MethodName, saveWallPhotoDDO.Params, &responseSaveWallPhoto)
	if err != nil {
		return extErr(errs.UnknownError,
			msg("saveWallPhoto failed with error: %s", err.Error()), log)
	}
	if len(responseSaveWallPhoto.Response) <= 0 {
		return extErr(errs.UnknownError,
			msg("method saveWallPhoto failed: Response index = 0; with error %s", err.Error()), log)
	}

	// Поход в домен - публикация поста на стене группы
	// vkMachineDDO = mapping.RecordDbOtoVkMachineDDO(&recordDBO)
	// ths.vkMachineDomain.UploadPhotoToServer(vkMachineDDO)

	// Поход в домен wall.post
	reqPostWallPhoto := ddo.ReqPostWallPhoto{
		URL:     recordDBO.URL,
		AuthURL: recordDBO.AuthURL,
		OwnerID: responseSaveWallPhoto.Response[0].OwnerID,
		ID:      responseSaveWallPhoto.Response[0].ID,
	}
	resPostWallPhoto := ths.vkMachineDomain.PostWallPhoto(&reqPostWallPhoto)

	_, err = lib.Request("v", resPostWallPhoto.MethodName, resPostWallPhoto.Params, nil)
	if err != nil {
		return extErr(errs.UnknownError,
			msg("ResPostWallPhoto failed with error: %s", err.Error()), log)
	}

	getUploadServer := ths.vkMachineDomain.GetUploadServer()
	resGetUploadServer := structs.VkGetWallUploadServer{}
	_, err = lib.Request("v", getUploadServer.MethodName, getUploadServer.Params, &resGetUploadServer)
	if err != nil {
		return extErr(errs.UnknownError,
			msg("GetUploadServer failed with error: %s", err.Error()), log)
	}

	// загрузить картинку по ссылке полученной ранее
	uploadPhoto, err := lib.PhotoGroup(resGetUploadServer.Response.UploadURL, path)
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
	_, err = lib.Request("v", photoSave.MethodName, photoSave.Params, nil)
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

	if err := ths.persister.SetArtPublishCount(ctx, &publishCounterDBO); err != nil {
		return extErr(errs.UnknownError,
			msg("failed to set artPublishCount  with error: %s", err.Error()), log)
	}

	// удалить файл
	if err = os.Remove(path); err != nil {
		ths.log.Error("error Remove file: " + fmt.Sprint(err))
	}

	return nil
}
