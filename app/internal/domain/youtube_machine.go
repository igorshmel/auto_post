package domain

import (
	"auto_post/app/internal/domain/models"
	"auto_post/app/pkg/helpers"
	"auto_post/app/pkg/vars/constants"
	"errors"
	"fmt"
)

type YouTubeMachine struct {
	models.VkPicIni
	models.VkGetWallUploadS
	saveWall models.VkSaveWallPhoto
}

func (ths *Domain) newYouTubeMachine() {

}

func (ths *Domain) uploadPhotoToServer() {
	l := ths.log

	l.Debug("conf VK: %s", ths.youTubeMachine.VkAlbumId)

	paramsWallPhotosGetUploadServer := map[string]string{
		"group_id":     ths.youTubeMachine.VkGroupId,
		"access_token": ths.youTubeMachine.VkToken,
		"v":            constants.ApiVersion,
	}

	l.Debug("postToVk | getWallUploadServerReq: %v", paramsWallPhotosGetUploadServer)

	responseWallUploadServer, err := helpers.Request("v", "photos.getWallUploadServer", paramsWallPhotosGetUploadServer, &ths.youTubeMachine)

	l.Debug("postToVk | getWallUploadServer Response: %s", ths.youTubeMachine.Response)

	//path := fmt.Sprint(postPic.Id) + ".jpg" // путь к файлу
	path := "6706.jpg" // путь к файлу
	if err != nil {
		err = errors.New(fmt.Sprintf("photos.getWallUploadServer failed: %v | with error: %s", responseWallUploadServer, err.Error()))
	}

	// подготовительный этап в формировании загрузки на сайт
	uploaded, err := helpers.PhotoWall(ths.youTubeMachine.Response.UploadUrl, path)

	l.Debug("postToVk | uploaded %v", uploaded)
	if err != nil {
		err = errors.New(fmt.Sprintf("error uploaded photoWall: %s", err.Error()))
	}
	// сохранить картинку для поста на стене группы
	paramsSaveWallPhoto := map[string]string{
		"group_id":     ths.youTubeMachine.VkGroupId,
		"photo":        uploaded.Photo,
		"hash":         uploaded.Hash,
		"server":       fmt.Sprint(uploaded.Server),
		"access_token": ths.youTubeMachine.VkToken,
		"v":            constants.ApiVersion,
	}

	responseSaveWallPhoto, err := helpers.Request("v", "photos.saveWallPhoto", paramsSaveWallPhoto, &ths.youTubeMachine.saveWall)

	l.Debug("postToVk | saveWallPhoto %v", ths.youTubeMachine.saveWall)
	if err != nil {
		err = errors.New(fmt.Sprintf("photos.saveWallPhoto failed: %s; \n response: %s", err.Error(), responseSaveWallPhoto))
	}

	if len(ths.youTubeMachine.saveWall.Response) <= 0 {
		err = errors.New(fmt.Sprintf("method saveWallPhoto failed: Response index = 0; with error %s", err))

	}

}
