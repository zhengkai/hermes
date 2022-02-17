package video

import (
	"bytes"
	"os/exec"
)

func (v *Video) exec(file string, width, height, firstFrames int, seek string) {

	args := makeArgs(file, width, height, firstFrames, seek)

	cmd := exec.Command(`ffmpeg`, args...)

	cmd.Stdin = &v.stdin
	cmd.Stdout = v

	var errOut bytes.Buffer
	cmd.Stderr = &errOut

	v.cmd = cmd
	v.ch = make(chan *Frame, 1)

	go func() {
		v.Err = cmd.Start()
		cmd.Wait()
		close(v.ch)
	}()
}
