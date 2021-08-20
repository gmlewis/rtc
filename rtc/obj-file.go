package rtc

import (
	"bufio"
	"io"
	"log"
	"strconv"
	"strings"
)

// ParseObjFile parses a Wavefront OBJ file and returns an Object
// and the number of lines that were ignored.
func ParseObjFile(r io.Reader) (*ObjFile, error) {
	obj := &ObjFile{Vertices: []Tuple{Point(0, 0, 0)}} // Vertex 0 is unused.

	b := bufio.NewReader(r)
	for {
		line, err := b.ReadString('\n')
		if err != nil && err != io.EOF {
			return nil, err
		}

		switch {
		case strings.HasPrefix(line, "v "):
			obj.Vertices = append(obj.Vertices, parseTuple(line[2:], Point))
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

func atof(s string) float64 {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatal(err)
	}
	return v
}

func parseTuple(s string, f func(x, y, z float64) Tuple) Tuple {
	parts := strings.Split(strings.TrimSpace(s), " ")
	x := atof(parts[0])
	y := atof(parts[1])
	z := atof(parts[2])
	return f(x, y, z)
}
