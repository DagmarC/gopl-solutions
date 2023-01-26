package archive

import (
	"errors"
	"fmt"
	"strings"
	"sync"
)

var ErrFormat = errors.New("archive: unknown format")

type format struct {
	name   string
	readFn func(string) error
}

var (
	formatsMu sync.Mutex
	formats   []format
)

func RegisterFormat(name string, readFn func(string) error) {
	formatsMu.Lock()
	frmt := format{name: name, readFn: readFn}
	formats = append(formats, frmt)
	formatsMu.Unlock()
}

func ArchiveReader(fname string) error {
	p := strings.LastIndexAny(fname, ".")
	if p == -1 {
		return errors.New("invalid file")
	}
	fmt.Println("REGISTERED FORMATS: ", formats)

	suffix := fname[p+1:]
	for _, frmt := range formats {
		if suffix == frmt.name {
			fmt.Println("Suffix: ", suffix)
			return frmt.readFn(fname)
		}
	}
	return ErrFormat // format not found
}
