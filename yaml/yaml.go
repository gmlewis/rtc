package yaml

import (
	"encoding/json"
	"fmt"
	"strings"
)

// YAMLFile represents a parsed yaml scene description file.
type YAMLFile struct {
	Items        []Item
	DefinedItems map[string]*Item
}

// Item represents anything that can be added to the scene.
type Item struct {
	Add    *string `json:"add,omitempty"`
	Define *string `json:"define,omitempty"`

	// camera
	Width  *int      `json:"width,omitempty"`
	Height *int      `json:"height,omitempty"`
	FOV    *float64  `json:"field-of-view,omitempty"`
	From   []float64 `json:"from,omitempty"`
	To     []float64 `json:"to,omitempty"`
	Up     []float64 `json:"up,omitempty"`

	// light
	At        []float64 `json:"at,omitempty"`
	Intensity []float64 `json:"intensity,omitempty"`

	// define
	Extend   *string         `json:"extend,omitempty"`
	RawValue json.RawMessage `json:"value,omitempty"`

	// object
	RawMaterial  json.RawMessage `json:"material,omitempty"`
	RawTransform json.RawMessage `json:"transform,omitempty"`

	// expanded raw messages:
	Material  *YAMLMaterial    `json:"-"`
	Transform []*YAMLTransform `json:"-"`
}

// YAMLMaterial represents either a named DefinedItems value or a Material.
type YAMLMaterial struct {
	NamedItem *string `json:"-"`

	Color           []float64 `json:"color,omitempty"`
	Diffuse         *float64  `json:"diffuse,omitempty"`
	Ambient         *float64  `json:"ambient,omitempty"`
	Specular        *float64  `json:"specular,omitempty"`
	Shininess       *float64  `json:"shininess,omitempty"`
	Reflective      *float64  `json:"reflective,omitempty"`
	Transparency    *float64  `json:"transparency,omitempty"`
	RefractiveIndex *float64  `json:"refractive-index,omitempty"`
}

// YAMLTransform is either a named DefinedItems value or a Transform.
type YAMLTransform struct {
	NamedItem *string

	Type *string
	Args []float64
}

func (i Item) String() string {
	var parts []string
	addInt := func(p []string, s *int, n string) []string {
		if s != nil {
			p = append(p, fmt.Sprintf("%v:I(%v)", n, *s))
		}
		return p
	}
	addFloat := func(p []string, s *float64, n string) []string {
		if s != nil {
			p = append(p, fmt.Sprintf("%v:F(%v)", n, *s))
		}
		return p
	}
	addFloatArray := func(p []string, s []float64, n string) []string {
		if len(s) > 0 {
			p = append(p, fmt.Sprintf("%v:%#v", n, s))
		}
		return p
	}
	addRaw := func(p []string, s json.RawMessage, n string) []string {
		if s != nil {
			p = append(p, fmt.Sprintf("%v:%s", n, s))
		}
		return p
	}
	addString := func(p []string, s *string, n string) []string {
		if s != nil {
			p = append(p, fmt.Sprintf("%v:S(%q)", n, *s))
		}
		return p
	}
	addYAMLMaterial := func(p []string, v *YAMLMaterial, n string) []string {
		if v == nil {
			return p
		}
		var p2 []string
		p2 = addString(p2, v.NamedItem, "NamedItem")
		p2 = addFloatArray(p2, v.Color, "Color")
		p2 = addFloat(p2, v.Diffuse, "Diffuse")
		p2 = addFloat(p2, v.Ambient, "Ambient")
		p2 = addFloat(p2, v.Specular, "Specular")
		p2 = addFloat(p2, v.Shininess, "Sininess")
		p2 = addFloat(p2, v.Reflective, "Reflective")
		p2 = addFloat(p2, v.Transparency, "Transparency")
		p2 = addFloat(p2, v.RefractiveIndex, "RefractiveIndex")
		p = append(p, fmt.Sprintf("%v:&YAMLMaterial{%v}", n, strings.Join(p2, ",")))
		return p
	}
	addYAMLTransforms := func(p []string, vs []*YAMLTransform, n string) []string {
		if len(vs) == 0 {
			return p
		}
		var p2 []string
		for _, v := range vs {
			var items []string
			items = addString(items, v.NamedItem, "NamedItem")
			items = addString(items, v.Type, "Type")
			items = addFloatArray(items, v.Args, "Args")
			p2 = append(p2, strings.Join(items, ","))
		}
		p = append(p, fmt.Sprintf("%v:[]*YAMLTransform{{%v}}", n, strings.Join(p2, "},{")))
		return p
	}

	parts = addString(parts, i.Add, "Add")
	parts = addString(parts, i.Define, "Define")
	parts = addInt(parts, i.Width, "Width")
	parts = addInt(parts, i.Height, "Height")
	parts = addFloat(parts, i.FOV, "FOV")
	parts = addFloatArray(parts, i.From, "From")
	parts = addFloatArray(parts, i.To, "To")
	parts = addFloatArray(parts, i.Up, "Up")
	parts = addFloatArray(parts, i.At, "At")
	parts = addFloatArray(parts, i.Intensity, "Intensity")
	parts = addString(parts, i.Extend, "Extend")
	parts = addRaw(parts, i.RawValue, "RawValue")
	parts = addRaw(parts, i.RawMaterial, "RawMaterial")
	parts = addRaw(parts, i.RawTransform, "RawTransform")
	parts = addYAMLMaterial(parts, i.Material, "Material")
	parts = addYAMLTransforms(parts, i.Transform, "Transform")
	return fmt.Sprintf("{%v}", strings.Join(parts, ","))
}
