package rtc

import (
	"bytes"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParseYAML_ToYAML(t *testing.T) {
	r := bytes.NewBufferString(coverYAML)
	y, err := ParseYAML(r)
	if err != nil {
		t.Fatal(err)
	}

	buf, err := y.ToYAML()
	if err != nil {
		t.Fatal(err)
	}

	gotLines := strings.Split(string(buf), "\n")
	wantLines := strings.Split(coverYAML, "\n")

	if got, want := len(gotLines), len(wantLines); got != want {
		t.Errorf("len(gotLines) = %v, want %v", got, want)
	}

	for i := 0; i < len(gotLines) && i < len(wantLines); i++ {
		if got, want := gotLines[i], wantLines[i]; got != want {
			t.Errorf("gotLines[%v] =\n%v\n, want\n%v", i+1, got, want)
		}
	}
}

func TestParseYAML(t *testing.T) {
	r := bytes.NewBufferString(coverYAML)
	y, err := ParseYAML(r)
	if err != nil {
		t.Fatal(err)
	}

	if got, want := len(y.Items), 30; got != want {
		t.Fatalf("len(y.Items) = %v, want %v", got, want)
	}

	I := func(i int) *int { return &i }
	F := func(f float64) *float64 { return &f }
	S := func(s string) *string { return &s }

	want := []Item{
		{Add: S("camera"), Width: I(100), Height: I(99), FOV: F(0.785), From: []float64{-6, 6, -10}, To: []float64{6, 0, 6}, Up: []float64{-0.45, 1, 0}},
		{Add: S("light"), At: []float64{50, 100, -50}, Intensity: []float64{1, 1, 1}},
		{Add: S("light"), At: []float64{-400, 50, -10}, Intensity: []float64{0.2, 0.2, 0.2}},
		{Define: S("white-material"), ValueMaterial: &YAMLMaterial{Color: []float64{1, 1, 1}, Diffuse: F(0.7), Ambient: F(0.1), Specular: F(0), Reflective: F(0.1)}},
		{Define: S("blue-material"), Extend: S("white-material"), ValueMaterial: &YAMLMaterial{Color: []float64{0.537, 0.831, 0.914}}},
		{Define: S("red-material"), Extend: S("white-material"), ValueMaterial: &YAMLMaterial{Color: []float64{0.941, 0.322, 0.388}}},
		{Define: S("purple-material"), Extend: S("white-material"), ValueMaterial: &YAMLMaterial{Color: []float64{0.373, 0.404, 0.55}}},
		{Define: S("standard-transform")},
		{Define: S("large-object")},
		{Define: S("medium-object")},
		{Define: S("small-object")},
		{Add: S("plane")},
		{Add: S("sphere")},
		{Add: S("cube")},
		{Add: S("cube")},
		{Add: S("cube")},
		{Add: S("cube")},
		{Add: S("cube")},
		{Add: S("cube")},
		{Add: S("cube")},
		{Add: S("cube")},
		{Add: S("cube")},
		{Add: S("cube")},
		{Add: S("cube")},
		{Add: S("cube")},
		{Add: S("cube")},
		{Add: S("cube")},
		{Add: S("cube")},
		{Add: S("cube")},
		{Add: S("cube")},
	}
	for i, got := range y.Items {
		if !cmp.Equal(got, want[i]) {
			t.Errorf("item#%v =\n%v,\nwant\n%v", i+1, got, want[i])
		}
	}
}

func TestYAMLFile_Camera(t *testing.T) {
	r := bytes.NewBufferString(coverYAML)
	y, err := ParseYAML(r)
	if err != nil {
		t.Fatal(err)
	}

	c := y.Camera()
	want := &CameraT{
		HSize:       100,
		VSize:       99,
		FieldOfView: 0.785,
		Transform: M4{
			Tuple{-0.6987699101721224, -0.3144464595774551, 0.4061600102875461, 0},
			Tuple{-0.41036467732879783, 0.9119215051751063, 0, 0},
			Tuple{-0.5746957711326908, 0.2873478855663454, -0.7662610281769211, 0},
			Tuple{-6, 6, -10, 0},
		},
	}

	if got, want := c.HSize, want.HSize; got != want {
		t.Errorf("c.HSize = %v, want %v", got, want)
	}

	if got, want := c.VSize, want.VSize; got != want {
		t.Errorf("c.VSize = %v, want %v", got, want)
	}

	if got, want := c.FieldOfView, want.FieldOfView; got != want {
		t.Errorf("c.FieldOfView = %v, want %v", got, want)
	}

	if got, want := c.Transform[0], want.Transform[0]; !got.Equal(want) {
		t.Errorf("c.Transform[0] = %v, want %v", got, want)
	}
	if got, want := c.Transform[1], want.Transform[1]; !got.Equal(want) {
		t.Errorf("c.Transform[1] = %v, want %v", got, want)
	}
	if got, want := c.Transform[2], want.Transform[2]; !got.Equal(want) {
		t.Errorf("c.Transform[2] = %v, want %v", got, want)
	}
	if got, want := c.Transform[3], want.Transform[3]; !got.Equal(want) {
		t.Errorf("c.Transform[3] = %v, want %v", got, want)
	}
}

var coverYAML = `- add: camera
  field-of-view: 0.785
  from:
  - -6
  - 6
  - -10
  height: 99
  to:
  - 6
  - 0
  - 6
  up:
  - -0.45
  - 1
  - 0
  width: 100
- add: light
  at:
  - 50
  - 100
  - -50
  intensity:
  - 1
  - 1
  - 1
- add: light
  at:
  - -400
  - 50
  - -10
  intensity:
  - 0.2
  - 0.2
  - 0.2
- define: white-material
  value:
    ambient: 0.1
    color:
    - 1
    - 1
    - 1
    diffuse: 0.7
    reflective: 0.1
    specular: 0
- define: blue-material
  extend: white-material
  value:
    color:
    - 0.537
    - 0.831
    - 0.914
- define: red-material
  extend: white-material
  value:
    color:
    - 0.941
    - 0.322
    - 0.388
- define: purple-material
  extend: white-material
  value:
    color:
    - 0.373
    - 0.404
    - 0.55
- define: standard-transform
  value:
  - - translate
    - 1
    - -1
    - 1
  - - scale
    - 0.5
    - 0.5
    - 0.5
- define: large-object
  value:
  - standard-transform
  - - scale
    - 3.5
    - 3.5
    - 3.5
- define: medium-object
  value:
  - standard-transform
  - - scale
    - 3
    - 3
    - 3
- define: small-object
  value:
  - standard-transform
  - - scale
    - 2
    - 2
    - 2
- add: plane
  material:
    ambient: 1
    color:
    - 1
    - 1
    - 1
    diffuse: 0
    specular: 0
  transform:
  - - rotate-x
    - 1.5707963267948966
  - - translate
    - 0
    - 0
    - 500
- add: sphere
  material:
    ambient: 0
    color:
    - 0.373
    - 0.404
    - 0.55
    diffuse: 0.2
    reflective: 0.7
    refractive-index: 1.5
    shininess: 200
    specular: 1
    transparency: 0.7
  transform:
  - large-object
- add: cube
  material: white-material
  transform:
  - medium-object
  - - translate
    - 4
    - 0
    - 0
- add: cube
  material: blue-material
  transform:
  - large-object
  - - translate
    - 8.5
    - 1.5
    - -0.5
- add: cube
  material: red-material
  transform:
  - large-object
  - - translate
    - 0
    - 0
    - 4
- add: cube
  material: white-material
  transform:
  - small-object
  - - translate
    - 4
    - 0
    - 4
- add: cube
  material: purple-material
  transform:
  - medium-object
  - - translate
    - 7.5
    - 0.5
    - 4
- add: cube
  material: white-material
  transform:
  - medium-object
  - - translate
    - -0.25
    - 0.25
    - 8
- add: cube
  material: blue-material
  transform:
  - large-object
  - - translate
    - 4
    - 1
    - 7.5
- add: cube
  material: red-material
  transform:
  - medium-object
  - - translate
    - 10
    - 2
    - 7.5
- add: cube
  material: white-material
  transform:
  - small-object
  - - translate
    - 8
    - 2
    - 12
- add: cube
  material: white-material
  transform:
  - small-object
  - - translate
    - 20
    - 1
    - 9
- add: cube
  material: blue-material
  transform:
  - large-object
  - - translate
    - -0.5
    - -5
    - 0.25
- add: cube
  material: red-material
  transform:
  - large-object
  - - translate
    - 4
    - -4
    - 0
- add: cube
  material: white-material
  transform:
  - large-object
  - - translate
    - 8.5
    - -4
    - 0
- add: cube
  material: white-material
  transform:
  - large-object
  - - translate
    - 0
    - -4
    - 4
- add: cube
  material: purple-material
  transform:
  - large-object
  - - translate
    - -0.5
    - -4.5
    - 8
- add: cube
  material: white-material
  transform:
  - large-object
  - - translate
    - 0
    - -8
    - 4
- add: cube
  material: white-material
  transform:
  - large-object
  - - translate
    - -0.5
    - -8.5
    - 8
`
