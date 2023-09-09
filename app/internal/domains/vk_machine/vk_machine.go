package vkmachine

import (
	"fmt"
	"github.com/igorshmel/lic_auto_post/app/internal/domains/vk_machine/structs"
	"github.com/igorshmel/lic_auto_post/app/pkg/config"
	"github.com/igorshmel/lic_auto_post/app/pkg/ddo"
	"github.com/igorshmel/lic_auto_post/app/pkg/lib"
	logger "github.com/igorshmel/lic_auto_post/app/pkg/log"
	"github.com/igorshmel/lic_auto_post/app/pkg/vars/constants"
	"strings"
)

// VkMachine --
type VkMachine struct {
	log            logger.Logger
	cfg            config.Config
	vkIni          structs.VkPicIni
	getWallUploads structs.VkGetWallUploadServer
	saveWall       structs.VkSaveWallPhoto
	getURL         structs.VkGetWallUploadServer
}

// NewVkMachine - инициализация домена VkMachine
func NewVkMachine(log logger.Logger, cfg config.Config) *VkMachine {
	log = log.WithMethod("VkMachineDomain")
	return &VkMachine{log: log, cfg: cfg}
}

// GetWallUploadServer --
func (ths VkMachine) GetWallUploadServer() *ddo.GetWallUploadServer {

	getWallUploadServerParams := map[string]string{
		"group_id":     ths.cfg.VkConfig.VkGroupID,
		"access_token": ths.cfg.VkConfig.VkToken,
		"v":            constants.APIVersion,
	}

	ths.log.Debug("getWallUploadServerParams: %v", getWallUploadServerParams)

	return &ddo.GetWallUploadServer{
		Params:     getWallUploadServerParams,
		MethodName: "photos.getWallUploadServer",
	}
}

// GetPath --
func (ths VkMachine) GetPath(req *ddo.VKMachine) string {
	return fmt.Sprintf("%s/%s.jpg", ths.cfg.DownloadMachine.Path, req.FileName) // путь к файлу
}

// SaveWallPhoto --
func (ths VkMachine) SaveWallPhoto(req *ddo.ReqSaveWallPhoto) *ddo.ResSaveWallPhoto {

	// сохранить картинку для поста на стене группы
	saveWallPhotoParams := map[string]string{
		"group_id":     ths.cfg.VkConfig.VkGroupID,
		"photo":        req.Photo,
		"hash":         req.Hash,
		"server":       fmt.Sprint(req.Server),
		"access_token": ths.cfg.VkConfig.VkToken,
		"v":            constants.APIVersion,
	}

	return &ddo.ResSaveWallPhoto{
		Params:     saveWallPhotoParams,
		MethodName: "photos.saveWallPhoto",
	}
}

// GetTags --
func (ths VkMachine) GetTags() string {
	// Подготовка хэш-тегов TODO оптимизировать процесс
	var tags []string
	var stringTags string
	if ths.cfg.VkConfig.VkHashTags != "" {
		tags = strings.Split(ths.cfg.VkConfig.VkHashTags, ",")
	}

	arr := lib.RangeInt(0, len(tags), ths.cfg.VkConfig.MaxHashTags)

	if arr != nil {
		for r := 0; r < len(arr); r++ {
			stringTags = stringTags + " #" + fmt.Sprint(tags[arr[r]])
		}
	}

	return stringTags
}

// PostWallPhoto --
func (ths VkMachine) PostWallPhoto(req *ddo.ReqPostWallPhoto) *ddo.ResPostWallPhoto {

	// Подготовка хэш-тегов TODO оптимизировать процесс
	var tags []string
	var stringTags string
	if ths.cfg.VkConfig.VkHashTags != "" {
		tags = strings.Split(ths.cfg.VkConfig.VkHashTags, ",")
	}

	arr := lib.RangeInt(0, len(tags), ths.cfg.VkConfig.MaxHashTags)

	if arr != nil {
		for r := 0; r < len(arr); r++ {
			stringTags = stringTags + " #" + fmt.Sprint(tags[arr[r]])
		}
	}

	// формируем правильное название для загружаемого файла
	photoID := "photo" + fmt.Sprint(req.OwnerID) + "_" + fmt.Sprint(req.ID)

	//пост картинки с описанием в группу
	postWallPhotoParams := map[string]string{
		"owner_id":     "-" + ths.cfg.VkConfig.VkGroupID,
		"from_group":   "1",
		"message":      "Источник: " + req.AuthURL + "\n\r\n\r" + strings.Trim(stringTags, " "),
		"attachments":  photoID,
		"access_token": ths.cfg.VkConfig.VkToken,
		"v":            constants.APIVersion,
	}

	return &ddo.ResPostWallPhoto{
		Params:     postWallPhotoParams,
		MethodName: "wall.post",
	}
}

// GetUploadServer --
func (ths VkMachine) GetUploadServer() *ddo.ResGetUploadServer {

	// ~~~~ Пост картинки с описанием в альбом группы ~~~~
	// получить ссылку по которой можно загрузить картинку на сайт ВК
	paramsGetUploadServer := map[string]string{
		"group_id":     ths.cfg.VkConfig.VkGroupID,
		"album_id":     ths.cfg.VkConfig.VkAlbumID,
		"access_token": ths.cfg.VkConfig.VkToken,
		"v":            constants.APIVersion,
	}

	return &ddo.ResGetUploadServer{
		Params:     paramsGetUploadServer,
		MethodName: "photos.getUploadServer",
	}
}

// PhotosSave --
func (ths VkMachine) PhotosSave(req ddo.ReqPhotosSave) *ddo.ResPhotosSave {

	// сохранить картинку в альбом
	paramsPhotosSave := map[string]string{
		"group_id":     ths.cfg.VkConfig.VkGroupID,
		"album_id":     ths.cfg.VkConfig.VkAlbumID,
		"photos_list":  req.PhotoList,
		"server":       fmt.Sprint(req.Server),
		"hash":         req.Hash,
		"caption":      "Источник: " + req.URL,
		"access_token": ths.cfg.VkConfig.VkToken,
		"v":            constants.APIVersion,
	}

	return &ddo.ResPhotosSave{
		Params:     paramsPhotosSave,
		MethodName: "photos.save",
	}
}
