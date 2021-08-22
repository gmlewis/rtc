package rtc

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

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
	y := &YAMLFile{DefinedItems: map[string]*Item{}}

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(b, &y.Items); err != nil {
		return nil, err
	}

	for i, item := range y.Items {
		if item.Define != nil {
			y.DefinedItems[*item.Define] = &y.Items[i]
		}
		if item.RawValue != nil {
			switch []byte(item.RawValue)[0] {
			case '[':
				log.Printf("item.Value is an array")
				y.Items[i].ValueArray, err = parseValueArray(item.RawValue)
				if err != nil {
					return nil, err
				}
				y.Items[i].RawValue = nil
			case '{':
				log.Printf("item.Value is an object")
				if err := json.Unmarshal(item.RawValue, &y.Items[i].ValueMaterial); err != nil {
					return nil, err
				}
				y.Items[i].RawValue = nil
			default:
				log.Fatalf("Unknown item.Value: %s", item.RawValue)
			}
		}
	}

	return y, nil
}

// ToYAML converts the scene back to a YAML file.
func (y *YAMLFile) ToYAML() ([]byte, error) {
	for i, item := range y.Items {
		if item.ValueMaterial != nil {
			buf, err := json.Marshal(item.ValueMaterial)
			if err != nil {
				return nil, err
			}
			y.Items[i].RawValue = buf
			y.Items[i].ValueMaterial = nil
			continue
		}

		if item.ValueArray != nil {
			var parts []string
			for _, v := range item.ValueArray {
				if v.NamedItem != nil {
					parts = append(parts, fmt.Sprintf("%q", *v.NamedItem))
					continue
				}
				if v.Type == nil || v.X == nil || v.Y == nil || v.Z == nil {
					return nil, fmt.Errorf("expected NamedItem or Type/X/Y/Z in ValueArray, got %#v", *v)
				}
				parts = append(parts, fmt.Sprintf("[%q,%v,%v,%v]", *v.Type, *v.X, *v.Y, *v.Z))
			}
			y.Items[i].RawValue = []byte(fmt.Sprintf("[%v]", strings.Join(parts, ",")))
			log.Printf("Created y.Items[%v].Value: %s", i, y.Items[i].RawValue)
			y.Items[i].ValueArray = nil
		}
	}

	return yaml.Marshal(y.Items)
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
	ValueMaterial *YAMLMaterial     `json:"-"`
	ValueArray    []*ValueArrayItem `json:"-"`
}

func parseValueArray(v json.RawMessage) ([]*ValueArrayItem, error) {
	var items []interface{}
	if err := json.Unmarshal(v, &items); err != nil {
		return nil, err
	}

	var result []*ValueArrayItem
	for _, item := range items {
		if s, ok := item.(string); ok {
			result = append(result, &ValueArrayItem{NamedItem: &s})
			continue
		}
		v, ok := item.([]interface{})
		if !ok {
			return nil, fmt.Errorf("expected string or transform array, but got %#v", item)
		}
		Type := v[0].(string)
		X := v[1].(float64)
		Y := v[2].(float64)
		Z := v[3].(float64)
		result = append(result, &ValueArrayItem{Type: &Type, X: &X, Y: &Y, Z: &Z})
	}
	return result, nil
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
	addValueMaterial := func(p []string, v *YAMLMaterial, n string) []string {
		if v == nil {
			return p
		}
		var p2 []string
		p2 = addString(p2, v.NamedItem, "NamedItem")
		p2 = addFloatArray(p2, v.Color, "Color")
		p2 = addFloat(p2, v.Diffuse, "Diffuse")
		p2 = addFloat(p2, v.Ambient, "Ambient")
		p2 = addFloat(p2, v.Specular, "Specular")
		p2 = addFloat(p2, v.Sininess, "Sininess")
		p2 = addFloat(p2, v.Reflective, "Reflective")
		p2 = addFloat(p2, v.Transparency, "Transparency")
		p2 = addFloat(p2, v.RefractiveIndex, "RefractiveIndex")
		p = append(p, fmt.Sprintf("%v:&YAMLMaterial{%v}", n, strings.Join(p2, ",")))
		return p
	}
	addValueArray := func(p []string, vs []*ValueArrayItem, n string) []string {
		if len(vs) == 0 {
			return p
		}
		var p2 []string
		for _, v := range vs {
			var items []string
			items = addString(items, v.NamedItem, "NamedItem")
			items = addString(items, v.Type, "Type")
			items = addFloat(items, v.X, "X")
			items = addFloat(items, v.Y, "Y")
			items = addFloat(items, v.Z, "Z")
			p2 = append(p2, strings.Join(items, ","))
		}
		p = append(p, fmt.Sprintf("%v:[]*ValueArrayItem{{%v}}", n, strings.Join(p2, "},{")))
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
	parts = addRaw(parts, i.RawValue, "Value")
	parts = addRaw(parts, i.RawMaterial, "Material")
	parts = addRaw(parts, i.RawTransform, "Transform")
	parts = addValueMaterial(parts, i.ValueMaterial, "ValueMaterial")
	parts = addValueArray(parts, i.ValueArray, "ValueArray")
	return fmt.Sprintf("{%v}", strings.Join(parts, ","))
}

// YAMLMaterial represents either a named DefinedItems value or a Material.
type YAMLMaterial struct {
	NamedItem *string `json:"-"`

	Color           []float64 `json:"color,omitempty"`
	Diffuse         *float64  `json:"diffuse,omitempty"`
	Ambient         *float64  `json:"ambient,omitempty"`
	Specular        *float64  `json:"specular,omitempty"`
	Sininess        *float64  `json:"shininess,omitempty"`
	Reflective      *float64  `json:"reflective,omitempty"`
	Transparency    *float64  `json:"transparency,omitempty"`
	RefractiveIndex *float64  `json:"refractive-index,omitempty"`
}

// ValueArrayItem is either a named DefinedItems value or a YAMLTransform.
type ValueArrayItem struct {
	NamedItem *string

	Type    *string
	X, Y, Z *float64
}
