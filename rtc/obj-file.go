package rtc

import "io"

// ParseObjFile parses a Wavefront OBJ file and returns an Object
// and the number of lines that were ignored.
func ParseObjFile(r io.Reader) *ObjFile {
	return &ObjFile{}
}

// ObjFile represents a parsed Wavefront OBJ file.
type ObjFile struct {
	Group        *GroupT
	IgnoredLines int

	Vertices []Tuple
}
