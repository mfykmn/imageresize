package main

import (
	"fmt"
	"os"
	"flag"

	"github.com/mafuyuk/imageresize"

	"github.com/go-ozzo/ozzo-validation"
)


type option struct {
  filepath string
}

func (o option) Validate() error {
	return validation.ValidateStruct(&o,
		validation.Field(&o.filepath, validation.Required),
	)
}



func main() {
	var opt = &option{}
	flag.StringVar(&opt.filepath, "filepath", "", "File path of the image you want to change")
	flag.Parse()

	if err := opt.Validate(); err != nil {
		fmt.Printf("Exit due to option error[%d]: %s", imageresize.ExitCodeError, err.Error())
		os.Exit(imageresize.ExitCodeError)
	}
	fmt.Printf("%#v", opt)

	os.Exit(Run(os.Args))
}

func Run(args []string) int {
	fmt.Println(args)

	return imageresize.ExitCodeOk
}
