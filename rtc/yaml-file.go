package rtc

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/go-yaml/yaml"
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
	Add    *string `yaml:"add,omitempty"`
	Define *string `yaml:"define,omitempty"`

	// camera
	Width  *int      `yaml:"width,omitempty"`
	Height *int      `yaml:"height,omitempty"`
	FOV    *float64  `yaml:"field-of-view,omitempty"`
	From   []float64 `yaml:"from,omitempty"`
	To     []float64 `yaml:"to,omitempty"`
	Up     []float64 `yaml:"up,omitempty"`

	// light
	At        []float64 `yaml:"at,omitempty"`
	Intensity []float64 `yaml:"intensity,omitempty"`

	// define
	Extend *string `yaml:"extend,omitempty"`
	Value  *Value  `yaml:"value,omitempty"`
}

// Value represents a value within a "define" section.
type Value struct {
	Color      []float64 `yaml:"color,omitempty"`
	Diffuse    float64   `yaml:"diffuse,omitempty"`
	Ambient    float64   `yaml:"ambient,omitempty"`
	Specular   float64   `yaml:"specular,omitempty"`
	Reflective float64   `yaml:"reflective,omitempty"`
}
