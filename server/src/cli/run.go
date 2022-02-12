package cli

import (
	"project/video"
	"project/zj"
)

// Run ...
func Run() {

	err := Init()
	if err != nil {
		zj.W(err)
		return
	}

	run()
}

func run() {
	v := video.Play(FileName)

	for {
		f, ok := v.Frame()
		if f != nil {
			zj.J(`read done`, f.Serial, len(f.Data), ok)
		}
		if !ok {
			break
		}
	}
}
