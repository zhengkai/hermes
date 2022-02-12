package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/zhengkai/zu"
)

// Init ...
func Init() (err error) {

	flag.StringVar(&FileName, `f`, `/www/hermes/static/forza5.mp4`, `video file name`)

	flag.Parse()

	if !zu.FileExists(FileName) {
		fmt.Fprintf(os.Stderr, "file \"%s\" not found\n", FileName)
		err = os.ErrNotExist
		return
	}

	return
}
