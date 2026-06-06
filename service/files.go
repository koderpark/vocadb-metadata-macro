package service

import (
	"os"
	"path/filepath"
	"strings"
)

// IsAudioFile은 주어진 경로가 지원하는 음원 포맷인지 확인한다.
func IsAudioFile(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	_, ok := supportedExtensions[ext]
	return ok
}

// ListFiles는 주어진 디렉토리 경로의 최상위에 존재하는 음원 파일의 경로를 반환한다.
func ListFiles(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var files []string
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		path := filepath.Join(dir, e.Name())
		if IsAudioFile(path) {
			files = append(files, path)
		}
	}

	return files, nil
}
