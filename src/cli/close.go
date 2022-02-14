package cli

import (
	"os"
	"os/signal"
	"syscall"
)

var isClose bool
var closeSignal os.Signal

func checkClose() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT)
	closeSignal = <-c
	isClose = true
}
