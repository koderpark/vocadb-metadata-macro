package vocadb

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// baseURL is the root of the VocaDB public API.
const baseURL = "https://vocadb.net/api"

// client is the HTTP client used for all VocaDB API requests,
// with a timeout to avoid hanging.
var client = &http.Client{Timeout: 10 * time.Second}

// request issues a GET to baseURL+path and returns the raw JSON response body.
// It returns an error on network failure or a non-200 status.
func request(path string) (json.RawMessage, error) {
	req, err := http.NewRequest(http.MethodGet, baseURL+path, nil)
	if err != nil {
		return nil, fmt.Errorf("building request for %s: %w", path, err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("requesting %s: %w", path, err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("vocadb returned status %d for %s", resp.StatusCode, path)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response for %s: %w", path, err)
	}

	return body, nil
}

// albumFields requests the complete album payload from the VocaDB API.
const albumFields = "Artists,Tracks,Discs,MainPicture"

// FetchAlbum retrieves the full album metadata from the VocaDB API for the given album ID.
// It returns an error on network failure, a non-200 status, or invalid JSON.
func FetchAlbum(id int) (*Album, error) {
	raw, err := request(fmt.Sprintf("/albums/%d?fields=%s", id, albumFields))
	if err != nil {
		return nil, fmt.Errorf("fetching album %d: %w", id, err)
	}

	var album Album
	if err := json.Unmarshal(raw, &album); err != nil {
		return nil, fmt.Errorf("decoding album %d: %w", id, err)
	}

	return &album, nil
}
