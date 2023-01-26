package tar

import (
	"archive/tar"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/DagmarC/gopl-solutions/ch10/10.2/archive"
)

// Register zip, ReadZip function to the Archive reader
func init() {
	archive.RegisterFormat("tar", ReadTar)
}

func ReadTar(f string) error {

	// Open the tar archive for reading.
	file, err := os.Open(f)
	if err != nil {
		return err
	}
	defer file.Close()

	tr := tar.NewReader(file)
	// read the complete content of the file h.Name into the bs []byte
	bs, err := ioutil.ReadAll(tr)
	if err != nil {
		return err
	}

	fmt.Println("========Reading TAR file=======")
	// convert the []byte to a string
	fmt.Println(string(bs))
	return nil
}
