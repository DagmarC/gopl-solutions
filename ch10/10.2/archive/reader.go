package archive

import (
	"errors"
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

func ArchiveReader(name string) error {
	p := strings.LastIndexAny(name, ".")
	if p == -1 {
		return errors.New("invalid file")
	}
	suffix := name[p+1:]

	for _, frmt := range formats {
		if suffix == frmt.name {
			return frmt.readFn(frmt.name)
		}
	}
	return ErrFormat // format not found
}
