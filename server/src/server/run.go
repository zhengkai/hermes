package server

import (
	"net"
	"project/config"
	"project/zj"
)

// Run ...
func Run() {

	ln, err := net.Listen(`tcp`, config.Port)
	if err != nil {
		zj.W(`listen start fail:`, err)
		return
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			zj.W(`listen fail:`, err)
			return
		}
		go handleConnection(c)
	}
}
