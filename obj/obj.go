// Package obj implements Wavefront OBJ file parsing.
package obj

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gmlewis/rtc/rtc"
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

func getFaceIndices(s string) (int, int) {
	var p, n int
	parseArgs(s, "/", func(args ...interface{}) {
		if len(args) != 3 {
			log.Fatalf("expect 3 parts to face argument, got %#v", args)
		}
		p = int(args[0].(float64))
		n = int(args[2].(float64))
	})
	return p, n
}

func get3Floats(v ...interface{}) (float64, float64, float64) {
	if len(v) != 3 {
		log.Fatalf("expect 3 floats, got %#v", v)
	}
	x, y, z := v[0].(float64), v[1].(float64), v[2].(float64)
	return x, y, z
}

// ParseObj parse a Wavefront OBJ and returns an ObjFile.
func ParseObj(r io.Reader) (*ObjFile, error) {
	obj := &ObjFile{
		DefaultGroup: rtc.Group(),
		Vertices:     []rtc.Tuple{rtc.Point(0, 0, 0)},  // Vertex 0 is unused.
		Normals:      []rtc.Tuple{rtc.Vector(0, 0, 0)}, // Normal 0 is unused.
		NamedGroups:  map[string]*rtc.GroupT{},
	}
	lastGroup := obj.DefaultGroup

	addVertex := func(v ...interface{}) {
		obj.Vertices = append(obj.Vertices, rtc.Point(get3Floats(v...)))
	}

	addNormal := func(v ...interface{}) {
		obj.Normals = append(obj.Normals, rtc.Vector(get3Floats(v...)))
	}

	addTriangle := func(v ...interface{}) {
		if len(v) < 3 {
			log.Fatalf("expect 3 or more face arguments, got %#v", v)
		}
		if _, ok := v[0].(float64); ok {
			p1 := int(v[0].(float64))
			for i := 2; i < len(v); i++ {
				p2, p3 := int(v[i-1].(float64)), int(v[i].(float64))
				lastGroup.AddChild(rtc.Triangle(obj.Vertices[p1], obj.Vertices[p2], obj.Vertices[p3]))
			}
			return
		}

		p1, n1 := getFaceIndices(v[0].(string))
		for i := 2; i < len(v); i++ {
			p2, n2 := getFaceIndices(v[i-1].(string))
			p3, n3 := getFaceIndices(v[i].(string))
			lastGroup.AddChild(rtc.SmoothTriangle(
				obj.Vertices[p1], obj.Vertices[p2], obj.Vertices[p3],
				obj.Normals[n1], obj.Normals[n2], obj.Normals[n3],
			))
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
			parseArgs(line[2:], " ", addVertex)
		case strings.HasPrefix(line, "vn "):
			parseArgs(line[3:], " ", addNormal)
		case strings.HasPrefix(line, "f "):
			parseArgs(line[2:], " ", addTriangle)
		case strings.HasPrefix(line, "g "):
			groupName := strings.TrimSpace(line[2:])
			lastGroup = rtc.Group()
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
	DefaultGroup *rtc.GroupT
	IgnoredLines int

	Vertices    []rtc.Tuple
	Normals     []rtc.Tuple
	NamedGroups map[string]*rtc.GroupT
}

func parseArgs(s, sep string, f func(args ...interface{})) {
	parts := strings.Split(strings.TrimSpace(s), sep)
	var args []interface{}
	for _, arg := range parts {
		v, err := strconv.ParseFloat(arg, 64)
		if err == nil {
			args = append(args, v)
		} else {
			args = append(args, arg)
		}
	}
	f(args...)
}

// ToGroup returns a GroupT representing the parsed Wavefront OBJ file.
func (o *ObjFile) ToGroup() *rtc.GroupT {
	g := rtc.Group()
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
