package img

import (
	"bytes"
	"fmt"
	"os"
	"project/config"
	"time"
)

const preFrame = time.Second / 60

// Stream ...
type Stream struct {
	t           time.Time
	c           chan []byte
	stop        bool
	frameCursor time.Duration
}

// NewStream ...
func NewStream() (s *Stream) {
	s = &Stream{
		t: time.Now(),
		c: make(chan []byte),
	}
	go s.loop()
	return
}

func (s *Stream) Read() (ab []byte, ok bool) {
	ab, ok = <-s.c
	return
}

func (s *Stream) loop() {
	for {
		s.frame()
		if s.stop {
			break
		}
	}
	close(s.c)
}

func (s *Stream) frame() {
	now := time.Now()
	f := now.Sub(s.t) / preFrame
	// zj.J(now, f)

	if f > 1799 {
		s.stop = true
		return
	}

	if s.frameCursor == f {
		sl := preFrame - now.Sub(s.t)%preFrame
		time.Sleep(sl)
		f++
	}
	s.frameCursor = f

	// msg := fmt.Sprintf("%s %d\n", now, f)

	filename := fmt.Sprintf(`%s/bf30/%04d.bmp`, config.StaticDir, f)
	file, err := os.Open(filename)
	if err != nil {
		s.stop = true
		return
	}
	defer file.Close()

	var out bytes.Buffer
	rect := Fill(file, &out)
	h := (rect.Max.Y - 1) / 2
	fmt.Fprintf(&out, "\033[%dF", h)

	for {
		var ok bool
		select {
		case s.c <- out.Bytes():
			ok = true
		case <-time.After(3 * time.Second):
		}
		if ok || s.stop {
			break
		}
	}
}
