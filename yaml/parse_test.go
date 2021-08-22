package yaml

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParse(t *testing.T) {
	r := bytes.NewBufferString(coverYAML)
	y, err := Parse(r)
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
		{Define: S("white-material"), Material: &YAMLMaterial{Color: []float64{1, 1, 1}, Diffuse: F(0.7), Ambient: F(0.1), Specular: F(0), Reflective: F(0.1)}},
		{Define: S("blue-material"), Extend: S("white-material"), Material: &YAMLMaterial{Color: []float64{0.537, 0.831, 0.914}}},
		{Define: S("red-material"), Extend: S("white-material"), Material: &YAMLMaterial{Color: []float64{0.941, 0.322, 0.388}}},
		{Define: S("purple-material"), Extend: S("white-material"), Material: &YAMLMaterial{Color: []float64{0.373, 0.404, 0.55}}},
		{Define: S("standard-transform"), Transform: []*YAMLTransform{{Type: S("translate"), Args: []float64{1, -1, 1}}, {Type: S("scale"), Args: []float64{0.5, 0.5, 0.5}}}},
		{Define: S("large-object"), Transform: []*YAMLTransform{{NamedItem: S("standard-transform")}, {Type: S("scale"), Args: []float64{3.5, 3.5, 3.5}}}},
		{Define: S("medium-object"), Transform: []*YAMLTransform{{NamedItem: S("standard-transform")}, {Type: S("scale"), Args: []float64{3, 3, 3}}}},
		{Define: S("small-object"), Transform: []*YAMLTransform{{NamedItem: S("standard-transform")}, {Type: S("scale"), Args: []float64{2, 2, 2}}}},
		{Add: S("plane"), Material: &YAMLMaterial{Color: []float64{1, 1, 1}, Diffuse: F(0), Ambient: F(1), Specular: F(0)}, Transform: []*YAMLTransform{{Type: S("rotate-x"), Args: []float64{1.5707963267948966}}, {Type: S("translate"), Args: []float64{0, 0, 500}}}},
		{Add: S("sphere"), Material: &YAMLMaterial{Color: []float64{0.373, 0.404, 0.55}, Diffuse: F(0.2), Ambient: F(0), Specular: F(1), Shininess: F(200), Reflective: F(0.7), Transparency: F(0.7), RefractiveIndex: F(1.5)}, Transform: []*YAMLTransform{{NamedItem: S("large-object")}}},
		{Add: S("cube"), Material: &YAMLMaterial{NamedItem: S("white-material")}, Transform: []*YAMLTransform{{NamedItem: S("medium-object")}, {Type: S("translate"), Args: []float64{4, 0, 0}}}},
		{Add: S("cube"), Material: &YAMLMaterial{NamedItem: S("blue-material")}, Transform: []*YAMLTransform{{NamedItem: S("large-object")}, {Type: S("translate"), Args: []float64{8.5, 1.5, -0.5}}}},
		{Add: S("cube"), Material: &YAMLMaterial{NamedItem: S("red-material")}, Transform: []*YAMLTransform{{NamedItem: S("large-object")}, {Type: S("translate"), Args: []float64{0, 0, 4}}}},
		{Add: S("cube"), Material: &YAMLMaterial{NamedItem: S("white-material")}, Transform: []*YAMLTransform{{NamedItem: S("small-object")}, {Type: S("translate"), Args: []float64{4, 0, 4}}}},
		{Add: S("cube"), Material: &YAMLMaterial{NamedItem: S("purple-material")}, Transform: []*YAMLTransform{{NamedItem: S("medium-object")}, {Type: S("translate"), Args: []float64{7.5, 0.5, 4}}}},
		{Add: S("cube"), Material: &YAMLMaterial{NamedItem: S("white-material")}, Transform: []*YAMLTransform{{NamedItem: S("medium-object")}, {Type: S("translate"), Args: []float64{-0.25, 0.25, 8}}}},
		{Add: S("cube"), Material: &YAMLMaterial{NamedItem: S("blue-material")}, Transform: []*YAMLTransform{{NamedItem: S("large-object")}, {Type: S("translate"), Args: []float64{4, 1, 7.5}}}},
		{Add: S("cube"), Material: &YAMLMaterial{NamedItem: S("red-material")}, Transform: []*YAMLTransform{{NamedItem: S("medium-object")}, {Type: S("translate"), Args: []float64{10, 2, 7.5}}}},
		{Add: S("cube"), Material: &YAMLMaterial{NamedItem: S("white-material")}, Transform: []*YAMLTransform{{NamedItem: S("small-object")}, {Type: S("translate"), Args: []float64{8, 2, 12}}}},
		{Add: S("cube"), Material: &YAMLMaterial{NamedItem: S("white-material")}, Transform: []*YAMLTransform{{NamedItem: S("small-object")}, {Type: S("translate"), Args: []float64{20, 1, 9}}}},
		{Add: S("cube"), Material: &YAMLMaterial{NamedItem: S("blue-material")}, Transform: []*YAMLTransform{{NamedItem: S("large-object")}, {Type: S("translate"), Args: []float64{-0.5, -5, 0.25}}}},
		{Add: S("cube"), Material: &YAMLMaterial{NamedItem: S("red-material")}, Transform: []*YAMLTransform{{NamedItem: S("large-object")}, {Type: S("translate"), Args: []float64{4, -4, 0}}}},
		{Add: S("cube"), Material: &YAMLMaterial{NamedItem: S("white-material")}, Transform: []*YAMLTransform{{NamedItem: S("large-object")}, {Type: S("translate"), Args: []float64{8.5, -4, 0}}}},
		{Add: S("cube"), Material: &YAMLMaterial{NamedItem: S("white-material")}, Transform: []*YAMLTransform{{NamedItem: S("large-object")}, {Type: S("translate"), Args: []float64{0, -4, 4}}}},
		{Add: S("cube"), Material: &YAMLMaterial{NamedItem: S("purple-material")}, Transform: []*YAMLTransform{{NamedItem: S("large-object")}, {Type: S("translate"), Args: []float64{-0.5, -4.5, 8}}}},
		{Add: S("cube"), Material: &YAMLMaterial{NamedItem: S("white-material")}, Transform: []*YAMLTransform{{NamedItem: S("large-object")}, {Type: S("translate"), Args: []float64{0, -8, 4}}}},
		{Add: S("cube"), Material: &YAMLMaterial{NamedItem: S("white-material")}, Transform: []*YAMLTransform{{NamedItem: S("large-object")}, {Type: S("translate"), Args: []float64{-0.5, -8.5, 8}}}},
	}
	for i, got := range y.Items {
		if !cmp.Equal(got, want[i]) {
			t.Errorf("item#%v =\n%v,\nwant\n%v", i+1, got, want[i])
		}
	}
}
