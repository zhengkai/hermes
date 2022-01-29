package img

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"project/zj"
	"time"
)

// Movie ...
func Movie(file string) (err error) {

	start := time.Now()
	max := time.Second * 10

	for {
		d := time.Now().Sub(start)

		if d > max {
			break
		}

		movieFrame(file, d+time.Second*15)
		time.Sleep(time.Second / 20)
	}

	return
}

func movieFrame(file string, t time.Duration) (err error) {

	ft := float64(t / time.Millisecond)
	s := fmt.Sprintf(`%.03f`, ft/1000)

	cmd := exec.Command(
		`ffmpeg`,
		`-ss`,
		s,
		`-i`,
		file,
		`-vframes`,
		`1`,
		`-vf`,
		`scale=w=120:h=80:force_original_aspect_ratio=decrease`,
		`-vcodec`,
		`bmp`,
		`-f`,
		`image2pipe`,
		`pipe:1`,
	)

	var pic bytes.Buffer
	// cmd.Stderr = os.Stdout
	cmd.Stdout = &pic
	err = cmd.Run()
	if err != nil {
		zj.W(`ffmpeg fail:`, err)
		return
	}

	var out bytes.Buffer
	rect := fill(&pic, &out)

	h := (rect.Max.Y + 1) / 2
	// zj.IO(rect, h)

	os.Stdout.Write([]byte(fmt.Sprintf("\033[%dF", h)))
	out.WriteTo(os.Stdout)
	// write(`/tmp/4.txt`, &out)
	return
}
