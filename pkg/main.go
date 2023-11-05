package pkg

import "os"

func Contains(way string, searchIt string) (bool, error) {
	content, err := os.ReadFile(way)
	if err != nil {
		return false, err
	}
	return searcher(string(content), searchIt), err
}

func searcher(content string, searchIt string) bool {
	resultAssist := false
	for i := 0; i < len(content); i++ {
		if content[i] == searchIt[0] && len(searchIt) <= len(content) {
			k := i + 1
			for j := 2; j < len(searchIt); j++ {
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
