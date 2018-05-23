package image

import (
	"os"
	"image"
)

type Image interface {
	Resize(width, height uint) error
}

func New(f *os.File) (Image, error) {
	_, format, err := image.DecodeConfig(f)
	if err != nil {
		// 画像フォーマットではない場合はエラーが発生する
		return Image, err
	}




	switch format {
	case "png":
		return &Png{
			File: f,
		}, nil
	case "jpg":
		return &Jpg{}, nil
	case "gif":
		return &Gif{}, nil
	default:
		//TODO
		return Image, err
	}
}
