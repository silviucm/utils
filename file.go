package utils

import "os"

// Checks if the file with the given path exists, returns true if yes
func FileExists(name string) bool {

	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func FileDelete(filePath string) error {

	return os.Remove(filePath)

}
