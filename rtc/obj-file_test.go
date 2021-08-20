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

	if got, want := obj.IgnoredLines, 5; got != want {
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

	if got, want := obj.IgnoredLines, 0; got != want {
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

func TestParseObjFile_Faces(t *testing.T) {
	fileData := `
v -1 1 0
v -1 0 0
v 1 0 0
v 1 1 0

f 1 2 3
f 1 3 4
`
	r := bytes.NewBufferString(fileData)
	obj, err := ParseObjFile(r)
	if err != nil {
		t.Fatal(err)
	}

	if got, want := obj.IgnoredLines, 0; got != want {
		t.Errorf("obj.IgnoredLines = %v, want %v", got, want)
	}

	if got, want := len(obj.Vertices), 5; got != want {
		t.Fatalf("len(obj.Vertices) = %v, want %v", got, want)
	}

	if got, want := obj.Vertices[1], Point(-1, 1, 0); !got.Equal(want) {
		t.Errorf("obj.Vertices[1] = %v, want %v", got, want)
	}

	if got, want := obj.Vertices[2], Point(-1, 0, 0); !got.Equal(want) {
		t.Errorf("obj.Vertices[2] = %v, want %v", got, want)
	}

	if got, want := obj.Vertices[3], Point(1, 0, 0); !got.Equal(want) {
		t.Errorf("obj.Vertices[3] = %v, want %v", got, want)
	}

	if got, want := obj.Vertices[4], Point(1, 1, 0); !got.Equal(want) {
		t.Errorf("obj.Vertices[4] = %v, want %v", got, want)
	}

	if got, want := len(obj.DefaultGroup.Children), 2; got != want {
		t.Fatalf("len(obj.DefaultGroup.Children) = %v, want %v", got, want)
	}

	t1, ok := obj.DefaultGroup.Children[0].(*TriangleT)
	if !ok {
		t.Fatalf("obj.DefaultGroup.Children[0] = %T, want *TriangleT", obj.DefaultGroup.Children[0])
	}

	t2, ok := obj.DefaultGroup.Children[1].(*TriangleT)
	if !ok {
		t.Fatalf("obj.DefaultGroup.Children[1] = %T, want *TriangleT", obj.DefaultGroup.Children[1])
	}

	if got, want := t1.P1, obj.Vertices[1]; !got.Equal(want) {
		t.Errorf("t1.P1 = %v, want %v", got, want)
	}

	if got, want := t1.P2, obj.Vertices[2]; !got.Equal(want) {
		t.Errorf("t1.P2 = %v, want %v", got, want)
	}

	if got, want := t1.P3, obj.Vertices[3]; !got.Equal(want) {
		t.Errorf("t1.P3 = %v, want %v", got, want)
	}

	if got, want := t2.P1, obj.Vertices[1]; !got.Equal(want) {
		t.Errorf("t2.P1 = %v, want %v", got, want)
	}

	if got, want := t2.P2, obj.Vertices[3]; !got.Equal(want) {
		t.Errorf("t2.P2 = %v, want %v", got, want)
	}

	if got, want := t2.P3, obj.Vertices[4]; !got.Equal(want) {
		t.Errorf("t2.P3 = %v, want %v", got, want)
	}
}
