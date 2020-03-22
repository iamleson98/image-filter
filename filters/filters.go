package filters

import (
	"image"
	"image/color"
	"math"
	"strings"

	"github.com/anthonynsimon/bild/adjust"
	"github.com/leminhson2398/image-filter/helpers"
	"github.com/leminhson2398/image-filter/utils"
)

// RGB represent RGB color object
type RGB struct {
	R float64
	G float64
	B float64
}

// Contrast applies contrast effect to the image.
// NOTE: NOT WORKED YET => NOT TESTED
func Contrast(img image.Image, adj float64) *image.RGBA {
	adj = math.Pow((adj+100)/100, 2)
	fn := func(c color.RGBA) color.RGBA {
		floatR := float64(c.R)
		floatG := float64(c.G)
		floatB := float64(c.B)

		c.R = uint8(math.Round(((floatR/255-0.5)*adj + 0.5) * 255))
		c.G = uint8(math.Round(((floatG/255-0.5)*adj + 0.5) * 255))
		c.B = uint8(math.Round(((floatB/255-0.5)*adj + 0.5) * 255))

		return c
	}

	return adjust.Apply(img, fn)
}

// Curves applies curves effect to the image
// TESTED PASS
func Curves(img image.Image, chans string, second, third, fourth, fifth []float64) *image.RGBA {
	bezier := helpers.Bezier(second, third, fourth, fifth, 0, 255)

	start := second
	for i := 0; i < int(start[0]); i++ {
		if start[0] > 0 {
			bezier[i] = start[1]
		}
	}

	end := fifth
	for i := int(end[0]); i <= 255; i++ {
		if end[0] < 255 {
			bezier[i] = end[1]
		}
	}

	fn := func(c color.RGBA) color.RGBA {
		if strings.Contains(chans, "r") {
			c.R = uint8(math.Round(bezier[int(c.R)]))
		}
		if strings.Contains(chans, "g") {
			c.G = uint8(math.Round(bezier[int(c.G)]))
		}
		if strings.Contains(chans, "b") {
			c.B = uint8(math.Round(bezier[int(c.B)]))
		}

		return c
	}

	return adjust.Apply(img, fn)
}

// Saturation applies Saturation to image
// adj must be negative
// PASSED
func Saturation(img image.Image, adj float64) *image.RGBA {
	adj *= -0.01

	fn := func(c color.RGBA) color.RGBA {
		max := utils.MaxUint8(c.R, utils.MaxUint8(c.G, c.B))
		if c.R != max {
			c.R += uint8(math.Round(float64(max-c.R) * adj))
		}
		if c.G != max {
			c.G += uint8(math.Round(float64(max-c.G) * adj))
		}
		if c.B != max {
			c.B += uint8(math.Round(float64(max-c.B) * adj))
		}

		return c
	}

	return adjust.Apply(img, fn)
}

// Vibrance apply vibrance effect to the image
// NOTE: NOT WORKED YET
func Vibrance(img image.Image, adj float64) *image.RGBA {
	adj *= -1

	fn := func(c color.RGBA) color.RGBA {

		floatR := float64(c.R)
		floatG := float64(c.G)
		floatB := float64(c.B)

		max := math.Max(floatR, math.Max(floatG, floatB))
		avg := (floatR + floatG + floatB) / 3
		amt := math.Abs(max-avg) * 2 / 255 * adj / 100

		if floatR != max {
			c.R += uint8(math.Round((max - floatR) * amt))
		}
		if floatG != max {
			c.G += uint8(math.Round((max - floatG) * amt))
		}
		if floatR != max {
			c.B += uint8(math.Round((max - floatB) * amt))
		}

		return c
	}
	return adjust.Apply(img, fn)
}

// Exposure applies exposure effect to image
func Exposure(img image.Image, adj float64) *image.RGBA {

	p := math.Abs(adj) / 100
	ctrl1 := []float64{0, 255 * p}
	ctrl2 := []float64{255 - (255 * p), 255}

	if adj < 0 {
		// reverse arrays
		ctrl1 = []float64{255 * p, 0}
		ctrl2 = []float64{255, 255 - (255 * p)}
	}
	return Curves(img, "rgb", []float64{0, 0}, ctrl1, ctrl2, []float64{255, 255})
}

// CamanGamma applies gamma effect tom image
func CamanGamma(img image.Image, adj float64) *image.RGBA {
	fn := func(c color.RGBA) color.RGBA {
		c.R = uint8(math.Pow(float64(c.R)/255, adj) * 255)
		c.G = uint8(math.Pow(float64(c.G)/255, adj) * 255)
		c.B = uint8(math.Pow(float64(c.B)/255, adj) * 255)

		return c
	}

	return adjust.Apply(img, fn)
}

