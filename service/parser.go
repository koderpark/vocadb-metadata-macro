package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

// albumIDPattern matches the album primary key in a VocaDB album URL,
// e.g. the "53767" in "https://vocadb.net/Al/53767/...".
var albumIDPattern = regexp.MustCompile(`/Al/(\d+)`)

// ParseAlbumID extracts the album primary key from a VocaDB album URL such as
// "https://vocadb.net/Al/53767/whatever" and returns it as an int.
// It returns an error when no album ID can be found or the ID is not a valid int.
func ParseAlbumID(rawURL string) (int, error) {
	m := albumIDPattern.FindStringSubmatch(rawURL)
	if m == nil {
		return 0, fmt.Errorf("no album ID found in URL %q", rawURL)
	}

	id, err := strconv.Atoi(m[1])
	if err != nil {
		return 0, fmt.Errorf("invalid album ID %q in URL %q: %w", m[1], rawURL, err)
	}

	return id, nil
}

// vocadbClient is used for VocaDB API requests, with a timeout to avoid hanging.
var vocadbClient = &http.Client{Timeout: 10 * time.Second}

// Album is the full album metadata returned by
// GET https://vocadb.net/api/albums/{id}?fields=... .
type Album struct {
	ID                  int             `json:"id"`
	Name                string          `json:"name"`
	DefaultName         string          `json:"defaultName"`
	DefaultNameLanguage string          `json:"defaultNameLanguage"`
	ArtistString        string          `json:"artistString"`
	Description         string          `json:"description"`
	DiscType            string          `json:"discType"`
	Status              string          `json:"status"`
	CreateDate          string          `json:"createDate"`
	RatingAverage       float64         `json:"ratingAverage"`
	RatingCount         int             `json:"ratingCount"`
	Version             int             `json:"version"`
	ReleaseDate         ReleaseDate     `json:"releaseDate"`
	MainPicture         Picture         `json:"mainPicture"`
	Names               []LocalizedName `json:"names"`
	Artists             []AlbumArtist   `json:"artists"`
	Discs               []Disc          `json:"discs"`
	Tracks              []Track         `json:"tracks"`
	Tags                []TagUsage      `json:"tags"`
	PVs                 []PV            `json:"pvs"`
	WebLinks            []WebLink       `json:"webLinks"`
	ReleaseEvents       []ReleaseEvent  `json:"releaseEvents"`
	Identifiers         []string        `json:"identifiers"`
}

// ReleaseDate is the release date object embedded in an Album.
type ReleaseDate struct {
	Year    int  `json:"year"`
	Month   int  `json:"month"`
	Day     int  `json:"day"`
	IsEmpty bool `json:"isEmpty"`
}

// Picture holds the URLs and MIME type of an album picture.
type Picture struct {
	Mime          string `json:"mime"`
	URLOriginal   string `json:"urlOriginal"`
	URLSmallThumb string `json:"urlSmallThumb"`
	URLThumb      string `json:"urlThumb"`
	URLTinyThumb  string `json:"urlTinyThumb"`
}

// LocalizedName is a name in a specific language.
type LocalizedName struct {
	Language string `json:"language"`
	Value    string `json:"value"`
}

// AlbumArtist links an Artist to the album along with its roles.
type AlbumArtist struct {
	Artist         Artist `json:"artist"`
	Name           string `json:"name"`
	Categories     string `json:"categories"`
	EffectiveRoles string `json:"effectiveRoles"`
	Roles          string `json:"roles"`
	IsSupport      bool   `json:"isSupport"`
}

// Artist is the artist record referenced by an album or song.
type Artist struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	AdditionalNames string `json:"additionalNames"`
	ArtistType      string `json:"artistType"`
	PictureMime     string `json:"pictureMime"`
	ReleaseDate     string `json:"releaseDate"`
	Status          string `json:"status"`
	Version         int    `json:"version"`
	Deleted         bool   `json:"deleted"`
}

// Disc describes a single disc of a multi-disc album.
type Disc struct {
	DiscNumber int    `json:"discNumber"`
	ID         int    `json:"id"`
	MediaType  string `json:"mediaType"`
	Name       string `json:"name"`
}

// Track is a single track on the album.
type Track struct {
	ID                   int      `json:"id"`
	Name                 string   `json:"name"`
	DiscNumber           int      `json:"discNumber"`
	TrackNumber          int      `json:"trackNumber"`
	Song                 Song     `json:"song"`
	ComputedCultureCodes []string `json:"computedCultureCodes"`
}

// Song is the song referenced by a track.
type Song struct {
	ID                  int      `json:"id"`
	Name                string   `json:"name"`
	DefaultName         string   `json:"defaultName"`
	DefaultNameLanguage string   `json:"defaultNameLanguage"`
	ArtistString        string   `json:"artistString"`
	SongType            string   `json:"songType"`
	LengthSeconds       int      `json:"lengthSeconds"`
	FavoritedTimes      int      `json:"favoritedTimes"`
	RatingScore         int      `json:"ratingScore"`
	PVServices          string   `json:"pvServices"`
	PublishDate         string   `json:"publishDate"`
	CreateDate          string   `json:"createDate"`
	Status              string   `json:"status"`
	Version             int      `json:"version"`
	CultureCodes        []string `json:"cultureCodes"`
}

// TagUsage is a tag applied to the album together with its usage count.
type TagUsage struct {
	Count int `json:"count"`
	Tag   Tag `json:"tag"`
}

// Tag is a VocaDB tag.
type Tag struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	AdditionalNames string `json:"additionalNames"`
	CategoryName    string `json:"categoryName"`
	URLSlug         string `json:"urlSlug"`
}

// PV is a promotional video/audio link associated with the album.
type PV struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Service     string `json:"service"`
	PVID        string `json:"pvId"`
	PVType      string `json:"pvType"`
	URL         string `json:"url"`
	Length      int    `json:"length"`
	PublishDate string `json:"publishDate"`
	Disabled    bool   `json:"disabled"`
}

// WebLink is an external link associated with the album.
type WebLink struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Category    string `json:"category"`
	Disabled    bool   `json:"disabled"`
}

// ReleaseEvent is an event at which the album was released.
type ReleaseEvent struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Category     string `json:"category"`
	Date         string `json:"date"`
	EndDate      string `json:"endDate"`
	SeriesID     int    `json:"seriesId"`
	SeriesNumber int    `json:"seriesNumber"`
	SeriesSuffix string `json:"seriesSuffix"`
	Status       string `json:"status"`
	URLSlug      string `json:"urlSlug"`
	VenueName    string `json:"venueName"`
	Version      int    `json:"version"`
}

// albumFields requests the complete album payload from the VocaDB API.
const albumFields = "Artists,Names,PVs,Tags,Tracks,WebLinks,Description,Discs,Identifiers,MainPicture,ReleaseEvent"

// FetchAlbum retrieves the full album metadata from the VocaDB API for the given album ID.
// It returns an error on network failure, a non-200 status, or invalid JSON.
func FetchAlbum(id int) (*Album, error) {
	url := fmt.Sprintf("https://vocadb.net/api/albums/%d?fields=%s&lang=Default", id, albumFields)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("building request for album %d: %w", id, err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := vocadbClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("requesting album %d: %w", id, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("vocadb returned status %d for album %d", resp.StatusCode, id)
	}

	var album Album
	if err := json.NewDecoder(resp.Body).Decode(&album); err != nil {
		return nil, fmt.Errorf("decoding album %d: %w", id, err)
	}

	return &album, nil
}
