package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/koderpark/vocadb-metadata-macro/service"
	"github.com/koderpark/vocadb-metadata-macro/service/vocadb"
)

func main() {
	dir := flag.String("dir", "", "directory of audio files to read tags from")
	url := flag.String("url", "", "VocaDB album URL to fetch metadata for")
	flag.Parse()

	if *url != "" {
		id, err := vocadb.ParseAlbumID(*url)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}

		album, err := vocadb.FetchAlbum(id)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}

		printJSON(album)
		return
	}

	if *dir == "" {
		fmt.Fprintln(os.Stderr, "usage: chk -url <vocadb-album-url> | -dir <directory>")
		os.Exit(1)
	}

	files, err := service.ListFiles(*dir)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	type fileMetadata struct {
		Path     string           `json:"path"`
		Metadata service.Metadata `json:"metadata"`
	}

	result := make([]fileMetadata, 0, len(files))
	for _, f := range files {
		meta, err := service.ReadMetadata(f)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}
		result = append(result, fileMetadata{Path: f, Metadata: meta})
	}

	printJSON(result)
}

// printJSON은 값을 들여쓰기된 JSON으로 표준출력에 기록한다.
func printJSON(v any) {
	out, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
	fmt.Println(string(out))
}
