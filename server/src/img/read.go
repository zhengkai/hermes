package img

import (
	"bytes"
	"os"
)

const (
	colorHead    = "\033[38;2;%d;%d;%d;48;2;%d;%d;%dm▄"
	colorFront   = "\033[38;2;%d;%d;%dm▄"
	colorBack    = "\033[48;2;%d;%d;%dm▄"
	colorBackOne = "\033[48;2;%d;%d;%dm "
	colorRight   = "\033[%dC"
)

var (
	colorEnd   = []byte("\033[0m")
	colorDot   = []byte(`▄`)
	colorSpace = []byte{' '}
	colorBR    = []byte{'\n'}
)

// Read ...
func Read(file string) (err error) {

	ab, err := os.ReadFile(file)
	if err != nil {
		return
	}

	var out bytes.Buffer

	in := bytes.NewReader(ab)

	Fill(in, &out)

	write(`/tmp/3.txt`, &out)

	return
}
