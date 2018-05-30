package image

import (
	"image"
	"image/png"
	"os"

	"github.com/nfnt/resize"
)

type PngService struct {
	image image.Image
}

func (p PngService) Resize(width, height uint) error {
	m := resize.Resize(width, height, p.image, resize.Lanczos3)

	out, err := os.Create("resized." + Png.String())
	if err != nil {
		return err
	}
	defer out.Close()

	return png.Encode(out, m)
}
