package rtc

import (
	"math"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPointLight(t *testing.T) {
	tests := []struct {
		name      string
		position  Tuple
		intensity Tuple
		want      *PointLightT
	}{
		{
			name:      "A point light has a position and intensity",
			position:  Point(0, 0, 0),
			intensity: Color(1, 1, 1),
			want:      &PointLightT{position: Point(0, 0, 0), intensity: Color(1, 1, 1)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PointLight(tt.position, tt.intensity); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PointLight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLighting(t *testing.T) {
	sq2 := math.Sqrt(2) / 2

	tests := []struct {
		name         string
		light        *PointLightT
		eyeVector    Tuple
		normalVector Tuple
		want         Tuple
	}{
		{
			name:         "Lighting with the eye between the light and the surface",
			eyeVector:    Vector(0, 0, -1),
			normalVector: Vector(0, 0, -1),
			light:        PointLight(Point(0, 0, -10), Color(1, 1, 1)),
			want:         Color(1.9, 1.9, 1.9),
		},
		{
			name:         "Lighting with the eye between light and surface, eye offset 45°",
			eyeVector:    Vector(0, sq2, -sq2),
			normalVector: Vector(0, 0, -1),
			light:        PointLight(Point(0, 0, -10), Color(1, 1, 1)),
			want:         Color(1, 1, 1),
		},
		{
			name:         "Lighting with eye opposite surface, light offset 45°",
			eyeVector:    Vector(0, 0, -1),
			normalVector: Vector(0, 0, -1),
			light:        PointLight(Point(0, 10, -10), Color(1, 1, 1)),
			want:         Color(0.7364, 0.7364, 0.7364),
		},
		{
			name:         "Lighting with eye in the path of the reflection vector",
			eyeVector:    Vector(0, -sq2, -sq2),
			normalVector: Vector(0, 0, -1),
			light:        PointLight(Point(0, 10, -10), Color(1, 1, 1)),
			want:         Color(1.6364, 1.6364, 1.6364),
		},
		{
			name:         "Lighting with the light behind the surface",
			eyeVector:    Vector(0, 0, -1),
			normalVector: Vector(0, 0, -1),
			light:        PointLight(Point(0, 0, 10), Color(1, 1, 1)),
			want:         Color(0.1, 0.1, 0.1),
		},
	}

	m := Material()
	position := Point(0, 0, 0)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Lighting(&m, tt.light, position, tt.eyeVector, tt.normalVector); !cmp.Equal(got, tt.want) {
				t.Errorf("Lighting() = %v, want %v", got, tt.want)
			}
		})
	}
}
