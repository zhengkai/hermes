package project

import (
	"project/cli"
	"project/config"
	"project/zj"
)

// Start ...
func Start() {

	zj.Init()

	cli.Run()
}

// Prod ...
func Prod() {

	config.Prod = true

	Start()
}
