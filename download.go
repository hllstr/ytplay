package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/kkdai/youtube/v2"
)

// function untuk Download & Save audio (auto pick best quality)
func DownloadAudio(videoID, dst string) error {
	client := youtube.Client{}
	video, err := client.GetVideo(videoID)
	if err != nil {
		return err
	}
	formats := video.Formats.WithAudioChannels()
	if len(formats) == 0 {
		fmt.Println(`Tidak dapat menemukan Audio Format :(`)
		return err
	}
	formats.Sort()
	best := formats[0]
	stream, _, err := client.GetStream(video, &best)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
		return err
	}

	file, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()

	_, err = io.Copy(file, stream)
	return err
}
