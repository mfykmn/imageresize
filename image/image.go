package image

import (
	"os"
	"image"
	"image/png"

	"github.com/nfnt/resize"
)

func New(filepath string) *Image {
	return &Image{
		filepath: filepath,
	}
}

type Image struct {
	filepath string
}

func(i *Image) Resize(width, height uint) error {
	// 画像ファイルを開く
	file, err := os.Open(i.filepath)
	if err != nil {
		return err
	}
	defer file.Close()


	// デコードしてイメージオブジェクトを準備
	img, _, err := image.Decode(file)
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
