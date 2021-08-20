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
	obj := &ObjFile{
		DefaultGroup: Group(),
		Vertices:     []Tuple{Point(0, 0, 0)}, // Vertex 0 is unused.
	}

	addVertex := func(v ...float64) {
		if len(v) != 3 {
			log.Fatalf("expect 3 vertices, got %#v", v)
		}
		obj.Vertices = append(obj.Vertices, Point(v[0], v[1], v[2]))
	}
	addTriangle := func(v ...float64) {
		if len(v) < 3 {
			log.Fatalf("expect 3 or more faces, got %#v", v)
		}
		for i := 2; i < len(v); i++ {
			obj.DefaultGroup.AddChild(Triangle(obj.Vertices[int(v[0])], obj.Vertices[int(v[i-1])], obj.Vertices[int(v[i])]))
		}
	}

	b := bufio.NewReader(r)
	for {
		line, err := b.ReadString('\n')
		if err != nil && err != io.EOF {
			return nil, err
		}

		line = strings.TrimSpace(line)
		switch {
		case line == "": // silently ignore
		case strings.HasPrefix(line, "v "):
			parseTriplet(line[2:], addVertex)
		case strings.HasPrefix(line, "f "):
			parseTriplet(line[2:], addTriangle)
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
	DefaultGroup *GroupT
	IgnoredLines int

	Vertices []Tuple
}

func parseTriplet(s string, f func(args ...float64)) {
	parts := strings.Split(strings.TrimSpace(s), " ")
	var args []float64
	for _, arg := range parts {
		v, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			log.Printf("WARNING: parsing error on line %q, ignored.", s)
			break
		}
		args = append(args, v)
	}
	f(args...)
}
