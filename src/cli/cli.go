package cli

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/zhengkai/zu"
)

// Init ...
func Init() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage:\n\n  %s [-s 80x40] video_filename\n\n", filepath.Base(os.Args[0]))
		flag.PrintDefaults()
	}

	flag.StringVar(&Size, `s`, ``, "Output size, example: \"80x40\"\nnote: one character can display two pixels height")

	flag.Parse()

	a := flag.Args()
	if len(a) > 0 {
		FileName = a[0]
	}

	if FileName == `` {
		flag.Usage()
		os.Exit(1)
		return
	}

	if Size != `` {
		size := strings.SplitN(Size, `x`, 3)
		if len(size) < 2 {
			sizeError()
		}
		var err error
		sizeW, err = strconv.ParseInt(strings.TrimSpace(size[0]), 10, 20)
		if err != nil {
			sizeError()
		}
		sizeH, err = strconv.ParseInt(strings.TrimSpace(size[1]), 10, 20)
		if err != nil {
			sizeError()
		}
		sizeSet = true
	}

	if !zu.FileExists(FileName) {
		fmt.Fprintf(os.Stderr, "file \"%s\" not found\n", FileName)
		os.Exit(3)
	}
}

func sizeError() {
	fmt.Fprintf(os.Stderr, "size \"%s\" illegal\n", Size)
	os.Exit(4)
}
