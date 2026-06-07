package service

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
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

// RenameSequential은 주어진 파일들을 슬라이스 순서대로 "01.ext", "02.ext" ...
// 같은 0 패딩 순번 이름으로 바꾼다. 원본(일본어 등) 이름은 버리고 순서와 확장자만 보존하며,
// 바뀐 경로를 같은 순서로 반환한다.
//
// 이름 변경은 2단계(임시 이름 → 최종 이름)로 진행하므로, 최종 이름이 이미 다른 파일이
// 쓰고 있더라도(예: 원본에 "02.flac"이 섞여 있어도) 덮어쓰기 없이 안전하다.
func RenameSequential(files []string) ([]string, error) {
	width := len(strconv.Itoa(len(files)))
	if width < 2 {
		width = 2
	}

	// 1단계: 최종 이름과 절대 겹치지 않는 임시 이름으로 먼저 옮긴다.
	temps := make([]string, len(files))
	for i, f := range files {
		tmp := filepath.Join(filepath.Dir(f), fmt.Sprintf(".vmm-tmp-%d%s", i, filepath.Ext(f)))
		if err := os.Rename(f, tmp); err != nil {
			return nil, fmt.Errorf("renaming %s: %w", f, err)
		}
		temps[i] = tmp
	}

	// 2단계: 임시 이름을 01.ext, 02.ext ... 순번으로 확정한다.
	newPaths := make([]string, len(files))
	for i, tmp := range temps {
		dst := filepath.Join(filepath.Dir(tmp), fmt.Sprintf("%0*d%s", width, i+1, filepath.Ext(tmp)))
		if err := os.Rename(tmp, dst); err != nil {
			return nil, fmt.Errorf("renaming %s to %s: %w", tmp, dst, err)
		}
		newPaths[i] = dst
	}

	return newPaths, nil
}
