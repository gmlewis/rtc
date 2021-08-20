package rtc

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// ParseObjFile parses a Wavefront OBJ file and returns an ObjFile.
func ParseObjFile(filename string) (*ObjFile, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	obj, err := ParseObj(f)
	if err != nil {
		return nil, err
	}

	if err := f.Close(); err != nil {
		return nil, err
	}

	return obj, nil
}

// ParseObj parse a Wavefront OBJ and returns an ObjFile.
func ParseObj(r io.Reader) (*ObjFile, error) {
	obj := &ObjFile{
		DefaultGroup: Group(),
		Vertices:     []Tuple{Point(0, 0, 0)},  // Vertex 0 is unused.
		Normals:      []Tuple{Vector(0, 0, 0)}, // Normal 0 is unused.
		NamedGroups:  map[string]*GroupT{},
	}
	lastGroup := obj.DefaultGroup

	addVertex := func(v ...float64) {
		if len(v) != 3 {
			log.Fatalf("expect 3 vertex values, got %#v", v)
		}
		obj.Vertices = append(obj.Vertices, Point(v[0], v[1], v[2]))
	}

	addNormal := func(v ...float64) {
		if len(v) != 3 {
			log.Fatalf("expect 3 normal values, got %#v", v)
		}
		obj.Normals = append(obj.Normals, Vector(v[0], v[1], v[2]))
	}

	addTriangle := func(v ...float64) {
		if len(v) < 3 {
			log.Fatalf("expect 3 or more faces, got %#v", v)
		}
		for i := 2; i < len(v); i++ {
			lastGroup.AddChild(Triangle(obj.Vertices[int(v[0])], obj.Vertices[int(v[i-1])], obj.Vertices[int(v[i])]))
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
		case strings.HasPrefix(line, "vn "):
			parseTriplet(line[3:], addNormal)
		case strings.HasPrefix(line, "f "):
			parseTriplet(line[2:], addTriangle)
		case strings.HasPrefix(line, "g "):
			groupName := strings.TrimSpace(line[2:])
			lastGroup = Group()
			obj.NamedGroups[groupName] = lastGroup
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

	Vertices    []Tuple
	Normals     []Tuple
	NamedGroups map[string]*GroupT
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

// ToGroup returns a GroupT representing the parsed Wavefront OBJ file.
func (o *ObjFile) ToGroup() *GroupT {
	g := Group()
	if len(o.DefaultGroup.Children) > 0 {
		g.AddChild(o.DefaultGroup)
	}
	for _, group := range o.NamedGroups {
		if len(group.Children) > 0 {
			g.AddChild(group)
		}
	}
	return g
}
