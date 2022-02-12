package video

import (
	"bytes"
	"os/exec"
	"regexp"
	"strconv"
)

var regexpFPS = regexp.MustCompile(`, ([0-9\.]+) fps`)

// GetFPS ...
func GetFPS(file string) (fps float64) {

	cmd := exec.Command(
		`ffmpeg`,
		`-i`,
		file,
	)

	var out bytes.Buffer
	cmd.Stderr = &out

	cmd.Run()

	re := regexpFPS.FindSubmatch(out.Bytes())
	if len(re) != 2 {
		return
	}

	fps, _ = strconv.ParseFloat(string(re[1]), 64)
	return
}
