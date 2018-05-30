package image

import (
	"image"
	"image/jpeg"
	"os"

	"github.com/nfnt/resize"
)

// 1〜100のスケールで画質を選択できる
const quality = 100

type JpegService struct {
	image image.Image
}

func (j *JpegService) Resize(width, height uint) error {
	m := resize.Resize(width, height, j.image, resize.Lanczos3)

	out, err := os.Create("resized." + Jpeg.String())
	if err != nil {
		return err
	}
	defer out.Close()

	opts := &jpeg.Options{Quality: quality}

	return jpeg.Encode(out, m, opts)
}
