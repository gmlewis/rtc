package rtc

import (
	"bytes"
	"strings"
	"testing"

	"sigs.k8s.io/yaml"
)

func TestParseYAML(t *testing.T) {
	r := bytes.NewBufferString(coverYAML)
	y, err := ParseYAML(r)
	if err != nil {
		t.Fatal(err)
	}

	buf, err := yaml.Marshal(y)
	if err != nil {
		t.Fatal(err)
	}

	gotLines := strings.Split(string(buf), "\n")
	wantLines := strings.Split(coverYAML, "\n")

	for i := 0; i < len(gotLines) && i < len(wantLines); i++ {
		if gotLines[i] == "Item:" {
			continue
		}

		if got, want := gotLines[i], wantLines[i]; got != want {
			t.Errorf("gotLines[%v] =\n%v\n, want\n%v", i+1, got, want)
		}
	}
}

var coverYAML = `
- add: camera
  field-of-view: 0.785
  from:
  - -6
  - 6
  - -10
  height: 100
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
