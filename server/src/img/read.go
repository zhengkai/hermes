package img

import (
	"bytes"
	"os"
)

const colorHead = "\033[38;2;%d;%d;%d;48;2;%d;%d;%dm▄"
const colorFront = "\033[38;2;%d;%d;%dm▄"
const colorBack = "\033[48;2;%d;%d;%dm▄"
const colorEnd = "\033[0m"
const colorRight = "\033[%dC"

// Read ...
func Read(file string) (err error) {

	ab, err := os.ReadFile(file)
	if err != nil {
		return
	}

	var out bytes.Buffer

	in := bytes.NewReader(ab)

	fill(in, &out)

	write(`/tmp/3.txt`, &out)

	return
}
