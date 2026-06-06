package service

import (
	"go.senan.xyz/taglib"
)

// Metadata는 음원 파일에서 읽거나 쓸 핵심 메타데이터를 표현한다.
// 모든 필드는 태그 저장 모델에 맞춰 문자열로 다룬다. (트랙/디스크 번호 포함)
type Metadata struct {
	Title       string `json:"title"`
	Artist      string `json:"artist"`
	Album       string `json:"album"`
	AlbumArtist string `json:"albumArtist"`
	Date        string `json:"date"`
	Genre       string `json:"genre"`
	TrackNumber string `json:"trackNumber"`
	DiscNumber  string `json:"discNumber"`
	Compilation string `json:"compilation"` // 컴필레이션이면 "1", 아니면 ""
}

// supportedExtensions는 읽기/쓰기를 지원하는 음원 확장자 집합이다.
// (소문자, 점 포함) 확장자를 추가하려면 이 맵에 항목만 추가하면 된다.
var supportedExtensions = map[string]struct{}{
	".flac": {},
	".mp3":  {},
}

// ReadMetadata는 음원 파일의 태그를 읽어 Metadata로 반환한다.
func ReadMetadata(path string) (Metadata, error) {
	tags, err := taglib.ReadTags(path)
	if err != nil {
		return Metadata{}, err
	}

	return Metadata{
		Title:       firstTag(tags[taglib.Title]),
		Artist:      firstTag(tags[taglib.Artist]),
		Album:       firstTag(tags[taglib.Album]),
		AlbumArtist: firstTag(tags[taglib.AlbumArtist]),
		Date:        firstTag(tags[taglib.Date]),
		Genre:       firstTag(tags[taglib.Genre]),
		TrackNumber: firstTag(tags[taglib.TrackNumber]),
		DiscNumber:  firstTag(tags[taglib.DiscNumber]),
		Compilation: firstTag(tags[taglib.Compilation]),
	}, nil
}

// WriteMetadata는 Metadata를 음원 파일의 태그로 기록한다.
// 명시한 태그만 갱신하고, 나머지 기존 태그는 보존한다.
func WriteMetadata(path string, m Metadata) error {
	return taglib.WriteTags(path, map[string][]string{
		taglib.Title:       {m.Title},
		taglib.Artist:      {m.Artist},
		taglib.Album:       {m.Album},
		taglib.AlbumArtist: {m.AlbumArtist},
		taglib.Date:        {m.Date},
		taglib.Genre:       {m.Genre},
		taglib.TrackNumber: {m.TrackNumber},
		taglib.DiscNumber:  {m.DiscNumber},
		taglib.Compilation: {m.Compilation},
	}, 0)
}

// firstTag는 멀티값 태그에서 첫 번째 값을 반환한다. (없으면 빈 문자열)
func firstTag(values []string) string {
	if len(values) == 0 {
		return ""
	}
	return values[0]
}
