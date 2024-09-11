package youtube

import (
	"fmt"
	"google.golang.org/api/youtube/v3"
	"testing"
)

func TestYouTubeClient_GetPlayListItems(t *testing.T) {
	type fields struct {
		client *YouTubeClient
	}
	type args struct {
		playListID string
	}

	client, err := NewYouTubeClient("AIzaSyCIuulOtq5ZebCqhA_ElCuQvIM01shbApw")
	if err != nil {
		return
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *youtube.PlaylistItemListResponse
		wantErr bool
	}{{
		name: "Valid playlistID",
		fields: fields{
			client: client,
		},
		args: args{
			playListID: "PLgicPnEfofJpsLS7wnv7Cs7mC7jPlHSbj",
		},
		want: &youtube.PlaylistItemListResponse{
			Items: []*youtube.PlaylistItem{
				{
					Id: "item1",
					Snippet: &youtube.PlaylistItemSnippet{
						Title: "Test Title 1",
					},
				},
			},
		},
		wantErr: false,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			items, err := client.GetPlayListItems(tt.args.playListID, "")
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPlayListItems() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Printf("NT: %s \n", items.NextPageToken)
			fmt.Printf("PT: %s \n", items.PrevPageToken)
			fmt.Printf("Pagination: %v \n", items.TokenPagination)
			for _, item := range items.Items {
				fmt.Printf("Title: %s	| ", item.Snippet.Title)
				fmt.Printf("ID: %s \n", item.Snippet.ResourceId.VideoId)
			}
			/*if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPlayListItems() got = %v, want %v", got, tt.want)
			}*/
		})
	}
}
