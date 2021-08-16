package rtc

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMaterial(t *testing.T) {
	tests := []struct {
		name string
		want *MaterialT
	}{
		{
			name: "The default material",
			want: &MaterialT{
				color:     Color(1, 1, 1),
				ambient:   0.1,
				diffuse:   0.9,
				specular:  0.9,
				shininess: 200.0,
			},
		},
	}

	opt := cmp.Comparer(func(a, b MaterialT) bool {
		return a.color == b.color &&
			a.ambient == b.ambient &&
			a.diffuse == b.diffuse &&
			a.specular == b.specular &&
			a.shininess == b.shininess
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Material()

			if !cmp.Equal(got, tt.want, opt) {
				t.Errorf("Material() = %v, want %v", got, tt.want)
			}
		})
	}
}
