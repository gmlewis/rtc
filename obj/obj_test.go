package obj

import (
	"bytes"
	"testing"

	"github.com/gmlewis/rtc/rtc"
)

func TestParseObj_Gibberish(t *testing.T) {
	gibberish := `
There was a young lady named Bright
who traveled much faster than light.
She set out on day
in a relative way,
and came back the previous night.
`
	r := bytes.NewBufferString(gibberish)
	obj, err := ParseObj(r)
	if err != nil {
		t.Fatal(err)
	}

	if got, want := obj.IgnoredLines, 5; got != want {
		t.Errorf("obj.IgnoredLines = %v, want %v", got, want)
	}
}

func TestParseObj_Vertices(t *testing.T) {
	vertices := `
v -1 1 0
v -1.0000 0.5000 0.0000
v 1 0 0
v 1 1 0
`
	r := bytes.NewBufferString(vertices)
	obj, err := ParseObj(r)
	if err != nil {
		t.Fatal(err)
	}

	if got, want := obj.IgnoredLines, 0; got != want {
		t.Errorf("obj.IgnoredLines = %v, want %v", got, want)
	}

	if got, want := len(obj.Vertices), 5; got != want {
		t.Fatalf("len(obj.Vertices) = %v, want %v", got, want)
	}

	if got, want := obj.Vertices[1], rtc.Point(-1, 1, 0); !got.Equal(want) {
		t.Errorf("obj.Vertices[1] = %v, want %v", got, want)
	}

	if got, want := obj.Vertices[2], rtc.Point(-1, 0.5, 0); !got.Equal(want) {
		t.Errorf("obj.Vertices[2] = %v, want %v", got, want)
	}

	if got, want := obj.Vertices[3], rtc.Point(1, 0, 0); !got.Equal(want) {
		t.Errorf("obj.Vertices[3] = %v, want %v", got, want)
	}

	if got, want := obj.Vertices[4], rtc.Point(1, 1, 0); !got.Equal(want) {
		t.Errorf("obj.Vertices[4] = %v, want %v", got, want)
	}
}

