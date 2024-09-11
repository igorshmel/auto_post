package youtube_usecase

import (
	"context"
	"fmt"
	"github.com/igorshmel/lic_auto_post/app/internal/adapters/port"
	transport "github.com/igorshmel/lic_auto_post/app/internal/adapters/transport/youtube"
	"github.com/igorshmel/lic_auto_post/app/pkg/config"
	"github.com/igorshmel/lic_auto_post/app/pkg/deo"
	"github.com/igorshmel/lic_auto_post/app/pkg/dto"
	logger "github.com/igorshmel/lic_auto_post/app/pkg/log"
	"github.com/igorshmel/lic_auto_post/app/pkg/mapping"
	"github.com/igorshmel/lic_auto_post/app/pkg/vars/constants"
	"github.com/nuttech/bell/v2"
	"google.golang.org/api/youtube/v3"
	"time"
)

// GetPlayListUseCase --
type GetPlayListUseCase struct {
	cfg           config.Config
	log           logger.Logger
	bell          *bell.Events
	persister     port.Persister
	extractor     port.Extractor
	youtubeClient *transport.YouTubeClient
}

// NewGetPlayListUseCase --
func NewGetPlayListUseCase(
	cfg config.Config,
	log logger.Logger,
	events *bell.Events,
	persister port.Persister,
	extractor port.Extractor,
	youtubeClient *transport.YouTubeClient,
) port.GetPlayListUseCase {
	return GetPlayListUseCase{
		cfg:           cfg,
		log:           log,
		bell:          events,
		persister:     persister,
		extractor:     extractor,
		youtubeClient: youtubeClient,
	}
}

// Execute _
func (ths GetPlayListUseCase) Execute(ctx context.Context, req *dto.GetPlayListReqDTO) error {
	var videosInfo []dto.VideoInfo
	var nextPageToken, prevPageToken string
	allItems := youtube.PlaylistItemListResponse{}

	log := ths.log.WithMethod("usecase GetPlayList")
	log.Info("Try GetPlayList with req: %v", req)
	// -- Бизнес логика --
	// ---------------------------------------------------------------------------------------------------------------------------

	// -- Инфраструктурная логика --
	// ---------------------------------------------------------------------------------------------------------------------------
	for {
		time.Sleep(100 * time.Microsecond)
		response, err := ths.youtubeClient.GetPlayListItems(req.PlayListID, nextPageToken)
		if err != nil {
			return err
		}

		allItems.Items = append(allItems.Items, response.Items...)

		if response.NextPageToken == "" {
			break
		}
		nextPageToken = response.NextPageToken
		prevPageToken = response.PrevPageToken
	}
	allItems.PrevPageToken = prevPageToken

	for _, item := range allItems.Items {
		fmt.Printf("Title: %s	| ", item.Snippet.Title)
		fmt.Printf("VideoID: %s \n", item.Snippet.ResourceId.VideoId)
		videosInfo = append(videosInfo, dto.VideoInfo{Title: item.Snippet.Title, VideoId: item.Snippet.ResourceId.VideoId})
	}

	// -- Периферия --
	// ---------------------------------------------------------------------------------------------------------------------------

	// отправка события YouTubeGetPlayListDoneEventName для получения данных по playlist
	if err := ths.bell.Ring(
		constants.YouTubeGetPlayListDoneEventName,
		deo.SaveNewYoutubeItemsEvent{
			VideosInfo:           mapping.ConvertVideosInfoDTOtoDEO(videosInfo),
			NextCursorPagination: nextPageToken,
		}); err != nil {

		ths.log.Error("unable send event DownloadImage with error: %s", err.Error())
	}
	log.Debug("client: response body")

	return nil
}
