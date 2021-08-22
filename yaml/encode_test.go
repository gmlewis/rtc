package yaml

import (
	"bytes"
	"strings"
	"testing"
)

func TestToYAML(t *testing.T) {
	r := bytes.NewBufferString(coverYAML)
	y, err := Parse(r)
	if err != nil {
		t.Fatal(err)
	}

	buf, err := y.ToYAML()
	if err != nil {
		t.Fatal(err)
	}

	gotLines := strings.Split(string(buf), "\n")
	wantLines := strings.Split(coverYAML, "\n")

	if got, want := len(gotLines), len(wantLines); got != want {
		t.Errorf("len(gotLines) = %v, want %v", got, want)
	}

	for i := 0; i < len(gotLines) && i < len(wantLines); i++ {
		if got, want := gotLines[i], wantLines[i]; got != want {
			t.Errorf("gotLines[%v] =\n%v\n, want\n%v", i+1, got, want)
		}
	}
}
