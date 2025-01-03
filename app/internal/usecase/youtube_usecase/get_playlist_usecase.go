package youtube_usecase

import (
	"context"
	"fmt"
	"github.com/igorshmel/lic_auto_post/app/internal/adapters/port"
	transport "github.com/igorshmel/lic_auto_post/app/internal/adapters/transport/youtube"
	"github.com/igorshmel/lic_auto_post/app/pkg/config"
	logger "github.com/igorshmel/lic_auto_post/app/pkg/log"
	"github.com/igorshmel/lic_auto_post/app/pkg/vars"
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
	youtubeDomain port.YoutubeMachineDomain
}

// NewGetPlayListUseCase --
func NewGetPlayListUseCase(
	cfg config.Config,
	log logger.Logger,
	events *bell.Events,
	persister port.Persister,
	extractor port.Extractor,
	youtubeClient *transport.YouTubeClient,
	youtubeDomain port.YoutubeMachineDomain,
) port.GetPlayListUseCase {
	return GetPlayListUseCase{
		cfg:           cfg,
		log:           log,
		bell:          events,
		persister:     persister,
		extractor:     extractor,
		youtubeClient: youtubeClient,
		youtubeDomain: youtubeDomain,
	}
}

// Execute _
func (ths GetPlayListUseCase) Execute(ctx context.Context) error {
	var nextPageToken, prevPageToken string
	allItems := youtube.PlaylistItemListResponse{}
	playlistID := ths.cfg.YouTubeConfig.YouTubePlayListID

	requestID, ok := ctx.Value(vars.RequestIDKey).(string)
	if !ok {
		return fmt.Errorf("requestID not found in context")
	}

	log := ths.log.WithMethod(requestID + " | usecase GetPlayList")
	log.Info("start")

	// ~~~~ Инфраструктурная логика ~~~~
	for {
		time.Sleep(100 * time.Microsecond)
		response, err := ths.youtubeClient.GetPlayListItems(playlistID, nextPageToken)
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

	// ~~~~  Бизнес логика ~~~~
	for _, item := range allItems.Items {
		ths.youtubeDomain.KeepItems(requestID, item.Snippet.Title, item.Snippet.ResourceId.VideoId)
	}

	// ~~~~ Периферия ~~~~
	// отправка события Done_GetPlaylist_Event для получения данных по playlist
	if err := ths.bell.Ring(
		constants.Done_GetPlaylist_Event, ctx); err != nil {
		ths.log.Error("unable send event Done_GetPlaylist_Event with error: %s", err.Error())
	}
	log.Debug("client: response body")

	return nil
}
