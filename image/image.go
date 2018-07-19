package image

import (
	"errors"
	"image"
	"image/gif"
	"bytes"
	"io"
)

type Image interface {
	Resize(width, height uint) error
}

func New(in io.Reader) (Image, error) {
	buf := bytes.NewBuffer(nil)
	r := io.TeeReader(in, buf)

	// formatによって処理を切り分け
	img, format, err := image.Decode(r)
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
		gifimg, err := gif.DecodeAll(buf)
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
