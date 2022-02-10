package project

import (
	"project/config"
	"project/server"
)

// Start ...
func Start() {
	// img.Read(`/tmp/1.bmp`)
	// img.Movie(`/share/4.mp4`)
	server.Run()
}

// Prod ...
func Prod() {

	config.Prod = true

	Start()
}
