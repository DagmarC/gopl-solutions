package html

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/DagmarC/gopl-solutions/utils"
)

//!+sema
// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests - HTTP and file/folder cretion.
var tokens = make(chan struct{}, 10)

func SavePage(path string, dst string) error {

	url, err := GetURL(path)
	if err != nil {
		return err
	}

	fpath, dirPath := getFileDirPath(url, url.Host, dst)

	tokens <- struct{}{} // acquire a token
	resp, err := http.Get(path)

	os.MkdirAll(dirPath, os.ModePerm)

	f, err := os.Create(fpath)
	if err != nil {
		return err
	}
	<-tokens // release a token

	defer f.Close()

	io.Copy(f, resp.Body)

	return nil
}

func getFileDirPath(url *url.URL, host, dst string) (string, string) {
	wd, err := utils.GetCurrentWD()
	if err != nil {
		log.Fatal(err)
	}

	var fpath string
	if filepath.Ext(url.Path) == "" {
		fpath = filepath.Join(wd, dst, host, url.Path, "index.html")
	} else {
		fpath = filepath.Join(wd, dst, host, url.Path)
	}

	dirPath := filepath.Dir(fpath)

	return fpath, dirPath
}
