package image

import (
	"image/gif"
	"math"
	"image/color"
	"image"
	"image/draw"
	"os"

	"github.com/nfnt/resize"
)

type GifService struct {
	image *gif.GIF
}

func (g *GifService) Resize(width, height uint) error {
	ratio := getRatio()

	for index, frame := range g.image.Image {
		rect := frame.Bounds()
		tmpImage := frame.SubImage(rect)
		resizedImage := resize.Resize(uint(math.Floor(float64(rect.Dx())*ratio)),
			uint(math.Floor(float64(rect.Dy())*ratio)),
			tmpImage, resize.Lanczos3)
		// Add colors from original gif image
		var tmpPalette color.Palette
		for x := 1; x <= rect.Dx(); x++ {
			for y := 1; y <= rect.Dy(); y++ {
				if !contains(tmpPalette, g.image.Image[index].At(x, y)) {
					tmpPalette = append(tmpPalette, g.image.Image[index].At(x, y))
				}
			}
		}
		// After first image, image may contains only difference
		// bounds may not start from at (0,0)
		resizedBounds := resizedImage.Bounds()
		if index >= 1 {
			marginX := int(math.Floor(float64(rect.Min.X) * ratio))
			marginY := int(math.Floor(float64(rect.Min.Y) * ratio))
			resizedBounds = image.Rect(marginX, marginY, resizedBounds.Dx()+marginX,
				resizedBounds.Dy()+marginY)
		}
		resizedPalette := image.NewPaletted(resizedBounds, tmpPalette)
		draw.Draw(resizedPalette, resizedBounds, resizedImage, image.ZP, draw.Src)
		g.image.Image[index] = resizedPalette
	}

	// Set size to resized size
	g.image.Config.Width = int(math.Floor(float64(width) * ratio))
	g.image.Config.Height = int(math.Floor(float64(height) * ratio))

	out, err := os.Create("resized." + Gif.String())
	if err != nil {
		return err
	}
	defer out.Close()

	var gifImage gif.GIF
	gifImage = *g.image
	return gif.EncodeAll(out, &gifImage) // TODO error (Exit due to fail resize[1]: gif: image block is out of bounds%
}

func getRatio() float64 {
		return 1
}

// Check if color is already in the Palette
func contains(colorPalette color.Palette, c color.Color) bool {
	for _, tmpColor := range colorPalette {
		if tmpColor == c {
			return true
		}
	}
	return false
}