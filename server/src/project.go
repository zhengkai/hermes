package project

import (
	"project/cli"
	"project/config"
	"project/zj"
	"time"
)

// Start ...
func Start() {

	cli.Run()

	zj.J(`all done`)
	time.Sleep(time.Hour)

	// img.TestMovieFrame(cli.FileName)

	// server.Run()
}

// Prod ...
func Prod() {

	config.Prod = true

	Start()
}
