package filters

import (
	"image"
	"testing"

	"github.com/anthonynsimon/bild/util"
)

// PASSED
func TestCurves(t *testing.T) {

	// value struct contains input arguments of Curves()
	type value struct {
		img                          image.Image
		chans                        string
		second, third, fourth, fifth []float64
	}

	cases := []struct {
		*value
		expected *image.RGBA
	}{
		// first test
		{
			value: &value{
				&image.RGBA{},
				"b",
				[]float64{20, 0},
				[]float64{90, 120},
				[]float64{186, 144},
				[]float64{255, 230},
			},
			expected: &image.RGBA{
				Pix:    []uint8{},
				Stride: 0,
				Rect:   image.Rect(0, 0, 0, 0),
			},
		},
		// second test
		{
			value: &value{
				&image.RGBA{
					Pix:    []uint8{0x7f, 0x7f, 0x7f, 0x80, 0x7f, 0x7f, 0x7f, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x0, 0x0, 0x0, 0xFF},
					Stride: 8,
					Rect:   image.Rect(0, 0, 2, 2),
				},
				"r",
				[]float64{0, 0},
				[]float64{144, 90},
				[]float64{138, 120},
				[]float64{255, 255},
			},
			expected: &image.RGBA{
				Pix:    []uint8{99, 127, 127, 128, 99, 127, 127, 255, 255, 255, 255, 255, 0, 0, 0, 255},
				Stride: 8,
				Rect:   image.Rect(0, 0, 2, 2),
			},
		},
		// // third test
		{
			value: &value{
				&image.RGBA{
					Pix:    []uint8{0x7f, 0x7f, 0x7f, 0x80, 0x7f, 0x7f, 0x7f, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x0, 0x0, 0x0, 0xFF, 0x7f, 0x7f, 0x7f, 0x80, 0x7f, 0x7f, 0x7f, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x0, 0x0, 0x0, 0xFF},
					Stride: 8,
					Rect:   image.Rect(0, 0, 2, 4),
				},
				"g",
				[]float64{10, 0},
				[]float64{115, 105},
				[]float64{148, 100},
				[]float64{255, 248},
			},
			expected: &image.RGBA{
				Pix:    []uint8{127, 104, 127, 128, 127, 104, 127, 255, 255, 248, 255, 255, 0, 0, 0, 255, 127, 104, 127, 128, 127, 104, 127, 255, 255, 248, 255, 255, 0, 0, 0, 255},
				Stride: 8,
				Rect:   image.Rect(0, 0, 2, 4),
			},
		},
		// // fourth test
		{
			value: &value{
				&image.RGBA{
					Pix:    []uint8{0x7f, 0x7f, 0x7f, 0x80, 0x7f, 0x7f, 0x7f, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x0, 0x0, 0x0, 0xFF},
					Stride: 8,
					Rect:   image.Rect(0, 0, 2, 2),
				},
				"rgb",
				[]float64{0, 0},
				[]float64{120, 100},
				[]float64{128, 140},
				[]float64{255, 255},
			},
			expected: &image.RGBA{
				Pix:    []uint8{125, 125, 125, 128, 125, 125, 125, 255, 255, 255, 255, 255, 0, 0, 0, 255},
				Stride: 8,
				Rect:   image.Rect(0, 0, 2, 2),
			},
		},
	}

	for i, c := range cases {
		actual := Curves(c.img, c.chans, c.second, c.third, c.fourth, c.fifth)
		if !util.RGBAImageEqual(actual, c.expected) {
			t.Errorf("%s: case number %d\nexpected:\n%v\nactual:\n%v", "Curves", i+1, util.RGBAToString(c.expected), util.RGBAToString(actual))
		}
	}
}

func TestSaturation(t *testing.T) {
	cases := []struct {
		img    image.Image
		adjust float64

		expected *image.RGBA
	}{
		{
			&image.RGBA{},
			-34,
			&image.RGBA{
				Pix:    []uint8{},
				Stride: 0,
				Rect:   image.Rect(0, 0, 0, 0),
			},
		},
		{
			&image.RGBA{
				Pix:    []uint8{0x7f, 0x7f, 0x7f, 0x80, 0x7f, 0x7f, 0x7f, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x0, 0x0, 0x0, 0xFF},
				Stride: 8,
				Rect:   image.Rect(0, 0, 2, 2),
			},
			-23,
			&image.RGBA{
				Pix:    []uint8{127, 127, 127, 128, 127, 127, 127, 255, 255, 255, 255, 255, 0, 0, 0, 255},
				Stride: 8,
				Rect:   image.Rect(0, 0, 2, 2),
			},
		},
		{
			&image.RGBA{
				Pix:    []uint8{0x7f, 0x7f, 0x7f, 0x80, 0x7f, 0x7f, 0x7f, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x0, 0x0, 0x0, 0xFF, 0x7f, 0x7f, 0x7f, 0x80, 0x7f, 0x7f, 0x7f, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x0, 0x0, 0x0, 0xFF},
				Stride: 8,
				Rect:   image.Rect(0, 0, 2, 4),
			},
			-44,
			&image.RGBA{
				Pix:    []uint8{127, 127, 127, 128, 127, 127, 127, 255, 255, 255, 255, 255, 0, 0, 0, 255, 127, 127, 127, 128, 127, 127, 127, 255, 255, 255, 255, 255, 0, 0, 0, 255},
				Stride: 8,
				Rect:   image.Rect(0, 0, 2, 4),
			},
		},
		{
			&image.RGBA{
				Pix:    []uint8{0x7f, 0x7f, 0x7f, 0x80, 0x7f, 0x7f, 0x7f, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x0, 0x0, 0x0, 0xFF},
				Stride: 8,
				Rect:   image.Rect(0, 0, 2, 2),
			},
			-22,
			&image.RGBA{
				Pix:    []uint8{127, 127, 127, 128, 127, 127, 127, 255, 255, 255, 255, 255, 0, 0, 0, 255},
				Stride: 8,
				Rect:   image.Rect(0, 0, 2, 2),
			},
		},
	}

	for i, c := range cases {
		actual := Saturation(c.img, c.adjust)
		if !util.RGBAImageEqual(actual, c.expected) {
			t.Errorf("%s: case number %d\nexpected: \n\t%v\nactual: \n\t%v", "Saturation", i+1, c.expected, actual)
		}
	}
}
