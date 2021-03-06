package rtc

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMaterial(t *testing.T) {
	tests := []struct {
		name string
		want MaterialT
	}{
		{
			name: "The default material",
			want: MaterialT{
				Color:           Color(1, 1, 1),
				Ambient:         0.1,
				Diffuse:         0.9,
				Specular:        0.9,
				Shininess:       200.0,
				Reflective:      0,
				Transparency:    0,
				RefractiveIndex: 1,
			},
		},
	}

	opt := cmp.Comparer(func(a, b MaterialT) bool {
		return a.Color == b.Color &&
			a.Ambient == b.Ambient &&
			a.Diffuse == b.Diffuse &&
			a.Specular == b.Specular &&
			a.Shininess == b.Shininess &&
			a.Reflective == b.Reflective &&
			a.Transparency == b.Transparency &&
			a.RefractiveIndex == b.RefractiveIndex
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetMaterial()

			if !cmp.Equal(got, tt.want, opt) {
				t.Errorf("GetMaterial() = %v, want %v", got, tt.want)
			}
		})
	}
}
