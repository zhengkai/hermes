package img

import (
	"bytes"
	"os"
)

func write(file string, out *bytes.Buffer) (err error) {

	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return
	}

	_, err = out.WriteTo(f)
	if err != nil {
		return
	}

	err = f.Close()
	if err != nil {
		return
	}
	return
}
