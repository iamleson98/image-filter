package operation

import (
	"encoding/json"
	"image"
	"log"
	"os"
	"strings"

	"github.com/anthonynsimon/bild/adjust"
	"github.com/leminhson2398/image-filter/filters"
)

type filterFunc func(img *image.RGBA) *image.RGBA

type ImagePayload struct {
	Name   string `json:"name,omitempty"`
	Filter string `json:"filter,omitempty"`
}

var filterMap map[string]filterFunc

func init() {
	filterMap["sepiana"] = sepiana
}

func sepiana(img *image.RGBA) *image.RGBA {
	im := filters.SepiaCaman(img, 50)
	return im
}

func dawn(img *image.RGBA) *image.RGBA {
	colorize := filters.Colorize(img, &filters.RGB{R: 255, G: 105, B: 59}, 10)
	im := filters.CamanGamma(colorize, 1.2)
	return im
}

func javana(img *image.RGBA) *image.RGBA {
	saturation := filters.Saturation(img, -45)
	cur1 := filters.Curves(saturation, "b", []float64{20, 0}, []float64{90, 120}, []float64{186, 144}, []float64{255, 230})
	cur2 := filters.Curves(cur1, "r", []float64{0, 0}, []float64{144, 90}, []float64{138, 120}, []float64{255, 255})
	cur3 := filters.Curves(cur2, "g", []float64{10, 0}, []float64{115, 105}, []float64{148, 100}, []float64{255, 248})
	im := filters.Curves(cur3, "rgb", []float64{0, 0}, []float64{120, 100}, []float64{128, 140}, []float64{255, 255})
	return im
}

func charm(img *image.RGBA) *image.RGBA {
	im := filters.Channels(img, map[string]float64{"red": 8, "blue": 8})
	return im
}

func vintage(img *image.RGBA) *image.RGBA {
	grayscale := filters.GreyScale(img)
	sepia := filters.SepiaCaman(grayscale, 40)
	channels := filters.Channels(sepia, map[string]float64{"red": 8, "blue": 2, "green": 4})
	im := filters.CamanGamma(channels, 0.87)
	return im
}

func bright(img *image.RGBA) *image.RGBA {
	im := adjust.Brightness(img, 0.25)
	return im
}

func gameron(img *image.RGBA) *image.RGBA {
	im := filters.CamanGamma(img, 1.6)
	return im
}

func HandleInput() {
	if len(os.Args) == 1 {
		log.Println(help)
		return
	}
	cmd := os.Args[1]
	if !strings.EqualFold(cmd, "-images") || strings.EqualFold(cmd, "-help") {
		log.Println(help)
		return
	}
	imageByte := []byte(os.Args[2])

	imagesPayload := []ImagePayload{}
	if err := json.Unmarshal(imageByte, &imagesPayload); err != nil {
		log.Println(help)
		return
	}

	performTransform(&imagesPayload)
}

func performTransform(imageList *[]ImagePayload) {
	for _, image := range *imageList {
		switch strings.ToLower(image.Filter) {
		case "sepiana":
			break
		case "dawn":
			break
		case "javana":
			break
		case "charm":
			break
		case "vintage":
			break
		case "bright":
			break
		case "gameron":
			break
		default:
			break
		}

	}
}
