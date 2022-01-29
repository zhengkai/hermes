package project

import (
	"project/config"
	"project/img"
)

// Start ...
func Start() {
	// img.Read(`/tmp/1.bmp`)
	img.Movie(`/share/4.mp4`)
}

// Prod ...
func Prod() {

	config.Prod = true

	Start()
}
