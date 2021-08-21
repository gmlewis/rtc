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

	if err := yaml.Unmarshal(b, &y.Items); err != nil {
		return nil, err
	}

	return y, nil
}

// Camera returns the YAML camera definition.
func (y *YAMLFile) Camera() *CameraT {
	for _, item := range y.Items {
		if item.Add == nil || item.Width == nil || item.Height == nil || item.FOV == nil {
			continue
		}
		if *item.Add == "camera" {
			camera := Camera(*item.Width, *item.Height, *item.FOV)
			from := Vector(item.From[0], item.From[1], item.From[2])
			to := Vector(item.To[0], item.To[1], item.To[2])
			up := Vector(item.Up[0], item.Up[1], item.Up[2]).Normalize()
			forward := from.Sub(to).Normalize()
			right := up.Cross(forward)
			camera.Transform = M4{
				Tuple{right.X(), right.Y(), right.Z(), 0},
				Tuple{up.X(), up.Y(), up.Z(), 0},
				Tuple{forward.X(), forward.Y(), forward.Z(), 0},
				Tuple{from.X(), from.Y(), from.Z(), 0},
			}
			return camera
		}
	}
	return nil
}

// YAMLFile represents a parsed yaml scene description file.
type YAMLFile struct {
	Items []Item
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
	Material  json.RawMessage `json:"material,omitempty"`
	Transform json.RawMessage `json:"transform,omitempty"`
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
