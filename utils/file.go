package utils

import (
	"log"
	"os"
)

func GetFile(name string) *os.File {
    f, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		log.Fatal(err)
	}
	return f
}