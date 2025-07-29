package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/kkdai/youtube/v2"
)

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
	
	contentLength := best.ContentLength
	
	if contentLength <= 0 {
		fmt.Println("Mengunduh... (ukuran tidak diketahui)")
		_, err = io.Copy(file, stream)
		return err
	}

	fmt.Printf("Ukuran file: %s\n", formatBytes(contentLength))
	
	written := int64(0)
	ShowProgress(0, contentLength)

	buffer := make([]byte, 1024)
	
	for {
		n, err := stream.Read(buffer)
		if n > 0 {
			_, writeErr := file.Write(buffer[:n])
			if writeErr != nil {
				return writeErr
			}
			written += int64(n)
			ShowProgress(written, contentLength)
		}
		
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}
	
	return nil
}