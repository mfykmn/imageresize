package image

import (
	"os"
	"image"
	"github.com/nfnt/resize"
	"image/png"
)

const format = ".png"

type Png struct {
	image image.Image
}

func(p Png) Resize(width, height uint) error {
	// リサイズ
	m := resize.Resize(width, height, p.image, resize.Lanczos3)

	out, err := os.Create("resized"+format)
	if err != nil {
		return err
	}
	defer out.Close()

	return png.Encode(out, m)
}
