package vocadb

import (
	"fmt"
	"regexp"
	"strconv"
)

// ParseAlbumID extracts the album primary key from a VocaDB album URL such as
// "https://vocadb.net/Al/53767/whatever" and returns it as an int.
// It returns an error when no album ID can be found or the ID is not a valid int.
func ParseAlbumID(rawURL string) (int, error) {
	m := regexp.MustCompile(`/Al/(\d+)`).FindStringSubmatch(rawURL)
	if m == nil {
		return 0, fmt.Errorf("no album ID found in URL %q", rawURL)
	}

	id, err := strconv.Atoi(m[1])
	if err != nil {
		return 0, fmt.Errorf("invalid album ID %q in URL %q: %w", m[1], rawURL, err)
	}

	return id, nil
}
