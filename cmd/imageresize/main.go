package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mafuyuk/imageresize"
	"github.com/mafuyuk/imageresize/image"

	"github.com/go-ozzo/ozzo-validation"
)

type option struct {
	filepath string
	width    int
	height   int
}

func (o option) Validate() error {
	return validation.ValidateStruct(&o,
		validation.Field(&o.filepath, validation.Required),
		validation.Field(&o.width, validation.Min(0)),
		validation.Field(&o.height, validation.Min(0)),
	)
}

func main() {
	var opt = &option{}
	flag.StringVar(&opt.filepath, "f", "", "File path of the image you want to change")
	flag.IntVar(&opt.width, "w", 100, "")
	flag.IntVar(&opt.height, "h", 100, "")
	flag.Parse()

	if err := opt.Validate(); err != nil {
		fmt.Printf("Exit due to option error[%d]: %s", imageresize.ExitCodeError, err.Error())
		os.Exit(imageresize.ExitCodeError)
	}

	os.Exit(Run(opt))
}

func Run(option *option) int {
	// ファイルオープン
	file, err := os.Open(option.filepath)
	if err != nil {
		return imageresize.ExitCodeFileError
	}
	defer file.Close()

	// 画像オブジェクトの取得
	i, err := image.New(file)
	if err != nil {
		fmt.Println(err.Error())
		return imageresize.ExitCodeError
	}

	// リサイズ
	if err := i.Resize(uint(option.width), uint(option.height)); nil != err {
		fmt.Printf("Exit due to fail resize[%d]: %s", imageresize.ExitCodeError, err.Error())
		return imageresize.ExitCodeError
	}

	return imageresize.ExitCodeOk
}