// Channels aplies channels effect to the image
func Channels(img image.Image, options map[string]float64) *image.RGBA {
	for k := range options {
		options[k] /= 100
	}
	var maxColor float64 = 255

	fn := func(c color.RGBA) color.RGBA {
		if red, ok := options["red"]; ok {
			if red > 0 {
				c.R += uint8(math.Round((maxColor - float64(c.R)) * red))
			} else {
				c.R -= uint8(math.Round(float64(c.R) * -red))
			}
		}
		if green, ok := options["green"]; ok {
			if green > 0 {
				c.G += uint8(math.Round((maxColor - float64(c.G)) * green))
			} else {
				c.G -= uint8(math.Round(float64(c.G) * -green))
			}
		}
		if blue, ok := options["blue"]; ok {
			if blue > 0 {
				c.B += uint8(math.Round((maxColor - float64(c.B)) * blue))
			} else {
				c.B -= uint8(math.Round(float64(c.B) * -blue))
			}
		}
		return c
	}

	return adjust.Apply(img, fn)
}

// SepiaCaman applies sepia effect to the image
func SepiaCaman(img image.Image, adj float64) *image.RGBA {
	adj /= 100
	var maxColor float64 = 255

	fn := func(rgba color.RGBA) color.RGBA {
		floatR := float64(rgba.R)
		floatG := float64(rgba.G)
		floatB := float64(rgba.B)

		rgba.R = uint8(math.Min(maxColor, (floatR*(1-(0.607*adj)))+(floatG*(0.769*adj))+(floatB*(0.189*adj))))
		rgba.G = uint8(math.Min(maxColor, (floatR*(0.349*adj))+(floatG*(1-(0.314*adj)))+(floatB*(0.168*adj))))
		rgba.B = uint8(math.Min(maxColor, (floatR*(0.272*adj))+(floatG*(0.534*adj))+(floatB*(1-(0.869*adj)))))

		return rgba
	}

	return adjust.Apply(img, fn)
}

// Colorize applies colorize effect to image
func Colorize(img image.Image, rgb *RGB, level float64) *image.RGBA {
	fn := func(c color.RGBA) color.RGBA {
		floatR := float64(c.R)
		floatG := float64(c.G)
		floatB := float64(c.B)

		c.R -= uint8(math.Round((floatR - rgb.R) * level / 100))
		c.G -= uint8(math.Round((floatG - rgb.G) * level / 100))
		c.B -= uint8(math.Round((floatB - rgb.B) * level / 100))
		return c
	}

	return adjust.Apply(img, fn)
}

// Posterize applies Posterize to image
func Posterize(img image.Image, adj float64) *image.RGBA {
	numOfAreas := 256 / adj
	numOfValues := 255 / (adj - 1)

	fn := func(c color.RGBA) color.RGBA {
		c.R = uint8(math.Floor(math.Floor(float64(c.R)/numOfAreas) * numOfValues))
		c.G = uint8(math.Floor(math.Floor(float64(c.G)/numOfAreas) * numOfValues))
		c.B = uint8(math.Floor(math.Floor(float64(c.B)/numOfAreas) * numOfValues))

		return c
	}

	return adjust.Apply(img, fn)
}

// Vignette applies vignette effect to the image, size must not exceed the range (0, 100) inclusively
// func Vignette(img image.Image, sizePercent, strength float64) *image.RGBA {
// 	bounds := img.Bounds()
// 	height, width := float64(bounds.Dy()), float64(bounds.Dx())

// 	if height > width {
// 		sizePercent = width * (sizePercent / 100)
// 	} else {
// 		sizePercent = height * (sizePercent / 100)
// 	}

// 	strength /= 100
// 	center := []float64{width / 2, height / 2}
// 	start := math.Sqrt(math.Pow(center[0], 2) + math.Pow(center[1], 2))
// 	end := start - sizePercent
// 	bezier := helpers.Bezier([]float64{0, 1}, []float64{30, 30}, []float64{70, 60}, []float64{100, 80}, 0, 255)

// 	fn := func(c color.RGBA, coord *adjust.Coord) color.RGBA {
// 		dist := helpers.Distance(coord.X, coord.Y, center[0], center[1])
// 		if dist > end {
// 			div := math.Max(1, bezier[int(math.Round((dist-end)/sizePercent*100))]/10)
// 			c.R = uint8(math.Round(math.Pow(float64(c.R/255), div) * 255))
// 			c.G = uint8(math.Round(math.Pow(float64(c.G/255), div) * 255))
// 			c.B = uint8(math.Round(math.Pow(float64(c.B/255), div) * 255))
// 		}
// 		return c
// 	}

// 	return adjust.CustomApply(img, fn)
// }
