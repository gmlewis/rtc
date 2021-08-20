package rtc

import (
	"bytes"
	"testing"
)

func TestParseObjFile_Gibberish(t *testing.T) {
	gibberish := `
There was a young lady named Bright
who traveled much faster than light.
She set out on day
in a relative way,
and came back the previous night.
`
	r := bytes.NewBufferString(gibberish)
	obj, err := ParseObjFile(r)
	if err != nil {
		t.Fatal(err)
	}

	if got, want := obj.IgnoredLines, 7; got != want {
		t.Errorf("obj.IgnoredLines = %v, want %v", got, want)
	}
}

func TestParseObjFile_Vertices(t *testing.T) {
	vertices := `
v -1 1 0
v -1.0000 0.5000 0.0000
v 1 0 0
v 1 1 0
`
	r := bytes.NewBufferString(vertices)
	obj, err := ParseObjFile(r)
	if err != nil {
		t.Fatal(err)
	}

	if got, want := obj.IgnoredLines, 2; got != want {
		t.Errorf("obj.IgnoredLines = %v, want %v", got, want)
	}

	if got, want := len(obj.Vertices), 5; got != want {
		t.Fatalf("len(obj.Vertices) = %v, want %v", got, want)
	}

	if got, want := obj.Vertices[1], Point(-1, 1, 0); !got.Equal(want) {
		t.Errorf("obj.Vertices[1] = %v, want %v", got, want)
	}

	if got, want := obj.Vertices[2], Point(-1, 0.5, 0); !got.Equal(want) {
		t.Errorf("obj.Vertices[2] = %v, want %v", got, want)
	}

	if got, want := obj.Vertices[3], Point(1, 0, 0); !got.Equal(want) {
		t.Errorf("obj.Vertices[3] = %v, want %v", got, want)
	}

	if got, want := obj.Vertices[4], Point(1, 1, 0); !got.Equal(want) {
		t.Errorf("obj.Vertices[4] = %v, want %v", got, want)
	}
}
