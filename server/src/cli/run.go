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
		h *= 2
	}

	v := video.Play(FileName, w, h)

	var buf *bytes.Buffer

	for {
		f, ok := v.Frame()
		if f != nil {

			if buf == nil {
				ab := make([]byte, 0, len(f.Data)*5)
				buf = bytes.NewBuffer(ab)
			} else {
				buf.Reset()
			}

			rect := img.Fill(bytes.NewReader(f.Data), buf)

			// zj.J(`read done`, ok, f.Serial, len(f.Data), f.Skip, f.Serial-f.Skip, buf.Len(), time.Now().Sub(now))

			fmt.Fprintln(os.Stdout, rect.Max.X, rect.Max.Y, f.Serial, f.Skip, len(f.Data), buf.Len())
			buf.WriteTo(os.Stdout)

			h := (rect.Max.Y - 1) / 2
			fmt.Fprintf(os.Stdout, "\033[%dF", h+1)
		}
		if !ok {
			break
		}
	}
}
