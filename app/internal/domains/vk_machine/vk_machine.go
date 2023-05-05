package vkmachine

import (
	"auto_post/app/internal/domains/vk_machine/structs"
	"auto_post/app/pkg/config"
	"auto_post/app/pkg/ddo"
	"auto_post/app/pkg/helpers"
	logger "auto_post/app/pkg/log"
	"auto_post/app/pkg/vars/constants"
	"fmt"
)

// VkMachine --
type VkMachine struct {
	log            logger.Logger
	cfg            config.Config
	vkIni          structs.VkPicIni
	getWallUploads structs.VkGetWallUploadS
	saveWall       structs.VkSaveWallPhoto
}

// NewVkMachine - инициализация домена VkMachine
func NewVkMachine(log logger.Logger, cfg config.Config) *VkMachine {
	log = log.WithMethod("youTubeMachineDomain")
	return &VkMachine{log: log, cfg: cfg}
}

// UploadPhotoToServer --
func (ths VkMachine) UploadPhotoToServer(ddo *ddo.VKMachineDDO) {
	l := ths.log

	l.Debug("conf VK: %s", ths.cfg.VkConfig.VkAlbumID)

	paramsWallPhotosGetUploadServer := map[string]string{
		"group_id":     ths.cfg.VkConfig.VkGroupID,
		"access_token": ths.cfg.VkConfig.VkToken,
		"v":            constants.APIVersion,
	}

	l.Debug("postToVk | getWallUploadServerReq: %v", paramsWallPhotosGetUploadServer)

	responseWallUploadServer, err := helpers.Request("v", "photos.getWallUploadServer", paramsWallPhotosGetUploadServer, &ths.getWallUploads)

	l.Debug("postToVk | getWallUploadServer Response: %s", ths.getWallUploads.Response)

	//path := fmt.Sprint(postPic.Id) + ".jpg" // путь к файлу
	path := ths.cfg.DownloadMachine.Path + ddo.FileName + ".jpg" // путь к файлу
	if err != nil {
		err = fmt.Errorf("photos.getWallUploadServer failed: %v | with error: %s", responseWallUploadServer, err.Error())
	}

	// подготовительный этап в формировании загрузки на сайт
	uploaded, err := helpers.PhotoWall(ths.getWallUploads.Response.UploadURL, path)

	l.Debug("postToVk | uploaded %v", uploaded)
	if err != nil {
		err = fmt.Errorf("error uploaded photoWall: %s", err.Error())
	}
	// сохранить картинку для поста на стене группы
	paramsSaveWallPhoto := map[string]string{
		"group_id":     ths.cfg.VkConfig.VkGroupID,
		"photo":        uploaded.Photo,
		"hash":         uploaded.Hash,
		"server":       fmt.Sprint(uploaded.Server),
		"access_token": ths.cfg.VkConfig.VkToken,
		"v":            constants.APIVersion,
	}

	responseSaveWallPhoto, err := helpers.Request("v", "photos.saveWallPhoto", paramsSaveWallPhoto, &ths.saveWall)

	l.Debug("postToVk | saveWallPhoto %v", ths.saveWall)
	if err != nil {
		err = fmt.Errorf("photos.saveWallPhoto failed: %s; \n response: %s", err.Error(), responseSaveWallPhoto)
	}

	if len(ths.saveWall.Response) <= 0 {
		err = fmt.Errorf("method saveWallPhoto failed: Response index = 0; with error %s", err)
	}
}
