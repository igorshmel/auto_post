package main

import (
	"github.com/kkdai/youtube/v2"
	"io"
	"os"
)

func main() {

	// ExampleDownload : Example code for how to use this package for download video.
	videoID := "ke7s8COc1_k"
	client := youtube.Client{}

	video, err := client.GetVideo(videoID)
	if err != nil {
		panic(err)
	}

	formats := video.Formats.WithAudioChannels() // only get videos with audio
	stream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		panic(err)
	}
	defer func(stream io.ReadCloser) {
		err := stream.Close()
		if err != nil {

		}
	}(stream)

	file, err := os.Create("video.mp4")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		panic(err)
	}

}
