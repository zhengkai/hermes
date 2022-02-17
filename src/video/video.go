package video

import (
	"bytes"
	"encoding/binary"
	"os/exec"
	"time"
)

// Video ...
type Video struct {
	cmd         *exec.Cmd
	ch          chan *Frame
	stdin       bytes.Buffer
	time        time.Time
	frameSerial int
	skip        int
	pipe        bytes.Buffer
	filesize    int
	Err         error
}

// receive ffmpeg pipe
func (v *Video) Write(p []byte) (n int, err error) {

	n = len(p)

	if v.chunkWrite(&p, n) {
		return
	}

	f := &Frame{
		Serial: v.frameSerial,
		Data:   p,
		Skip:   v.skip,
	}
	if v.frameSerial == 0 {
		v.time = time.Now()
	} else {
		f.Duration = time.Now().Sub(v.time)
	}
	v.frameSerial++

	select {
	case v.ch <- f:
	default:
		v.skip++
	}

	return
}

// 如果单张 bmp 文件大于上限、需要多次写入（我不知道 32K 是 pipeline 还是 ffmpeg 的上限）
func (v *Video) chunkWrite(p *[]byte, size int) (chunk bool) {

	ab := *p

	if v.filesize == 0 {
		fsize := int(binary.LittleEndian.Uint32(ab[2:6]))
		if fsize == size {
			return
		}
		v.filesize = fsize
	}

	v.pipe.Write(ab)
	if v.pipe.Len() < v.filesize {
		chunk = true
		return
	}

	if v.pipe.Len() == v.filesize {
		ab = v.pipe.Bytes()
		v.pipe.Reset()
	} else {
		ab = v.pipe.Next(v.filesize)
	}
	*p = ab
	v.filesize = 0
	return
}

// Frame ...
func (v *Video) Frame() (f *Frame, ok bool) {
	f, ok = <-v.ch
	return
}
