package youtube

import (
	"context"
	"fmt"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

// YouTubeClient представляет собой клиента для работы с YouTube API
type YouTubeClient struct {
	service *youtube.Service
}

// NewYouTubeClient создает новый экземпляр YouTubeClient
func NewYouTubeClient(apiKey string) (*YouTubeClient, error) {
	ctx := context.Background()
	service, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}

	return &YouTubeClient{service: service}, nil
}

// GetVideo получает информацию о видео по его ID
func (c *YouTubeClient) GetVideo(videoID string) (*youtube.Video, error) {
	call := c.service.Videos.List([]string{"id", "snippet", "contentDetails", "statistics"}).Id(videoID)
	response, err := call.Do()
	if err != nil {
		return nil, err
	}

	if len(response.Items) == 0 {
		return nil, fmt.Errorf("video with ID %s not found", videoID)
	}

	return response.Items[0], nil
}

// GetChannel получает информацию о канале по его ID
func (c *YouTubeClient) GetChannel(channelID string) (*youtube.Channel, error) {
	call := c.service.Channels.List([]string{"id", "snippet", "statistics"}).Id(channelID)
	response, err := call.Do()
	if err != nil {
		return nil, err
	}

	if len(response.Items) == 0 {
		return nil, fmt.Errorf("channel with ID %s not found", channelID)
	}

	return response.Items[0], nil
}

// GetComments получает комментарии к видео по его ID
func (c *YouTubeClient) GetComments(videoID string) ([]*youtube.Comment, error) {
	call := c.service.CommentThreads.List([]string{"id", "snippet"}).VideoId(videoID)
	response, err := call.Do()
	if err != nil {
		return nil, err
	}

	comments := []*youtube.Comment{}
	for _, item := range response.Items {
		comments = append(comments, item.Snippet.TopLevelComment)
	}

	return comments, nil
}

// GetPlayListItems получает видео из плейлиста
func (c *YouTubeClient) GetPlayListItems(playListID, nextPageToken string) (*youtube.PlaylistItemListResponse, error) {

	call := c.service.PlaylistItems.List([]string{"id", "snippet"}).PlaylistId(playListID).PageToken(nextPageToken)
	response, err := call.Do()
	if err != nil {
		return nil, err
	}

	playListItems := youtube.PlaylistItemListResponse{
		Items:         response.Items,
		NextPageToken: response.NextPageToken,
		PrevPageToken: response.PrevPageToken,
	}

	return &playListItems, nil
}
