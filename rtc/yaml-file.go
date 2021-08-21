package rtc

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"

	"sigs.k8s.io/yaml"
)

// ParseYAMLFile parses a yaml scene description file and returns a YAMLFile.
func ParseYAMLFile(filename string) (*YAMLFile, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	obj, err := ParseYAML(f)
	if err != nil {
		return nil, err
	}

	if err := f.Close(); err != nil {
		return nil, err
	}

	return obj, nil
}

// ParseYAML parse a yaml scene description and returns a YAMLFile.
func ParseYAML(r io.Reader) (*YAMLFile, error) {
	y := &YAMLFile{}

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(b, &y.Item); err != nil {
		return nil, err
	}

	return y, nil
}

// YAMLFile represents a parsed yaml scene description file.
type YAMLFile struct {
	Item []Item
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
	Extend *string         `json:"extend,omitempty"`
	Value  json.RawMessage `json:"value,omitempty"`

	// object
	Material json.RawMessage `json:"material,omitempty"`
}

// YAMLMaterial represents a Material.
type YAMLMaterial struct {
	Color           []float64 `json:"color,omitempty"`
	Diffuse         *float64  `json:"diffuse,omitempty"`
	Ambient         *float64  `json:"ambient,omitempty"`
	Specular        *float64  `json:"specular,omitempty"`
	Sininess        *float64  `json:"shininess,omitempty"`
	Reflective      *float64  `json:"reflective,omitempty"`
	Transparency    *float64  `json:"transparency,omitempty"`
	RefractiveIndex *float64  `json:"refractive-index,omitempty"`
}
