package image

import (
	"errors"
	"image"
	"bytes"
	"io"
	"image/gif"
	"io/ioutil"
)

type Image interface {
	Resize(width, height uint) error
}

func New(file io.Reader) (Image, error) {
	buf, _ := ioutil.ReadAll(file)
	b0 := bytes.NewBufferString(string(buf))
	b1 := bytes.NewBufferString(string(buf))

	// formatによって処理を切り分け
	img, format, err := image.Decode(b0)
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
		img, err := gif.DecodeAll(b1)
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
