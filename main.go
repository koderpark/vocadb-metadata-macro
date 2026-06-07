package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/koderpark/vocadb-metadata-macro/service"
	"github.com/koderpark/vocadb-metadata-macro/service/vocadb"
)

func main() {
	url := flag.String("url", "", "VocaDB album URL to source metadata from")
	dir := flag.String("dir", "", "directory of audio files to write tags to")
	flag.Parse()

	if *url == "" || *dir == "" {
		fmt.Fprintln(os.Stderr, "usage: vocadb-metadata-macro -url <vocadb-album-url> -dir <directory>")
		os.Exit(1)
	}

	// VocaDB URL에서 앨범 ID를 뽑아 앨범 메타데이터를 받아온다.
	id, err := vocadb.ParseAlbumID(*url)
	if err != nil {
		fatal(err)
	}
	album, err := vocadb.FetchAlbum(id)
	if err != nil {
		fatal(err)
	}

	// 앨범을 트랙별 태그 메타데이터로 변환한다. (앨범의 트랙 순서 보존)
	metas := service.AlbumToMetadata(album)

	// 태그를 기록할 대상 음원 파일 목록. (파일명 정렬 순)
	files, err := service.ListFiles(*dir)
	if err != nil {
		fatal(err)
	}

	// 트랙 수와 파일 수가 다르면 잘못된 짝이 생기므로 기록하지 않고 중단한다.
	if len(files) != len(metas) {
		fatal(fmt.Errorf("track count mismatch: album %q has %d tracks but %s has %d audio files",
			album.Name, len(metas), *dir, len(files)))
	}

	// 순서대로 짝지어 각 파일에 태그를 기록한다.
	for i, f := range files {
		if err := service.WriteMetadata(f, metas[i]); err != nil {
			fatal(fmt.Errorf("writing tags to %s: %w", f, err))
		}
		fmt.Printf("tagged %s\t%s - %s\n", f, metas[i].Artist, metas[i].Title)
	}

	fmt.Printf("done: wrote tags to %d files in %s\n", len(files), *dir)
}

// fatal은 에러를 표준에러에 출력하고 프로세스를 종료한다.
func fatal(err error) {
	fmt.Fprintln(os.Stderr, "error:", err)
	os.Exit(1)
}
