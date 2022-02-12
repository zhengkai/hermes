package video

// Play ...
func Play(file string, width, height int) (v *Video) {

	if width < 8 {
		width = 8
	}
	if height < 8 {
		height = 8
	}

	v = &Video{}
	v.exec(file, width, height)

	return v
}
