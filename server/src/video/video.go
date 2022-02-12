package video

import (
	"bytes"
	"fmt"
	"os/exec"
	"time"
)

// Video ...
type Video struct {
	cmd         *exec.Cmd
	ch          chan *Frame
	stdin       bytes.Buffer
	interval    time.Duration
	frameSerial int
	skip        int
}

// receive ffmpeg pipe
func (v *Video) Write(p []byte) (n int, err error) {

	n = len(p)
	// zj.J(`write`, v.frameSerial, len(p), v.interval)

	f := &Frame{
		Serial: v.frameSerial,
		Data:   p,
		Skip:   v.skip,
	}
	v.frameSerial++

	select {
	case v.ch <- f:
	default:
		v.skip++
	}

	return
}

func (v *Video) exec(file string, width, height int) {

	cmd := exec.Command(
		`ffmpeg`,
		`-i`,
		file,
		// `-vframes`, // 只生成前 n 帧，开发用
		// `10`,
		`-vf`,
		fmt.Sprintf(`scale=w=%d:h=%d:force_original_aspect_ratio=decrease,realtime`, width, height),
		`-vcodec`,
		`bmp`,
		`-f`,
		`image2pipe`,
		`pipe:1`,
	)
	cmd.Stdin = &v.stdin
	cmd.Stdout = v
	// cmd.Stderr = zj.ErrCfg

	v.cmd = cmd
	v.ch = make(chan *Frame, 1)
	v.interval = time.Second / 60

	go func() {
		cmd.Start()
		cmd.Wait()
		time.Sleep(time.Second)
		close(v.ch)
	}()
}

// Frame ...
func (v *Video) Frame() (f *Frame, ok bool) {
	f, ok = <-v.ch
	return
}
