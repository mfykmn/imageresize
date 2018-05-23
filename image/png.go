package image

import (
	"os"
	"image"
	"github.com/nfnt/resize"
	"image/png"
)

type Png struct {
	File *os.File
}

func(p *Png) Resize(width, height uint) error {
	// デコードしてイメージオブジェクトを準備
	img, _, err := image.Decode(p.File)
	if err != nil {
		return err
	}

	// リサイズ
	m := resize.Resize(width, height, img, resize.Lanczos3)

	out, err := os.Create("resize.png")
	if err != nil {
		return err
	}
	defer out.Close()

	return png.Encode(out, m)
}
