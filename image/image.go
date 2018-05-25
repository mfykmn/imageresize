package image

import (
	"os"
	"image"
	"errors"
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
	case "png":
		return &Png{
			image: img,
		}, nil
	case "jpg":
		return &Jpg{}, nil
	case "gif":
		return &Gif{}, nil
	default:
		return nil, errors.New("image unknown")
	}
}
