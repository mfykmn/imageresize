package image

import (
	"errors"
	"image"
	"os"
)

type Image interface {
	Resize(width, height uint) error
}

func New(file *os.File) (Image, error) {
	img, format, err := image.Decode(file)
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
		return &GifService{}, nil
	default:
		return nil, errors.New("image unknown")
	}
}
