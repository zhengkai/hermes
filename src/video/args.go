package video

import (
	"fmt"
	"strconv"
)

func makeArgs(file string, width, height, firstFrames int, seek string) (args []string) {

	if seek != `` {
		args = append(args,
			`-ss`,
			seek,
		)
	}
	args = append(args,
		`-i`,
		file,
	)
	if firstFrames > 0 { // 只生成前 n 帧
		args = append(args,
			`-vframes`,
			strconv.Itoa(firstFrames),
		)
	}

	vf := fmt.Sprintf(
		`scale=w=%d:h=%d:force_original_aspect_ratio=decrease,realtime`,
		width,
		height,
	)

	args = append(args,
		`-vf`, vf,
		`-vcodec`, `bmp`,
		`-f`, `image2pipe`,
		`pipe:1`,
	)
	return
}
