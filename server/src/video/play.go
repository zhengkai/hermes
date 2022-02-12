package video

// Play ...
func Play(file string) (v *Video) {

	fps := GetFPS(file)
	if fps == 0 {
		fps = 30
	}

	v = &Video{}
	v.exec(file)

	return v
}
