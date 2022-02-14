package cli

import (
	"bytes"
	"fmt"
	"os"
	"project/img"
	"project/video"

	"golang.org/x/term"
)

// Run ...
func Run() {

	Init()

	run()
}

func run() {

	var w, h int
	if sizeSet {
		w = int(sizeW)
		h = int(sizeH)
	} else {
		w, h, _ = term.GetSize(0)
		if Verbose {
			h--
		}
		h *= 2
	}

	v := video.Play(FileName, w, h)

	var buf *bytes.Buffer

	first := true

	go checkClose()

	for {
		f, ok := v.Frame()
		if !ok || f == nil || isClose {
			break
		}

		if buf == nil {
			ab := make([]byte, 0, len(f.Data)*5)
			buf = bytes.NewBuffer(ab)
		} else {
			buf.Reset()
		}

		// TODO: 优化空间：应比较前后两帧的差异，只改写变化的地方
		rect := img.Fill(bytes.NewReader(f.Data), buf)

		if first {
			first = false
		} else {
			h := (rect.Max.Y - 1) / 2
			if Verbose {
				h++
			}
			fmt.Fprintf(os.Stdout, "\033[%dF", h)
		}
		if Verbose {
			verboseLine(f, rect, buf.Len())
		}

		if isClose {
			break
		}
		buf.WriteTo(os.Stdout)
		if isClose {
			break
		}
	}

	if isClose {
		fmt.Println("\033[0m")
		if closeSignal == os.Interrupt {
			os.Stderr.WriteString("Ctrl+C pressed in Terminal\n")
		} else {
			fmt.Fprintln(os.Stderr, closeSignal, `signal received, stop`)
		}
	}

	if Verbose {
		verboseSummary()
	}
}
