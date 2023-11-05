package pkg

import (
	"errors"
	"os"
)

func Contains(way string, searchIt string) (bool, error) {
	content, err := os.ReadFile(way)
	if err != nil {
		return false, errors.New("problems with the path")
	}
	if len(searchIt) < 1 {
		return false, errors.New("empty character")
	}
	return searcher(string(content), searchIt), err
}

func searcher(content string, searchIt string) bool {
	resultAssist := false
	for i := 0; i < len(content); i++ {
		if content[i] == searchIt[0] && len(searchIt) <= len(content) {
			resultAssist = true
			k := i + 1
			for j := 1; j < len(searchIt); j++ {
				if content[k] != searchIt[j] {
					resultAssist = false
					break
				}
				k++
			}
			if resultAssist == true {
				return true
			}
		}
	}
	return false
}
