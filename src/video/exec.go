package video

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
)

func (v *Video) exec(file string, width, height, firstFrames int, seek string) {

	vf := fmt.Sprintf(`scale=w=%d:h=%d:force_original_aspect_ratio=decrease,realtime`, width, height)

	var arg []string

	if seek != `` {
		arg = append(arg,
			`-ss`,
			seek,
		)
	}
	arg = append(arg,
		`-i`,
		file,
	)
	if firstFrames > 0 { // 只生成前 n 帧
		arg = append(arg,
			`-vframes`,
			strconv.Itoa(firstFrames),
		)
	}

	arg = append(arg,
		`-vf`, vf,
		`-vcodec`, `bmp`,
		`-f`, `image2pipe`,
		`pipe:1`,
	)

	cmd := exec.Command(`ffmpeg`, arg...)

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
