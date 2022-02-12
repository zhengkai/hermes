package project

import (
	"project/cli"
	"project/config"
)

// Start ...
func Start() {

	cli.Run()
}

// Prod ...
func Prod() {

	config.Prod = true

	Start()
}
