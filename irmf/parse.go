package irmf

import (
	"io"
	"io/ioutil"

	"github.com/gmlewis/rtc/rtc"
)

// ParseString returns a new IRMF object from a string.
func ParseString(s string) *IRMFT {
	bounds := rtc.Bounds()

	return &IRMFT{
		Shape:  rtc.Shape{Transform: rtc.M4Identity(), Material: rtc.GetMaterial()},
		bounds: bounds,
	}
}

// Parse returns a new Parse object.
func Parse(r io.Reader) (*IRMFT, error) {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return ParseString(string(buf)), nil
}

// ParseFile returns a new IRMF object from a file.
func ParseFile(filename string) (*IRMFT, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return ParseString(string(buf)), nil
}
