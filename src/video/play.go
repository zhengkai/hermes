package video

// Play ...
func Play(file string, width, height, firstFrames int, seek string) (v *Video) {

	if width < 1 {
		width = 1
	}
	if height < 1 {
		height = 1
	}

	v = &Video{}
	v.exec(file, width, height, firstFrames, seek)

	return v
}
