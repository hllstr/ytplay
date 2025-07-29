package main

import (
	"fmt"
	"io"
)

type ProgressReader struct {
	reader      io.Reader
	total       int64
	current     int64
	lastUpdate  int64
	finished    bool
	onProgress  func(current, total int64)
}

func NewProgressReader(reader io.Reader, total int64, onProgress func(current, total int64)) *ProgressReader {
	return &ProgressReader{
		reader:     reader,
		total:      total,
		onProgress: onProgress,
	}
}

func (pr *ProgressReader) Read(p []byte) (n int, err error) {
	n, err = pr.reader.Read(p)
	pr.current += int64(n)
	
	if pr.finished {
		return
	}
	
	threshold := pr.total / 200
	if threshold < 1024 {
		threshold = 1024
	}

	isComplete := pr.current >= pr.total || err == io.EOF
	shouldUpdate := pr.current-pr.lastUpdate >= threshold || isComplete
	
	if pr.onProgress != nil && shouldUpdate && !pr.finished {
		if isComplete {
			pr.finished = true
			pr.current = pr.total
		}
		
		pr.onProgress(pr.current, pr.total)
		pr.lastUpdate = pr.current
	}
	return
}

func ShowProgress(current, total int64) {
	if total <= 0 {
		return
	}
	
	percentage := float64(current) / float64(total) * 100
	if percentage > 100 {
		percentage = 100
	}
	
	barWidth := 25
	filled := int(percentage / 100 * float64(barWidth))
	
	bar := ""
	for i := 0; i < barWidth; i++ {
		if i < filled {
			bar += "#"
		} else {
			bar += "."
		}
	}
	
	currentSize := formatBytes(current)
	totalSize := formatBytes(total)
	
	fmt.Printf("\r\033[K")
	fmt.Printf("Progress: [%3.0f%%] [%s] %s/%s", 
		percentage, bar, currentSize, totalSize)
	
	if current >= total {
		fmt.Println()
	}
}

func formatBytes(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%dB", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f%cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}