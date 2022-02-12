package server

import (
	"bytes"
	"fmt"
	"net"
	"project/img"
	"project/zj"
	"sync"
	"time"
)

var ctrlC = []byte{0xff, 0xf4, 0xff, 0xfd, 0x06}

func handleConnection(c net.Conn) {

	start := time.Now()

	var wg sync.WaitGroup
	wg.Add(2)

	s := img.NewStream()

	var in, out, inCount, outCount int
	go func() {
		out, outCount = connWrite(c, s)
		wg.Done()
	}()
	go func() {
		in, inCount = connRead(c, s)
		wg.Done()
	}()

	wg.Wait()
	t := time.Now().Sub(start)
	t -= t % time.Millisecond
	zj.Access(fmt.Sprintf(`%21s %10s %8d %10d %8d %8d`, c.RemoteAddr().String(), t, in, out, inCount, outCount))
}

func connWrite(c net.Conn, s *img.Stream) (size, count int) {

	for {

		ab, ok := s.Read()
		if !ok {
			break
		}

		n, err := c.Write(ab)
		size += n
		count++
		if err != nil {
			break
		}
	}

	c.Close()
	return
}

func connRead(c net.Conn, s *img.Stream) (size, count int) {

	ab := make([]byte, 1024)

	for {
		n, err := c.Read(ab)
		size += n
		count++
		if err != nil {
			break
		}
		if bytes.Contains(ab[:n], ctrlC) {
			break
		}
	}

	c.Close()
	return
}
