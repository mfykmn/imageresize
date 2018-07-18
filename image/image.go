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

func New(in io.Reader) (Image, error) {
	buf := new(bytes.Buffer)
	gifDecodeTarget := new(bytes.Buffer)

	w := io.MultiWriter(buf, gifDecodeTarget)
	io.Copy(w, in)

	// formatによって処理を切り分け
	img, format, err := image.Decode(buf)
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
		img, err := gif.DecodeAll(gifDecodeTarget)
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
