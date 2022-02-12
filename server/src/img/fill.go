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

	var prevB, prevF color.RGBA

	var skip, skipF, skipB, skipOne int
	for y := rect.Min.Y; y < rect.Max.Y; y += 2 {

		for x := rect.Min.X; x < rect.Max.X; x++ {

			b := im.At(x, y).(color.RGBA)
			f := im.At(x, y+1).(color.RGBA)

			if prevB == b {

				if f == b {
					skipOne++
					out.WriteString(` `)
					continue
				}

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
				skipF++
				prevB = b
				fmt.Fprintf(out, colorBack, b.R, b.G, b.B)
				continue
			}

			if f == b {
				skipOne++
				prevB = b
				fmt.Fprintf(out, colorBackOne, b.R, b.G, b.B)
				continue
			}

			fmt.Fprintf(out, colorHead, f.R, f.G, f.B, b.R, b.G, b.B)
			prevB = b
			prevF = f
		}
		if y+2 < rect.Max.Y {
			out.WriteByte('\n')
		}
	}
	out.WriteString(colorEnd)

	// zj.IOF(`skip %5d %5d %5d %5d`, skip, skipB, skipF, skipOne)

	return
}
