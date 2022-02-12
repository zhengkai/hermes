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

	fps := MovieInfo(file)
	if fps == 0 {
		fps = 30
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

// TestMovieFrame ...
func TestMovieFrame(file string) (err error) {

	cmd := exec.Command(
		`ffmpeg`,
		`-i`,
		file,
		`-vframes`,
		`2`,
		`-vf`,
		`scale=w=120:h=80:force_original_aspect_ratio=decrease`,
		`-vcodec`,
		`bmp`,
		`-f`,
		`image2pipe`,
		`pipe:1`,
	)

	o := &PIO{}

	cmd.Stdout = o

	var in bytes.Buffer
	cmd.Stdin = &in

	go func() {
		err = cmd.Start()
		if err != nil {
			zj.W(`ffmpeg fail:`, err)
			return
		}
	}()

	time.Sleep(time.Second)
	in.WriteString(`q`)

	return
}
