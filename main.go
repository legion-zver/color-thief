package color_thief

import (
	"errors"
	"github.com/legion-zver/color-thief/helper"
	"github.com/legion-zver/color-thief/wsm"
	"github.com/legion-zver/color-thief/wu"
	"image"
	"image/color"
)

// GetColorFromUri return the base color from the image file or uri
func GetColorFromUri(uri string) (color.Color, error) {
	colors, err := GetPaletteFromUri(uri, 10, 0)
	if err != nil {
		return color.RGBA{}, err
	}
	return colors[0], nil
}

// GetColor return the base color from the image
func GetColor(img image.Image, numColors, functionType int) (color.Color, error) {
	colors, err := GetPalette(img, numColors, functionType)
	if err != nil {
		return color.RGBA{}, err
	}
	return colors[0], nil
}

// GetPaletteFromUri return cluster similar colors from the image file or uri
func GetPaletteFromUri(uri string, numColors, functionType int) ([]color.Color, error) {
	img, err := helper.ReadImage(uri)
	if err != nil {
		return nil, err
	}
	return GetPalette(img, numColors, functionType)
}

// GetPalette return cluster similar colors by the median cut algorithm
func GetPalette(img image.Image, numColors, functionType int) ([]color.Color, error) {
	if numColors < 1 {
		return nil, errors.New("number of colors should be greater than 0")
	}
	var (
		pixels  = helper.SubsamplingPixelsFromImage(img)
		palette [][3]int
		colors  []color.Color
	)
	switch functionType {
	case 0:
		palette = wu.QuantWu(pixels, numColors)
		break
	case 1:
		palette = wsm.WSM(pixels, numColors)
		break
	default:
		return nil, errors.New("function type should be either 0 or 1")
	}
	colors = make([]color.Color, len(palette), len(palette))
	for i, v := range palette {
		colors[i] = helper.Color(v)
	}
	return colors, nil
}
