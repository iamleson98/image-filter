package main

import (
	"fmt"
	_ "image/jpeg"
	_ "image/png"
	"path/filepath"
	"strings"

	// "github.com/anthonynsimon/bild/adjust"
	// "github.com/anthonynsimon/bild/effect"
	"github.com/anthonynsimon/bild/imgio"
	"github.com/leminhson2398/image-filter/filters"
	// "path/filepath"
	// "strings"
)

func main() {

	images := []string{
		"./upload-933310244.jpg",
	}

	for _, filename := range images {

		img, err := imgio.Open(filename)
		if err != nil {
			fmt.Println(err)
		}

		// -------------------transform:
		// transform := transform.Resize(img, 500, 330, transform.Linear)
		// if err := imgio.Save("Resize.jpg", transform, imgio.JPEGEncoder(80)); err != nil {
		// 	fmt.Println(err)
		// 	return
		// }

		// 	// ------------------Brightness:
		// Brightness := adjust.Brightness(img, 0.25)
		// savename0 := fmt.Sprintf("Brightness__%s.jpg", strings.Split(filepath.Base(filename), ".")[0])

		// if err := imgio.Save(savename0, Brightness, imgio.JPEGEncoder(80)); err != nil {
		// 	fmt.Println(err)
		// 	return
		// }

		// // 	// ------------------gamma
		// Gamma := filters.CamanGamma(img, 1.6)
		// savename6 := fmt.Sprintf("Gamma__%s.jpg", strings.Split(filepath.Base(filename), ".")[0])

		// if err := imgio.Save(savename6, Gamma, imgio.JPEGEncoder(80)); err != nil {
		// 	fmt.Println(err)
		// 	return
		// }

		// // 	// ---------------------jarques---------------------
		// Saturation := filters.Saturation(img, -45)

		// cur1 := filters.Curves(Saturation, "b", []float64{20, 0}, []float64{90, 120}, []float64{186, 144}, []float64{255, 230})
		// cur2 := filters.Curves(cur1, "r", []float64{0, 0}, []float64{144, 90}, []float64{138, 120}, []float64{255, 255})
		// cur3 := filters.Curves(cur2, "g", []float64{10, 0}, []float64{115, 105}, []float64{148, 100}, []float64{255, 248})
		// cur4 := filters.Curves(cur3, "rgb", []float64{0, 0}, []float64{120, 100}, []float64{128, 140}, []float64{255, 255})

		// savename2 := fmt.Sprintf("Jarques__%s.jpg", strings.Split(filepath.Base(filename), ".")[0])
		// if err := imgio.Save(savename2, cur4, imgio.JPEGEncoder(80)); err != nil {
		// 	fmt.Println(err)
		// 	return
		// }

		// // 	// ----------------------------sepiana----------------        (0-100)
		// SepiaCaman := filters.SepiaCaman(img, 50)
		// savename2 = fmt.Sprintf("Sepia__%s.jpg", strings.Split(filepath.Base(filename), ".")[0])

		// if err := imgio.Save(savename2, SepiaCaman, imgio.JPEGEncoder(80)); err != nil {
		// 	fmt.Println(err)
		// 	return
		// }

		// // 	// ---------------Grayscale-------------------
		// Grayscale := effect.Grayscale(img)
		// savename3 := fmt.Sprintf("Grayscale__%s.jpg", strings.Split(filepath.Base(filename), ".")[0])

		// if err := imgio.Save(savename3, Grayscale, imgio.JPEGEncoder(80)); err != nil {
		// 	fmt.Println(err)
		// 	return
		// }

		// // 	// ------------------Channels:
		Channels := filters.Channels(img, map[string]float64{"red": 8, "blue": 5})
		savename4 := fmt.Sprintf("Channels__%s.jpg", strings.Split(filepath.Base(filename), ".")[0])

		if err := imgio.Save(savename4, Channels, imgio.JPEGEncoder(80)); err != nil {
			fmt.Println(err)
			return
		}

		// // dawn:
		colorize := filters.Colorize(img, &filters.RGB{R: 255, G: 105, B: 59}, 10)
		gamm := filters.CamanGamma(colorize, 1.2)

		savename7 := fmt.Sprintf("Dawn_%s.jpg", strings.Split(filepath.Base(filename), ".")[0])
		if err := imgio.Save(savename7, gamm, imgio.JPEGEncoder(80)); err != nil {
			fmt.Println(err)
			return
		}
	}
}
