package main

import (
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"os"
	"path/filepath"
	"strings"

	"github.com/anthonynsimon/bild/adjust"
	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/transform"
	"github.com/leminhson2398/image-filter/filters"
)

type filterFunc func(img *image.RGBA) *image.RGBA

// ImagePayload represents image object to perform transformation
type ImagePayload struct {
	Name     string `json:"name,omitempty"`
	Mimetype string `json:"mimetype,omitempty"`
	Filter   string `json:"filter,omitempty"`
}

const (
	stdDimen int    = 600
	help     string = `-help for help message
-images '[{"name": "...", "filter": "...", "mimetype": "..."}, ...]'
	+) name: path to image
	+) filter: sepiana | dawn | javana | charm | original | vintage | bright | gameron
	+) mimetype: ^image/(png|jpg|jpeg)$
	-nodejs: use JSON.stringify()
	-python: use json.dumps()
		then put the result after "$ ./image-filter -images "
-watermark: place watermark.png alongside the program
`
)

var (
	stdErr    = os.Stderr
	filterMap = map[string]filterFunc{
		"sepiana": func(img *image.RGBA) *image.RGBA {
			im := filters.SepiaCaman(img, 50)
			return im
		},
		"dawn": func(img *image.RGBA) *image.RGBA {
			colorize := filters.Colorize(img, &filters.RGB{R: 255, G: 105, B: 59}, 10)
			im := filters.CamanGamma(colorize, 1.2)
			return im
		},
		"javana": func(img *image.RGBA) *image.RGBA {
			saturation := filters.Saturation(img, -45)
			cur1 := filters.Curves(saturation, "b", []float64{20, 0}, []float64{90, 120}, []float64{186, 144}, []float64{255, 230})
			cur2 := filters.Curves(cur1, "r", []float64{0, 0}, []float64{144, 90}, []float64{138, 120}, []float64{255, 255})
			cur3 := filters.Curves(cur2, "g", []float64{10, 0}, []float64{115, 105}, []float64{148, 100}, []float64{255, 248})
			im := filters.Curves(cur3, "rgb", []float64{0, 0}, []float64{120, 100}, []float64{128, 140}, []float64{255, 255})
			return im
		},
		"charm": func(img *image.RGBA) *image.RGBA {
			im := filters.Channels(img, map[string]float64{"red": 8, "blue": 8})
			return im
		},
		"vintage": func(img *image.RGBA) *image.RGBA {
			grayscale := filters.GreyScale(img)
			sepia := filters.SepiaCaman(grayscale, 40)
			channels := filters.Channels(sepia, map[string]float64{"red": 8, "blue": 2, "green": 4})
			im := filters.CamanGamma(channels, 0.87)
			return im
		},
		"bright": func(img *image.RGBA) *image.RGBA {
			im := adjust.Brightness(img, 0.25)
			return im
		},
		"gameron": func(img *image.RGBA) *image.RGBA {
			im := filters.CamanGamma(img, 1.6)
			return im
		},
	}
)

func stdResizeImage(img *image.Image) *image.RGBA {
	bound := (*img).Bounds()
	width, height := bound.Dx(), bound.Dy()
	// calculate new dimensions of image, always less than stdDimen = 600px
	if width > stdDimen || height > stdDimen {
		if width > height {
			height = stdDimen * height / width
			width = stdDimen
		} else {
			width = stdDimen * width / height
			height = stdDimen
		}
	}
	return transform.Resize(*img, width, height, transform.Linear)
}

func setWhiteBackgroundAndWaterMark(img *image.RGBA) *image.RGBA {
	bound := img.Bounds()
	width, height := bound.Dx(), bound.Dy()

	// create new 600x600 image
	rec := image.Rect(0, 0, stdDimen, stdDimen)
	squareImage := image.NewRGBA(rec)
	whiteColor := color.RGBA{255, 255, 255, 0}

	// write 600x600 image, fill with white color
	draw.Draw(squareImage, squareImage.Bounds(), &image.Uniform{whiteColor}, image.ZP, draw.Src)

	// calculate rectangle to draw resized image over square image
	startPosX, startPosY := (stdDimen-width)/2, (stdDimen-height)/2
	overlayRec := image.Rect(startPosX, startPosY, startPosX+width, startPosY+height)
	draw.Draw(squareImage, overlayRec, img, image.ZP, draw.Over)

	// add watermark to image
	waterMarkImg := openImage("./watermark.png")
	if waterMarkImg != nil {
		offset := image.Pt(stdDimen-130, stdDimen-50)
		draw.Draw(squareImage, (*waterMarkImg).Bounds().Add(offset), *waterMarkImg, image.ZP, draw.Over)
	}

	return squareImage
}

func saveImageWithNewName(name string, img *image.RGBA) bool {
	encoder := imgio.JPEGEncoder(80)
	if err := imgio.Save(name, img, encoder); err != nil {
		return false
	}
	return true
}

func openImage(path string) *image.Image {
	img, err := imgio.Open(path)
	if err != nil {
		return nil
	}
	return &img
}

func createJPGName(oldName string) string {
	extension := filepath.Ext(oldName)
	if strings.ToLower(extension) == ".png" {
		return fmt.Sprintf("%s.jpg", oldName[:strings.LastIndex(oldName, extension)])
	}
	return oldName
}

func handleInput() {
	if len(os.Args) == 1 {
		fmt.Fprint(stdErr, help)
		return
	}
	cmd := os.Args[1]
	if !strings.EqualFold(cmd, "-images") || strings.EqualFold(cmd, "-help") {
		fmt.Fprint(stdErr, help)
		return
	}
	imageByte := []byte(os.Args[2])

	imagesPayload := []ImagePayload{}
	if err := json.Unmarshal(imageByte, &imagesPayload); err != nil {
		fmt.Fprint(stdErr, help)
		return
	}

	performTransform(&imagesPayload)
}

func performTransform(imageList *[]ImagePayload) {
	outputResult := []string{}

	for _, image := range *imageList {
		filterName := strings.ToLower(image.Filter)
		if filterFunction, ok := filterMap[filterName]; ok {
			img := openImage(image.Name)
			if img != nil {
				resizedImg := stdResizeImage(img)
				filterResult := filterFunction(resizedImg)
				setBackgrResult := setWhiteBackgroundAndWaterMark(filterResult)
				jpgName := createJPGName(image.Name)
				saveSuccess := saveImageWithNewName(jpgName, setBackgrResult)
				if saveSuccess {
					// append new name path to result list
					outputResult = append(outputResult, jpgName)
					// if input image's content type is "image/png", remove the image
					if image.Mimetype == "image/png" {
						os.Remove(image.Name)
					}
				} else {
					// if save not success, append old name to output list
					outputResult = append(outputResult, image.Name)
					// remove save-failed image path if it does exist
					os.Remove(jpgName)
				}
			}
		}
	}
	fmt.Print(strings.Join(outputResult, " "))
}

func main() {
	handleInput()
}
