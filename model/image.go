package model

import (
	"image"
	"image/color"
)

type Image struct {
	Width, Height int
	ColourSeed    uint8
}

// ColorModel returns the Image's color model.
func (img *Image) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds returns the domain for which At can return non-zero color.
// The bounds do not necessarily contain the point (0, 0).
func (img *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.Width, img.Height)
}

// At returns the color of the pixel at (x, y).
// At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
// At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
func (img *Image) At(x, y int) color.Color {
	return color.RGBA{
		R: img.ColourSeed + uint8(x*y),
		G: img.ColourSeed + uint8(x+y),
		B: 255,
		A: 255,
	}
}
