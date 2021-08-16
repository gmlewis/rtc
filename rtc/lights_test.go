package rtc

import (
	"reflect"
	"testing"
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
