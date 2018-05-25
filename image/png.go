package image

import (
	"os"
	"image"
	"github.com/nfnt/resize"
	"image/png"
)

type Png struct {
	Image image.Image
}

func(p Png) Resize(width, height uint) error {
	// リサイズ
	m := resize.Resize(width, height, p.Image, resize.Lanczos3)

	out, err := os.Create("resize.png")
	if err != nil {
		return err
	}
	defer out.Close()

	return png.Encode(out, m)
}
