package cli

import (
	_ "embed" //
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"project/build"

	"github.com/zhengkai/zu"
)

//go:embed help.txt
var helpText string

// flag ...
func flagInit() {

	flag.CommandLine.SetOutput(os.Stderr)
	flag.Usage = func() {
		name := `Hermes`
		if build.BuildGit != `` {
			name += ` ` + build.BuildGit
		}
		fmt.Fprint(os.Stderr, name, "\n\n")
		fmt.Fprintf(os.Stderr, helpText, filepath.Base(os.Args[0]))
		os.Stderr.WriteString("\n\n")
		flag.PrintDefaults()
	}

	var defaultSize string
	w, h, ok := parseSize(os.Getenv(`HERMES_DEFAULT_SIZE`))
	if ok {
		defaultSize = fmt.Sprintf(`%dx%d`, w, h)
	}

	flag.StringVar(&Size, `size`, defaultSize, "Output size, example: \"80x40\"\nnote: one character can display two pixels height")
	flag.IntVar(&FirstFrames, `frames`, 0, "Show first N frames only")
	flag.BoolVar(&Verbose, `verbose`, false, "Print verbose information")
	flag.BoolVar(&Version, `version`, false, "Print version")
	flag.StringVar(&Seek, `seek`, ``, `Seek to the closest seek point before position, like "-ss" in ffmpeg`)

	flag.Parse()

	if Version {
		build.Ver()
		os.Exit(0)
	}

	a := flag.Args()
	if len(a) > 0 {
		FileName = a[0]
	}

	if FileName == `` {
		flag.Usage()
		os.Exit(exitFilename)
	}
	if !zu.FileExists(FileName) {
		fmt.Fprintf(os.Stderr, "file \"%s\" not found\n", FileName)
		os.Exit(exitFilename)
	}

	if Size != `` {
		sizeW, sizeH, sizeSet = parseSize(Size)
		if !sizeSet {
			fmt.Fprintf(os.Stderr, "invalid size \"%s\"\n", Size)
			os.Exit(exitSize)
		}
	}

	finalSeek, ok = parseSeek(Seek)
	if !ok {
		fmt.Fprintf(os.Stderr, "invalid seek \"%s\"\n", Seek)
		os.Exit(exitSize)
	}
}
