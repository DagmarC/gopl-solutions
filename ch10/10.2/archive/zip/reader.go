package zip

import (
	"archive/zip"
	"fmt"
	"io/ioutil"

	"github.com/DagmarC/gopl-solutions/ch10/10.2/archive"
)

// Register zip, ReadZip function to the Archive reader
func init() {
	archive.RegisterFormat("zip", ReadZip)
}

func ReadZip(f string) error {

	zf, err := zip.OpenReader(f)
	if err != nil {
		return err
	}
	defer zf.Close()

	for _, file := range zf.File {
		fmt.Printf("=%s\n", file.Name)
		res, err := readAll(file)
		if err != nil {
			return err
		}
		fmt.Printf("%s\n\n", res) // file content
	}
	return nil
}

// readAll is a wrapper function for ioutil.ReadAll. It accepts a zip.File as
// its parameter, opens it, reads its content and returns it as a byte slice.
func readAll(file *zip.File) ([]byte, error) {
	fc, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer fc.Close()

	cnt, err := ioutil.ReadAll(fc)
	if err != nil {
		return nil, err
	}

	return cnt, nil
}
