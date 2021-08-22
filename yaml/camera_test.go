package yaml

import (
	"bytes"
	"testing"

	"github.com/gmlewis/rtc/rtc"
)

func TestCamera(t *testing.T) {
	r := bytes.NewBufferString(coverYAML)
	y, err := Parse(r)
	if err != nil {
		t.Fatal(err)
	}

	c := y.Camera()
	want := &rtc.CameraT{
		HSize:       100,
		VSize:       99,
		FieldOfView: 0.785,
		Transform: rtc.M4{
			rtc.Tuple{-0.6987699101721224, -0.3144464595774551, 0.4061600102875461, 0},
			rtc.Tuple{-0.41036467732879783, 0.9119215051751063, 0, 0},
			rtc.Tuple{-0.5746957711326908, 0.2873478855663454, -0.7662610281769211, 0},
			rtc.Tuple{-6, 6, -10, 0},
		},
	}

	if got, want := c.HSize, want.HSize; got != want {
		t.Errorf("c.HSize = %v, want %v", got, want)
	}

	if got, want := c.VSize, want.VSize; got != want {
		t.Errorf("c.VSize = %v, want %v", got, want)
	}

	if got, want := c.FieldOfView, want.FieldOfView; got != want {
		t.Errorf("c.FieldOfView = %v, want %v", got, want)
	}

	if got, want := c.Transform[0], want.Transform[0]; !got.Equal(want) {
		t.Errorf("c.Transform[0] = %v, want %v", got, want)
	}
	if got, want := c.Transform[1], want.Transform[1]; !got.Equal(want) {
		t.Errorf("c.Transform[1] = %v, want %v", got, want)
	}
	if got, want := c.Transform[2], want.Transform[2]; !got.Equal(want) {
		t.Errorf("c.Transform[2] = %v, want %v", got, want)
	}
	if got, want := c.Transform[3], want.Transform[3]; !got.Equal(want) {
		t.Errorf("c.Transform[3] = %v, want %v", got, want)
	}
}
