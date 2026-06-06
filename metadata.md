# 음원 메타데이터 정리

음원 파일이 메타데이터에 기록하는 값들을 정리한 문서. 포맷마다 내부 표현은 다르지만,
이 프로젝트에서 쓰는 [`go.senan.xyz/taglib`](https://pkg.go.dev/go.senan.xyz/taglib)이
포맷별 표현을 공통 키로 정규화해주므로 그 키 집합을 기준으로 정리한다.

## 포맷별 태그 규격

| 포맷 | 네이티브 태그 규격 |
|---|---|
| MP3 | **ID3** (ID3v1: 고정 128바이트 / ID3v2: 프레임 기반, 확장 가능) |
| FLAC | **Vorbis Comment** (`VORBIS_COMMENT` 블록, `KEY=VALUE`) — ID3 아님 |
| Ogg Vorbis / Opus | Vorbis Comment |
| M4A / AAC / ALAC | MP4 atom (`©nam`, `©ART` 등) |

> FLAC은 ID3를 쓰지 않는다. 일부 SW가 FLAC에 ID3를 박아넣기도 하지만 비표준이며 권장되지 않는다.
> taglib은 위 어떤 포맷이든 아래 **공통 대문자 키**로 정규화해서 읽고/쓴다.

## 가장 흔히 쓰이는 핵심 값

거의 모든 플레이어가 읽고, 이 프로젝트(VocaDB 앨범 태깅)에서 실제로 쓸 값들:

`TITLE` · `ARTIST` · `ALBUM` · `ALBUMARTIST` · `DATE` · `GENRE` · `TRACKNUMBER` · `DISCNUMBER` · `COMMENT`

## taglib 지원 키 전체 (카테고리별)

**기본/식별**
`TITLE` `ARTIST` `ALBUM` `ALBUMARTIST` `ARTISTS` `DATE` `GENRE` `TRACKNUMBER` `DISCNUMBER` `COMMENT` `COMPILATION` `SUBTITLE` `DISCSUBTITLE` `GROUPING`

**정렬용 (Sort)**
`TITLESORT` `ARTISTSORT` `ALBUMSORT` `ALBUMARTISTSORT` `COMPOSERSORT` `SHOWSORT`

**제작/크레딧**
`COMPOSER` `CONDUCTOR` `ARRANGER` `LYRICIST` `PERFORMER` `PRODUCER` `ENGINEER` `MIXER` `DJMIXER` `REMIXER` `INVOLVEDPEOPLE` `MUSICIANCREDITS`

**클래식/구성 (Work·Movement)**
`WORK` `MOVEMENTNAME` `MOVEMENTNUMBER` `MOVEMENTCOUNT` `SHOWWORKMOVEMENT`

**발매/유통**
`LABEL` `CATALOGNUMBER` `BARCODE` `ISRC` `RELEASEDATE` `RELEASECOUNTRY` `RELEASESTATUS` `RELEASETYPE` `MEDIA` `SCRIPT` `COPYRIGHT` `LICENSE` `OWNER` `ENCODEDBY` `ENCODINGTIME` `TAGGINGDATE`

**음악적 속성**
`BPM` `INITIALKEY` `MOOD` `LANGUAGE` `LENGTH` `LYRICS` `GAPLESSPLAYBACK` `PLAYLISTDELAY` `ENCODING` `FILETYPE`

**원본 (오리지널 출처)**
`ORIGINALALBUM` `ORIGINALARTIST` `ORIGINALLYRICIST` `ORIGINALDATE` `ORIGINALFILENAME`

**URL/웹**
`URL` `ARTISTWEBPAGE` `FILEWEBPAGE` `AUDIOSOURCEWEBPAGE` `COPYRIGHTURL` `PUBLISHERWEBPAGE` `PAYMENTWEBPAGE` `RADIOSTATION` `RADIOSTATIONOWNER` `RADIOSTATIONWEBPAGE`

**외부 DB 식별자 (MusicBrainz / AcoustID 등)**
`MUSICBRAINZ_ALBUMID` `MUSICBRAINZ_ALBUMARTISTID` `MUSICBRAINZ_ARTISTID` `MUSICBRAINZ_RELEASEGROUPID` `MUSICBRAINZ_RELEASETRACKID` `MUSICBRAINZ_TRACKID` `MUSICBRAINZ_WORKID` `ACOUSTID_ID` `ACOUSTID_FINGERPRINT` `ASIN` `MUSICIP_PUID`

**팟캐스트/TV**
`PODCAST` `PODCASTCATEGORY` `PODCASTDESC` `PODCASTID` `PODCASTURL` `TVSHOW` `TVEPISODE` `TVEPISODEID` `TVSEASON` `TVNETWORK`

## 알아둘 점

1. **닫힌 집합이 아니다.** `WriteTags`의 시그니처가 `map[string][]string`이라 **임의의 키**도 쓸 수 있다.
   taglib은 위 키들만 포맷별 표준 필드로 매핑하고, 그 외 키는 그대로(ID3면 `TXXX`, FLAC면 Vorbis 필드) 저장한다.
2. **모든 값은 `[]string`(멀티값)** 이다. 아티스트 여러 명 등 다중 값을 가질 수 있다.
3. **오디오 속성은 태그가 아니다.** 재생 시간·비트레이트·샘플레이트·채널 수는 `ReadProperties`
   (`Length`/`Bitrate`/`SampleRate`/`Channels`)로 따로 읽는다.
4. **앨범 아트(이미지)** 는 바이너리라 이 문자열-맵 API로는 다루지 않는다.

## VocaDB → 태그 매핑

| VocaDB 필드 | 태그 키 |
|---|---|
| `Track.Name` | `TITLE` |
| `Track.Song.ArtistString` | `ARTIST` |
| `Album.Name` | `ALBUM` |
| `Album.ArtistString` | `ALBUMARTIST` |
| `Track.TrackNumber` | `TRACKNUMBER` |
| `Track.DiscNumber` | `DISCNUMBER` |
| `Album.ReleaseDate` | `DATE` 또는 `RELEASEDATE` |
