package img

import (
	"fmt"
	"image"
	"image/color"
	"io"

	// bmp
	_ "golang.org/x/image/bmp"
)

// Fill bmp binary to ansi color code
func Fill(in io.Reader, out io.Writer) (rect image.Rectangle) {

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
					out.Write(colorSpace)
					continue
				}

				if prevF == f {
					skip++
					out.Write(colorDot)
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
			out.Write(colorBR)
		}
	}
	out.Write(colorEnd)

	return
}
