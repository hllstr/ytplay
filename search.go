package main

import (
	"fmt"
	"github.com/raitonoberu/ytsearch"
)

func SearchSongs(query string) ([]*ytsearch.VideoItem, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Terjadi Panic saat search lagu! %v\n", r)
		}
	}()

	search := ytsearch.VideoSearch(query)
	result, err := search.Next()
	if err != nil {
		return nil, err
	}
	results := result.Videos
	//Limit ke 8 biar gak kebanyakan
	if len(results) > 8 {
		results = results[:8]
	}
	return results, nil
}
