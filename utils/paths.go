package utils

import "os"

func GetCurrentWD() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return path, nil
}
