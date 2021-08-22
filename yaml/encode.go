package yaml

import (
	"encoding/json"
	"fmt"
	"strings"

	"sigs.k8s.io/yaml"
)

// ToYAML converts the scene back to a YAML file.
func (y *YAMLFile) ToYAML() ([]byte, error) {
	for i, item := range y.Items {
		if v := item.Material; v != nil {
			var buf []byte
			if v.NamedItem != nil {
				buf = []byte(fmt.Sprintf("%q", *v.NamedItem))
			} else {
				var err error
				buf, err = json.Marshal(v)
				if err != nil {
					return nil, err
				}
			}
			if y.Items[i].Define != nil {
				y.Items[i].RawValue = buf
			} else {
				y.Items[i].RawMaterial = buf
			}
			y.Items[i].Material = nil
		}

		if item.Transform != nil {
			var parts []string

			for _, v := range item.Transform {
				if v.NamedItem != nil {
					parts = append(parts, fmt.Sprintf("%q", *v.NamedItem))
					continue
				}
				if v.Type == nil {
					return nil, fmt.Errorf("expected NamedItem or Type/Args in ValueArray, got %#v", *v)
				}
				var p2 []string
				for _, arg := range v.Args {
					p2 = append(p2, fmt.Sprintf("%v", arg))
				}
				parts = append(parts, fmt.Sprintf("[%q,%v]", *v.Type, strings.Join(p2, ",")))
			}

			if y.Items[i].Define != nil {
				y.Items[i].RawValue = []byte(fmt.Sprintf("[%v]", strings.Join(parts, ",")))
			} else {
				y.Items[i].RawTransform = []byte(fmt.Sprintf("[%v]", strings.Join(parts, ",")))
			}
			y.Items[i].Transform = nil
		}
	}

	return yaml.Marshal(y.Items)
}
