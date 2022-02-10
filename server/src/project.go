package project

import (
	"project/config"
	"project/server"
)

// Start ...
func Start() {

	server.Run()
}

// Prod ...
func Prod() {

	config.Prod = true

	Start()
}
