package service

import (
	"fmt"
	"strconv"

	"github.com/koderpark/vocadb-metadata-macro/service/vocadb"
)

// discTypeCompilation은 VocaDB가 컴필레이션 앨범에 부여하는 discType 값이다.
// 이 값일 때만 COMPILATION 태그를 "1"로 채운다. (plan.md 방식 #1)
const discTypeCompilation = "Compilation"

// AlbumToMetadata는 VocaDB에서 긁어온 Album을 taglib로 쓸 수 있는
// 트랙별 Metadata 슬라이스로 변환한다. (앨범의 트랙 순서를 그대로 보존)
// 매핑 규칙은 metadata.md의 "VocaDB → 태그 매핑" 표를 따른다.
func AlbumToMetadata(album *vocadb.Album) []Metadata {
	if album == nil {
		return nil
	}

	date := formatReleaseDate(album.ReleaseDate)
	compilation := ""
	if album.DiscType == discTypeCompilation {
		compilation = "1"
	}

	out := make([]Metadata, 0, len(album.Tracks))
	for _, t := range album.Tracks {
		out = append(out, Metadata{
			Title:       t.Name,
			Artist:      t.Song.ArtistString,
			Album:       album.Name,
			AlbumArtist: album.ArtistString,
			Date:        date,
			Genre:       "", // VocaDB 앨범 API는 장르를 제공하지 않는다.
			TrackNumber: numTag(t.TrackNumber),
			DiscNumber:  numTag(t.DiscNumber),
			Compilation: compilation,
		})
	}

	return out
}

// formatReleaseDate는 VocaDB ReleaseDate를 ISO 8601 부분 날짜 문자열로 변환한다.
// 정보가 없는 만큼만 잘라낸다: "2020-05-15" / "2020-05" / "2020" / "".
func formatReleaseDate(d vocadb.ReleaseDate) string {
	if d.IsEmpty || d.Year == 0 {
		return ""
	}
	if d.Month == 0 {
		return fmt.Sprintf("%04d", d.Year)
	}
	if d.Day == 0 {
		return fmt.Sprintf("%04d-%02d", d.Year, d.Month)
	}
	return fmt.Sprintf("%04d-%02d-%02d", d.Year, d.Month, d.Day)
}

// numTag는 트랙/디스크 번호를 태그 문자열로 변환한다.
// 0 이하(미지정)는 빈 문자열로 둔다.
func numTag(n int) string {
	if n <= 0 {
		return ""
	}
	return strconv.Itoa(n)
}
