package image

import (
	//"os"
	"image/gif"

	//"github.com/nfnt/resize"
	"fmt"
)

type GifService struct {
	image *gif.GIF
}

func (g *GifService) Resize(width, height uint) error {
	//m := resize.Resize(width, height, g.image, resize.Lanczos3)

	for index, frame := range g.image.Image {
		fmt.Print(index)
		fmt.Print(frame)
	}
	return nil

	//out, err := os.Create("resized." + Gif.String())
	//if err != nil {
	//	return err
	//}
	//defer out.Close()
	//
	//return gif.EncodeAll(out, g.image)
}
