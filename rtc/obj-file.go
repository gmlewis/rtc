package rtc

import (
	"bufio"
	"io"
	"strings"
)

// ParseObjFile parses a Wavefront OBJ file and returns an Object
// and the number of lines that were ignored.
func ParseObjFile(r io.Reader) (*ObjFile, error) {
	obj := &ObjFile{}

	b := bufio.NewReader(r)
	for {
		line, err := b.ReadString('\n')
		if err != nil && err != io.EOF {
			return nil, err
		}

		switch {
		case strings.HasPrefix(line, "v "):
		default:
			obj.IgnoredLines++
		}

		if err == io.EOF {
			break
		}
	}
	return obj, nil
}

// ObjFile represents a parsed Wavefront OBJ file.
type ObjFile struct {
	Group        *GroupT
	IgnoredLines int

	Vertices []Tuple
}
