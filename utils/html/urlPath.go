package html

import (
	"log"
	"net/url"
)

func GetURL(path string) (*url.URL, error){
	url, err := url.Parse(path)
	if err != nil {
		return nil, err
	}
	return url, nil
}

func GetHost(path string) string{
	url, err := GetURL(path)
	if err != nil {
		log.Fatal(err)
	}
	return url.Host
}

func GetPath(path string) string{
	url, err := GetURL(path)
	if err != nil {
		log.Fatal(err)
	}
	return url.Path
}