package image

import (
	"errors"
	"image"
	"bytes"
	"io"
	"image/gif"
)

type Image interface {
	Resize(width, height uint) error
}

func New(file io.Reader) (Image, error) {
	buf := new(bytes.Buffer)
	buf2 := new(bytes.Buffer)
	io.Copy(buf, file)
	io.Copy(buf2, file)

	// formatによって処理を切り分け
	img, format, err := image.Decode(buf)
	if err != nil {
		return nil, err
	}
	_, _, err = image.Decode(buf2)
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
		img, err := gif.DecodeAll(buf)
		if err != nil {
			return nil, err
		}
		return &GifService{
			image: img,
		}, nil
	default:
		return nil, errors.New("image unknown")
	}
}
