package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		// Prompt & Read user Input
		fmt.Print("Mau cari lagu apa? ")
		query, _ := reader.ReadString('\n')
		query = strings.TrimSpace(query)

		// Search
		videos, err := SearchSongs(query)
		if err != nil {
			fmt.Println("Err nyari lagu : ", err)
			continue
		}
		if len(videos) == 0 {
			fmt.Println("Tidak ditemukan hasil :(\nCoba cari dengan kata kunci lain.")
			continue
		}
		// Show list (format: [n] Title - Channel)
		for i, v := range videos {
			fmt.Printf("[%d] %s - %s\n", i+1, v.Title, v.Channel.Title)
		}
		fmt.Printf("Pilih nomor (1-%d) ", len(videos))
		userChoice, _ := reader.ReadString('\n')
		choice, err := strconv.Atoi(strings.TrimSpace(userChoice))
		if err != nil || choice < 1 || choice > len(videos) {
			fmt.Println("Pilihan tidak valid.")
			continue
		}
		vid := videos[choice-1]
		// Download & Save Audio
		dst := filepath.Join(userDownloads(), fmt.Sprintf("%s - %s [%s]", vid.Title, vid.Channel.Title, vid.ID))
		if _, err := os.Stat(dst); err == nil {
			fmt.Println("File audio sudah ada skipping Download...")
			if err := Play(dst); err != nil {
				fmt.Println("Err Playing Audio : ", err)
			}
			continue
		}
		// Kalau belum ada, kanjut donglod
		fmt.Println("Mengunduh : ", vid.Title)
		if err := DownloadAudio(vid.ID, dst); err != nil {
			fmt.Println("Gagal Download :", err)
			continue
		}

		fmt.Println("Playing : ", dst)
		if err := Play(dst); err != nil {
			fmt.Println("Error memutar : ", err)
		}

		fmt.Print("Mau cari lagi? (y/n) : ")
		again, _ := reader.ReadString('\n')
		if strings.ToLower(strings.TrimSpace(again)) != "y" {
			fmt.Println("Makasih udah pake program ini <3 - hllstr.")
			break
		}
	}
}

// Menentukan direktori Downloads
func userDownloads() string {
	// Untuk termux/android kita simpen hasil Download nya ke /sdcard/Download
	// buat dan simpan ke folder YTPLAY biar rapih (gak campur aduk sama downloadan lainnya)
	if runtime.GOOS == "android" {
		return filepath.Join("/sdcard/Download/YTPLAY")
	} else {
		home, _ := os.UserHomeDir()
		return filepath.Join(home, "Downloads/YTPLAY")
	}
}
