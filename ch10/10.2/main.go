package main

import (
	"log"

	"github.com/DagmarC/gopl-solutions/ch10/10.2/archive"
	_ "github.com/DagmarC/gopl-solutions/ch10/10.2/archive/zip"
	_ "github.com/DagmarC/gopl-solutions/ch10/10.2/archive/tar"


)

func main() {
	err := archive.ArchiveReader("/Users/dagmarmac/go/src/github.com/DagmarC/gopl-solutions/ch10/test.zip")
	if err != nil {
		log.Fatal(err)
	}
}