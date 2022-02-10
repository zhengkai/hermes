package server

import (
	"net"
	"project/zj"
)

// Run ...
func Run() {

	ln, err := net.Listen(`tcp`, `:23`)
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
