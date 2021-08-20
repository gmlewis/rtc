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

	addVertex := func(x, y, z float64) {
		obj.Vertices = append(obj.Vertices, Point(x, y, z))
	}
	addTriangle := func(x, y, z float64) {
		obj.DefaultGroup.AddChild(Triangle(obj.Vertices[int(x)], obj.Vertices[int(y)], obj.Vertices[int(z)]))
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

func atof(s string) float64 {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatal(err)
	}
	return v
}

func parseTriplet(s string, f func(x, y, z float64)) {
	parts := strings.Split(strings.TrimSpace(s), " ")
	x := atof(parts[0])
	y := atof(parts[1])
	z := atof(parts[2])
	f(x, y, z)
}
