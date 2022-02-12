package img

import (
	"project/zj"
	"time"
)

// PIO ...
type PIO struct {
	i    int
	echo bool
}

func (pio *PIO) Write(p []byte) (n int, err error) {

	if !pio.echo {
		time.Sleep(time.Hour)
	}

	pio.i++

	zj.J(pio.i, len(p))

	if pio.echo {
		zj.J(string(p))
	}

	return len(p), nil
}
