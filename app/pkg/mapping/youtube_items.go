package mapping

import (
	"github.com/google/uuid"
	"github.com/igorshmel/lic_auto_post/app/internal/adapters/repository/models"
	"github.com/igorshmel/lic_auto_post/app/internal/adapters/repository/models/basis"
	"github.com/igorshmel/lic_auto_post/app/pkg/dbo"
	"github.com/igorshmel/lic_auto_post/app/pkg/ddo"
	"github.com/igorshmel/lic_auto_post/app/pkg/deo"
	"github.com/igorshmel/lic_auto_post/app/pkg/dto"
	status "github.com/igorshmel/lic_auto_post/app/pkg/vars/statuses"
	"time"
)

// ConvertVideosInfoDTOtoDEO --
func ConvertVideosInfoDTOtoDEO(items []dto.VideoInfo) []deo.VideoInfo {
	var videosInfo []deo.VideoInfo
	for _, item := range items {
		videosInfo = append(videosInfo,
			deo.VideoInfo{
				Title:   item.Title,
				VideoID: item.VideoID,
			})
	}
	return videosInfo
}

// ConvertVideosInfoDEOtoDTO --
func ConvertVideosInfoDEOtoDTO(items []deo.VideoInfo) []dto.VideoInfo {
	var videosInfo []dto.VideoInfo
	for _, item := range items {
		videosInfo = append(videosInfo,
			dto.VideoInfo{
				Title:   item.Title,
				VideoID: item.VideoID,
			})
	}
	return videosInfo
}

// ConvertVideosInfoDTOtoDBO --
func ConvertVideosInfoDTOtoDBO(dto dto.VideoInfo) dbo.IsVideoExistsDBO {
	return dbo.IsVideoExistsDBO{
		VideoID: dto.VideoID,
	}
}

// ConvertDDOItemToDBO --
func ConvertDDOItemToDBO(ddo ddo.ResGetItems) dbo.IsVideoExistsDBO {
	return dbo.IsVideoExistsDBO{
		VideoID: ddo.VideoID,
	}
}

// ConvertIsVideoExistsDBOtoModel --
func ConvertIsVideoExistsDBOtoModel(dbo *dbo.IsVideoExistsDBO) *models.YoutubeItemsModel {
	return &models.YoutubeItemsModel{
		VideoID: dbo.VideoID,
	}
}

// YoutubeSaveNewYoutubeItemsDBOtoModel --
func YoutubeSaveNewYoutubeItemsDBOtoModel(dbo *dbo.SaveNewYoutubeItemsDBO) *models.YoutubeItemsModel {
	base := basis.BaseModel{}
	base.UUID = dbo.UUID
	base.CreatedAt = dbo.CreatedAt
	base.UpdatedAt = dbo.UpdatedAt
	return &models.YoutubeItemsModel{
		Title:     dbo.Title,
		VideoID:   dbo.VideoID,
		Status:    dbo.Status,
		BaseModel: base,
	}
}

// YoutubeSaveNewYoutubeItemsDtOtoDBO --
func YoutubeSaveNewYoutubeItemsDtOtoDBO(dto *dto.VideoInfo) *dbo.SaveNewYoutubeItemsDBO {
	return &dbo.SaveNewYoutubeItemsDBO{
		UUID:      uuid.New().String(),
		Title:     dto.Title,
		VideoID:   dto.VideoID,
		Status:    status.RecordActiveStatus,
		CreatedAt: time.Now(),
	}
}

// YoutubeNextCursorDBOtoModel --
func YoutubeNextCursorDBOtoModel(dbo *dbo.YoutubeNextCursorDBO) *models.YoutubeNextCursorModel {
	base := basis.BaseModel{}
	base.CreatedAt = dbo.CreatedAt
	base.UpdatedAt = dbo.UpdatedAt
	return &models.YoutubeNextCursorModel{
		NextPageToken: dbo.NextPageToken,
		BaseModel:     base,
	}
}

// ConvertDBOToYoutubeDDO --
func ConvertDBOToYoutubeDDO(dbo dbo.YoutubeItemDBO) *ddo.YoutubeItemDDO {
	return &ddo.YoutubeItemDDO{
		Title:   dbo.Title,
		VideoID: dbo.VideoID,
	}
}
