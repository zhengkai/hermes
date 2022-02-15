package cli

import (
	"fmt"
	"image"
	"os"
	"project/video"
	"time"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var ansiLenSum int
var frameCount int
var lastDur time.Duration

func verboseLine(f *video.Frame, rect image.Rectangle, ansiLen int) {
	t := time.Unix(0, 0).UTC()
	dur := t.Add(f.Duration).Format(`15:04:05.000`)
	lastDur = f.Duration
	var fps float64
	if f.Duration > 0 {
		fps = float64(time.Second) * float64(f.Serial-f.Skip) / float64(f.Duration)
	}
	fmt.Fprintf(
		os.Stderr,
		"%s %dx%d, fps:%.2f, frame:%4d, skip:%d, bmp:%d, ansi:%d\n",
		dur,
		rect.Max.X,
		rect.Max.Y,
		fps,
		f.Serial,
		f.Skip,
		len(f.Data),
		ansiLen,
	)

	if ansiLenSum > 0 {
		ansiLenSum += 4
	}
	h := (rect.Max.Y-1)/2 + 1
	if h >= 10 {
		ansiLenSum++
		if h >= 100 {
			ansiLenSum++
		}
	}
	ansiLenSum += ansiLen

	frameCount++
}

func verboseSummary() {
	p := message.NewPrinter(language.English)

	bps := float64(ansiLenSum) / (float64(lastDur) / float64(time.Second))

	p.Fprintf(os.Stderr, "total output: %d\n", ansiLenSum)
	p.Fprintf(os.Stderr, "   pre frame: %d\n", ansiLenSum/frameCount)
	p.Fprintf(os.Stderr, "  pre second: %d\n", int(bps))
}
