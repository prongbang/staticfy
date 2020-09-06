package core

import (
	"fmt"
	"path/filepath"
	"strings"
)

func IsWhiteListExt(extList []string, extFile string) (string, error) {
	if len(extList) > 0 {
		for _, ext := range extList {
			if strings.ToLower(filepath.Ext(ext)) == strings.ToLower(extFile) {
				return ext, nil
			}
		}
	}
	return "", fmt.Errorf("%s %s", "Unsupported extension", extFile)
}
