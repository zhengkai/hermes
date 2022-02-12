package video

import (
	"bytes"
	"os/exec"
	"project/zj"
	"time"
)

// Video ...
type Video struct {
	cmd         *exec.Cmd
	ch          chan *Frame
	stdin       bytes.Buffer
	nextTime    time.Time
	interval    time.Duration
	frameSerial int
	skip        int
}

// receive ffmpeg pipe
func (v *Video) Write(p []byte) (n int, err error) {

	n = len(p)
	// zj.J(`write`, v.frameSerial, len(p), v.interval)

	now := time.Now()
	if v.frameSerial == 0 {
		v.nextTime = now
	} else {
		diff := v.nextTime.Sub(now)
		// zj.J(`write diff`, diff, now.Format(`15:04:05.000`), v.nextTime.Format(`15:04:05.000`))
		if diff > 0 {
			time.Sleep(diff)
		}
	}

	f := &Frame{
		Serial: v.frameSerial,
		Data:   p,
	}
	v.frameSerial++
	v.nextTime = v.nextTime.Add(v.interval)

	select {
	case v.ch <- f:
	default:
		v.skip++
	}

	return
}

func (v *Video) exec(file string) {

	cmd := exec.Command(
		`ffmpeg`,
		`-i`,
		file,
		`-vframes`,
		`20`,
		`-vf`,
		`scale=w=120:h=80:force_original_aspect_ratio=decrease,realtime`,
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
	zj.J(`skip`, v.skip)
	return
}
