package img

import (
	"bytes"
	"fmt"
	"image"
	"io"

	// bmp
	_ "golang.org/x/image/bmp"
)

func fill(in io.Reader, out *bytes.Buffer) {

	im, _, err := image.Decode(in)
	if err != nil {
		return
	}

	rect := im.Bounds()

	for y := rect.Min.Y; y <= rect.Max.Y; y += 2 {
		for x := rect.Min.X; x <= rect.Max.X; x++ {
			br, bg, bb, _ := im.At(x, y).RGBA()
			fr, fg, fb, _ := im.At(x, y+1).RGBA()
			fmt.Fprintf(out, colorHead, uint8(fr), uint8(fg), uint8(fb), uint8(br), uint8(bg), uint8(bb))
		}
		out.WriteString(colorEnd)
		out.WriteByte('\n')
	}
}
