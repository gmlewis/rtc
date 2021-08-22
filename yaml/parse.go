// Package yaml provides a RTC yaml file parser as described in the book.
package yaml

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"sigs.k8s.io/yaml"
)

// ParseFile parses a yaml scene description file and returns a YAMLFile.
func ParseFile(filename string) (*YAMLFile, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	obj, err := Parse(f)
	if err != nil {
		return nil, err
	}

	if err := f.Close(); err != nil {
		return nil, err
	}

	return obj, nil
}

// Parse parse a yaml scene description and returns a YAMLFile.
func Parse(r io.Reader) (*YAMLFile, error) {
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
				y.Items[i].Transform, err = parseTransform(item.RawValue)
				if err != nil {
					return nil, err
				}
				y.Items[i].RawValue = nil
			case '{':
				if err := json.Unmarshal(item.RawValue, &y.Items[i].Material); err != nil {
					return nil, err
				}
				y.Items[i].RawValue = nil
			default:
				return nil, fmt.Errorf("unknown item.Value: %s", item.RawValue)
			}
		}

		if item.RawMaterial != nil {
			material := &YAMLMaterial{}
			switch []byte(item.RawMaterial)[0] {
			case '"':
				namedItem := strings.Trim(string(item.RawMaterial), "\"")
				material.NamedItem = &namedItem
			case '{':
				if err := json.Unmarshal(item.RawMaterial, &material); err != nil {
					return nil, err
				}
			}
			y.Items[i].RawMaterial = nil
			y.Items[i].Material = material
		}

		if item.RawTransform != nil {
			if []byte(item.RawTransform)[0] != '[' {
				return nil, fmt.Errorf("expected RawTransform to start with '[', got %s", item.RawTransform)
			}
			y.Items[i].Transform, err = parseTransform(item.RawTransform)
			if err != nil {
				return nil, err
			}
			y.Items[i].RawTransform = nil
		}
	}

	return y, nil
}

func parseTransform(v json.RawMessage) ([]*YAMLTransform, error) {
	var items []interface{}
	if err := json.Unmarshal(v, &items); err != nil {
		return nil, err
	}

	var result []*YAMLTransform
	for _, item := range items {
		if s, ok := item.(string); ok {
			result = append(result, &YAMLTransform{NamedItem: &s})
			continue
		}
		v, ok := item.([]interface{})
		if !ok {
			return nil, fmt.Errorf("expected string or transform array, but got %#v", item)
		}
		Type := v[0].(string)
		var args []float64
		for i := 1; i < len(v); i++ {
			args = append(args, v[i].(float64))
		}
		result = append(result, &YAMLTransform{Type: &Type, Args: args})
	}
	return result, nil
}
