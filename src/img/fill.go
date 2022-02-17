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

	var skip, skipF, skipB, skipOne int
	for y := rect.Min.Y; y < rect.Max.Y; y += 2 {

		var prevB, prevF color.RGBA
		for x := rect.Min.X; x < rect.Max.X; x++ {

			b := getPixel(im, x, y)
			f := getPixel(im, x, y+1)

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
		out.Write(colorEnd)
	}

	return
}

func getPixel(im image.Image, x, y int) (p color.RGBA) {

	v := im.At(x, y)
	switch c := v.(type) {
	case color.RGBA:
		p = c
	case color.NRGBA:
		p = color.RGBA{c.R, c.G, c.B, c.A}
	case color.Gray:
		p = color.RGBA{c.Y, c.Y, c.Y, 0}
	case color.Alpha:
		p = color.RGBA{c.A, c.A, c.A, 0}
	default:
		r, g, b, a := v.RGBA()
		p = color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
	}

	return
}
