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

	c := y.Camera(nil, nil, nil)
	want := &rtc.CameraT{
		HSize:       100,
		VSize:       99,
		FieldOfView: 0.785,
		Transform: rtc.M4{
			rtc.Tuple{-0.6987699101721224, -0.3144464595774551, 0.4061600102875461, 1.7556593993074578},
			rtc.Tuple{-0.12423884726468193, 0.7688585901430482, 0.38150110675215454, -1.5435735569248359},
			rtc.Tuple{-0.5746957711326908, 0.2873478855663454, -0.7662610281769211, -12.834872221963428},
			rtc.Tuple{0, 0, 0, 1},
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
