package img

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"io"

	// bmp
	_ "golang.org/x/image/bmp"
)

func fill(in io.Reader, out *bytes.Buffer) (rect image.Rectangle) {

	im, _, err := image.Decode(in)
	if err != nil {
		return
	}

	rect = im.Bounds()

	var skip, skipF, skipB int
	for y := rect.Min.Y; y < rect.Max.Y; y += 2 {

		var prevB, prevF color.RGBA

		for x := rect.Min.X; x < rect.Max.X; x++ {

			b := im.At(x, y).(color.RGBA)
			f := im.At(x, y+1).(color.RGBA)

			if prevB == b {

				if prevF == f {
					skip++
					out.WriteString(`▄`)
					continue
				}

				prevF = f
				skipB++
				fmt.Fprintf(out, colorFront, f.R, f.G, f.B)
				continue
			}

			if prevF == f {
				prevB = b
				skipF++
				fmt.Fprintf(out, colorBack, b.R, b.G, b.B)
				continue
			}

			fmt.Fprintf(out, colorHead, f.R, f.G, f.B, b.R, b.G, b.B)
			prevB = b
			prevF = f
		}
		out.WriteString(colorEnd)
		out.WriteByte('\n')
	}

	if skip > 0 || skipB > 0 || skipF > 0 {
		// zj.IOF(`skip %5d %5d %5d`, skip, skipB, skipF)
	}

	return
}
