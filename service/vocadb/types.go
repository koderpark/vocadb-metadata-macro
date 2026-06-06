package vocadb

// Album is the full album metadata returned by
// GET https://vocadb.net/api/albums/{id}?fields=... .
type Album struct {
	ID            int           `json:"id"`
	Name          string        `json:"name"`
	ArtistString  string        `json:"artistString"`
	DiscType      string        `json:"discType"`
	Status        string        `json:"status"`
	CreateDate    string        `json:"createDate"`
	RatingAverage float64       `json:"ratingAverage"`
	RatingCount   int           `json:"ratingCount"`
	Version       int           `json:"version"`
	ReleaseDate   ReleaseDate   `json:"releaseDate"`
	MainPicture   Picture       `json:"mainPicture"`
	Artists       []AlbumArtist `json:"artists"`
	Discs         []Disc        `json:"discs"`
	Tracks        []Track       `json:"tracks"`
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
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ArtistType  string `json:"artistType"`
	PictureMime string `json:"pictureMime"`
	ReleaseDate string `json:"releaseDate"`
	Status      string `json:"status"`
	Version     int    `json:"version"`
	Deleted     bool   `json:"deleted"`
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
	ID             int      `json:"id"`
	Name           string   `json:"name"`
	ArtistString   string   `json:"artistString"`
	SongType       string   `json:"songType"`
	LengthSeconds  int      `json:"lengthSeconds"`
	FavoritedTimes int      `json:"favoritedTimes"`
	RatingScore    int      `json:"ratingScore"`
	PVServices     string   `json:"pvServices"`
	PublishDate    string   `json:"publishDate"`
	CreateDate     string   `json:"createDate"`
	Status         string   `json:"status"`
	Version        int      `json:"version"`
	CultureCodes   []string `json:"cultureCodes"`
}