func TestParseObj_Faces(t *testing.T) {
	fileData := `
v -1 1 0
v -1 0 0
v 1 0 0
v 1 1 0

f 1 2 3
f 1 3 4
`
	r := bytes.NewBufferString(fileData)
	obj, err := ParseObj(r)
	if err != nil {
		t.Fatal(err)
	}

	if got, want := obj.IgnoredLines, 0; got != want {
		t.Errorf("obj.IgnoredLines = %v, want %v", got, want)
	}

	if got, want := len(obj.Vertices), 5; got != want {
		t.Fatalf("len(obj.Vertices) = %v, want %v", got, want)
	}

	if got, want := obj.Vertices[1], rtc.Point(-1, 1, 0); !got.Equal(want) {
		t.Errorf("obj.Vertices[1] = %v, want %v", got, want)
	}

	if got, want := obj.Vertices[2], rtc.Point(-1, 0, 0); !got.Equal(want) {
		t.Errorf("obj.Vertices[2] = %v, want %v", got, want)
	}

	if got, want := obj.Vertices[3], rtc.Point(1, 0, 0); !got.Equal(want) {
		t.Errorf("obj.Vertices[3] = %v, want %v", got, want)
	}

	if got, want := obj.Vertices[4], rtc.Point(1, 1, 0); !got.Equal(want) {
		t.Errorf("obj.Vertices[4] = %v, want %v", got, want)
	}

	if got, want := len(obj.DefaultGroup.Children), 2; got != want {
		t.Fatalf("len(obj.DefaultGroup.Children) = %v, want %v", got, want)
	}

	t1, ok := obj.DefaultGroup.Children[0].(*rtc.TriangleT)
	if !ok {
		t.Fatalf("obj.DefaultGroup.Children[0] = %T, want *TriangleT", obj.DefaultGroup.Children[0])
	}

	t2, ok := obj.DefaultGroup.Children[1].(*rtc.TriangleT)
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

func TestParseObj_Polygon(t *testing.T) {
	fileData := `
v -1 1 0
v -1 0 0
v 1 0 0
v 1 1 0
v 0 2 0

f 1 2 3 4 5
`
	r := bytes.NewBufferString(fileData)
	obj, err := ParseObj(r)
	if err != nil {
		t.Fatal(err)
	}

	if got, want := obj.IgnoredLines, 0; got != want {
		t.Errorf("obj.IgnoredLines = %v, want %v", got, want)
	}

	if got, want := len(obj.Vertices), 6; got != want {
		t.Fatalf("len(obj.Vertices) = %v, want %v", got, want)
	}

	if got, want := obj.Vertices[1], rtc.Point(-1, 1, 0); !got.Equal(want) {
		t.Errorf("obj.Vertices[1] = %v, want %v", got, want)
	}

	if got, want := obj.Vertices[2], rtc.Point(-1, 0, 0); !got.Equal(want) {
		t.Errorf("obj.Vertices[2] = %v, want %v", got, want)
	}

	if got, want := obj.Vertices[3], rtc.Point(1, 0, 0); !got.Equal(want) {
		t.Errorf("obj.Vertices[3] = %v, want %v", got, want)
	}

	if got, want := obj.Vertices[4], rtc.Point(1, 1, 0); !got.Equal(want) {
		t.Errorf("obj.Vertices[4] = %v, want %v", got, want)
	}

	if got, want := len(obj.DefaultGroup.Children), 3; got != want {
		t.Fatalf("len(obj.DefaultGroup.Children) = %v, want %v", got, want)
	}

	t1, ok := obj.DefaultGroup.Children[0].(*rtc.TriangleT)
	if !ok {
		t.Fatalf("obj.DefaultGroup.Children[0] = %T, want *TriangleT", obj.DefaultGroup.Children[0])
	}

	t2, ok := obj.DefaultGroup.Children[1].(*rtc.TriangleT)
	if !ok {
		t.Fatalf("obj.DefaultGroup.Children[1] = %T, want *TriangleT", obj.DefaultGroup.Children[1])
	}

	t3, ok := obj.DefaultGroup.Children[2].(*rtc.TriangleT)
	if !ok {
		t.Fatalf("obj.DefaultGroup.Children[2] = %T, want *TriangleT", obj.DefaultGroup.Children[2])
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

	if got, want := t3.P1, obj.Vertices[1]; !got.Equal(want) {
		t.Errorf("t3.P1 = %v, want %v", got, want)
	}

	if got, want := t3.P2, obj.Vertices[4]; !got.Equal(want) {
		t.Errorf("t3.P2 = %v, want %v", got, want)
	}

	if got, want := t3.P3, obj.Vertices[5]; !got.Equal(want) {
		t.Errorf("t3.P3 = %v, want %v", got, want)
	}
}

var trianglesObjFileData = `
v -1 1 0
v -1 0 0
v 1 0 0
v 1 1 0

g FirstGroup
f 1 2 3
g SecondGroup
f 1 3 4
`

func TestParseObj_NamedGroups(t *testing.T) {
	r := bytes.NewBufferString(trianglesObjFileData)
	obj, err := ParseObj(r)
	if err != nil {
		t.Fatal(err)
	}

	if got, want := obj.IgnoredLines, 0; got != want {
		t.Errorf("obj.IgnoredLines = %v, want %v", got, want)
	}

	if got, want := len(obj.Vertices), 5; got != want {
		t.Fatalf("len(obj.Vertices) = %v, want %v", got, want)
	}

	if got, want := obj.Vertices[1], rtc.Point(-1, 1, 0); !got.Equal(want) {
		t.Errorf("obj.Vertices[1] = %v, want %v", got, want)
	}

	if got, want := obj.Vertices[2], rtc.Point(-1, 0, 0); !got.Equal(want) {
		t.Errorf("obj.Vertices[2] = %v, want %v", got, want)
	}

	if got, want := obj.Vertices[3], rtc.Point(1, 0, 0); !got.Equal(want) {
		t.Errorf("obj.Vertices[3] = %v, want %v", got, want)
	}

	if got, want := obj.Vertices[4], rtc.Point(1, 1, 0); !got.Equal(want) {
		t.Errorf("obj.Vertices[4] = %v, want %v", got, want)
	}

	if got, want := len(obj.DefaultGroup.Children), 0; got != want {
		t.Fatalf("len(obj.DefaultGroup.Children) = %v, want %v", got, want)
	}

	if got, want := len(obj.NamedGroups), 2; got != want {
		t.Fatalf("len(obj.NamedGroups) = %v, want %v", got, want)
	}

	g1, ok := obj.NamedGroups["FirstGroup"]
	if !ok {
		t.Fatalf("obj.NamedGroups['FirstGroup'] not found")
	}

	if got, want := len(g1.Children), 1; got != want {
		t.Fatalf("len(g1.Children) = %v, want %v", got, want)
	}

	g2, ok := obj.NamedGroups["SecondGroup"]
	if !ok {
		t.Fatalf("obj.NamedGroups['SecondGroup'] not found")
	}

	if got, want := len(g2.Children), 1; got != want {
		t.Fatalf("len(g2.Children) = %v, want %v", got, want)
	}

	t1, ok := g1.Children[0].(*rtc.TriangleT)
	if !ok {
		t.Fatalf("g1.Children[0] = %T, want *TriangleT", obj.DefaultGroup.Children[0])
	}

	t2, ok := g2.Children[0].(*rtc.TriangleT)
	if !ok {
		t.Fatalf("g2.Children[0] = %T, want *TriangleT", obj.DefaultGroup.Children[1])
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

func TestParseObj_VertexNormals(t *testing.T) {
	vertices := `
vn 0 0 1
vn 0.707 0 -0.707
vn 1 2 3
`
	r := bytes.NewBufferString(vertices)
	obj, err := ParseObj(r)
	if err != nil {
		t.Fatal(err)
	}

	if got, want := obj.IgnoredLines, 0; got != want {
		t.Errorf("obj.IgnoredLines = %v, want %v", got, want)
	}

	if got, want := len(obj.Vertices), 1; got != want {
		t.Fatalf("len(obj.Vertices) = %v, want %v", got, want)
	}

	if got, want := len(obj.Normals), 4; got != want {
		t.Fatalf("len(obj.Normals) = %v, want %v", got, want)
	}

	if got, want := obj.Normals[1], rtc.Vector(0, 0, 1); !got.Equal(want) {
		t.Errorf("obj.Normals[1] = %v, want %v", got, want)
	}

	if got, want := obj.Normals[2], rtc.Vector(0.707, 0, -0.707); !got.Equal(want) {
		t.Errorf("obj.Normals[2] = %v, want %v", got, want)
	}

	if got, want := obj.Normals[3], rtc.Vector(1, 2, 3); !got.Equal(want) {
		t.Errorf("obj.Normals[3] = %v, want %v", got, want)
	}
}

func TestParseObj_FacesWithNormals(t *testing.T) {
	vertices := `
v 0 1 0
v -1 0 0
v 1 0 0

vn -1 0 0
vn 1 0 0
vn 0 1 0

f 1//3 2//1 3//2
f 1/0/3 2/102/1 3/14/2
`
	r := bytes.NewBufferString(vertices)
	obj, err := ParseObj(r)
	if err != nil {
		t.Fatal(err)
	}

	if got, want := obj.IgnoredLines, 0; got != want {
		t.Errorf("obj.IgnoredLines = %v, want %v", got, want)
	}

	if got, want := len(obj.Vertices), 4; got != want {
		t.Fatalf("len(obj.Vertices) = %v, want %v", got, want)
	}

	if got, want := len(obj.Normals), 4; got != want {
		t.Fatalf("len(obj.Normals) = %v, want %v", got, want)
	}

	if got, want := len(obj.DefaultGroup.Children), 2; got != want {
		t.Fatalf("len(obj.DefaultGroup.Children) = %v, want %v", got, want)
	}

	t1, ok := obj.DefaultGroup.Children[0].(*rtc.SmoothTriangleT)
	if !ok {
		t.Fatalf("obj.DefaultGroup.Children[0] = %T, want *SmoothTriangleT", obj.DefaultGroup.Children[0])
	}

	t2, ok := obj.DefaultGroup.Children[1].(*rtc.SmoothTriangleT)
	if !ok {
		t.Fatalf("obj.DefaultGroup.Children[1] = %T, want *SmoothTriangleT", obj.DefaultGroup.Children[1])
	}

	if got, want := t1.P1, obj.Vertices[1]; got != want {
		t.Errorf("t1.P1 = %v, want %v", got, want)
	}

	if got, want := t1.P2, obj.Vertices[2]; got != want {
		t.Errorf("t1.P2 = %v, want %v", got, want)
	}

	if got, want := t1.P3, obj.Vertices[3]; got != want {
		t.Errorf("t1.P3 = %v, want %v", got, want)
	}

	if got, want := t1.N1, obj.Normals[3]; got != want {
		t.Errorf("t1.N1 = %v, want %v", got, want)
	}

	if got, want := t1.N2, obj.Normals[1]; got != want {
		t.Errorf("t1.N2 = %v, want %v", got, want)
	}

	if got, want := t1.N3, obj.Normals[2]; got != want {
		t.Errorf("t1.N3 = %v, want %v", got, want)
	}

	if got, want := t2.P1, obj.Vertices[1]; got != want {
		t.Errorf("t2.P1 = %v, want %v", got, want)
	}

	if got, want := t2.P2, obj.Vertices[2]; got != want {
		t.Errorf("t2.P2 = %v, want %v", got, want)
	}

	if got, want := t2.P3, obj.Vertices[3]; got != want {
		t.Errorf("t2.P3 = %v, want %v", got, want)
	}

	if got, want := t2.N1, obj.Normals[3]; got != want {
		t.Errorf("t2.N1 = %v, want %v", got, want)
	}

	if got, want := t2.N2, obj.Normals[1]; got != want {
		t.Errorf("t2.N2 = %v, want %v", got, want)
	}

	if got, want := t2.N3, obj.Normals[2]; got != want {
		t.Errorf("t2.N3 = %v, want %v", got, want)
	}
}
