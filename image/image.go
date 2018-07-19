package image

import (
	"errors"
	"image"
	"image/gif"
	"os"
)

type Image interface {
	Resize(width, height uint) error
}

func New(in *os.File) (Image, error) {
	// formatによって処理を切り分け
	img, format, err := image.Decode(in)
	if err != nil {
		return nil, err
	}

	switch format {
	case Png.String():
		return &PngService{
			image: img,
		}, nil
	case Jpeg.String():
		return &JpegService{
			image: img,
		}, nil
	case Gif.String():
		in.Seek(0, 0) // インデックスを先頭に戻す
		gifimg, err := gif.DecodeAll(in)
		if err != nil {
			return nil, err
		}
		return &GifService{
			image: gifimg,
		}, nil
	default:
		return nil, errors.New("image unknown")
	}
}
