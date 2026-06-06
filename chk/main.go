package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/koderpark/vocadb-metadata-macro/service"
)

func main() {
	dir := flag.String("dir", "", "files to list in this directory")
	url := flag.String("url", "", "VocaDB album URL to fetch metadata for")
	jsonOut := flag.Bool("json", false, "with -url, print the full album metadata as JSON")
	flag.Parse()

	if *url != "" {
		id, err := service.ParseAlbumID(*url)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}

		album, err := service.FetchAlbum(id)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}

		if *jsonOut {
			out, err := json.MarshalIndent(album, "", "  ")
			if err != nil {
				fmt.Fprintln(os.Stderr, "error:", err)
				os.Exit(1)
			}
			fmt.Println(string(out))
			return
		}

		fmt.Printf("ID:           %d\n", album.ID)
		fmt.Printf("Name:         %s\n", album.Name)
		fmt.Printf("Artist:       %s\n", album.ArtistString)
		fmt.Printf("Disc Type:    %s\n", album.DiscType)
		fmt.Printf("Status:       %s\n", album.Status)
		if !album.ReleaseDate.IsEmpty {
			fmt.Printf("Release Date: %04d-%02d-%02d\n",
				album.ReleaseDate.Year, album.ReleaseDate.Month, album.ReleaseDate.Day)
		}
		fmt.Printf("Rating:       %.2f (%d)\n", album.RatingAverage, album.RatingCount)
		return
	}

	if *dir == "" {
		fmt.Fprintln(os.Stderr, "usage: chk -url <vocadb-album-url> [-json] | -dir <directory>")
		os.Exit(1)
	}

	files, err := service.ListFiles(*dir)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	for _, f := range files {
		fmt.Println(f)
	}
}
