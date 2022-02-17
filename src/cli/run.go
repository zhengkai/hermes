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

	flagInit()

	run()
}

func run() {

	w := sizeW
	h := sizeH
	if !sizeSet {
		w, h, _ = term.GetSize(0)
		if Verbose {
			h--
		}
		h *= 2
	}

	if Verbose {
		if FirstFrames > 0 {
			fmt.Fprintln(os.Stderr, `-frames`, FirstFrames)
		}
		fmt.Fprintf(os.Stderr, "-size %dx%d\n", w, h)
		if finalSeek != `` {
			fmt.Fprintln(os.Stderr, `-seek`, finalSeek)
		}
	}

	v := video.Play(FileName, w, h, FirstFrames, finalSeek)

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
			if h == 0 {
				fmt.Fprint(os.Stdout, "\033[0G")
			} else {
				fmt.Fprintf(os.Stdout, "\033[%dF", h)
			}
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

	if !first {
		fmt.Println("\033[0m")
	}

	fmt.Fprintln(os.Stderr, v.Err)

	if isClose {
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
