package server

import (
	"bytes"
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

	go func() {
		connWrite(c, s)
		wg.Done()
	}()
	go func() {
		connRead(c, s)
		wg.Done()
	}()

	wg.Wait()
	t := time.Now().Sub(start)
	t -= t % time.Millisecond
	zj.Access(c.RemoteAddr().String(), t)
}

func connWrite(c net.Conn, s *img.Stream) {

	for {

		ab, ok := s.Read()
		if !ok {
			break
		}

		_, err := c.Write(ab)
		if err != nil {
			break
		}
	}

	zj.J(`read`)

	c.Close()
}

func connRead(c net.Conn, s *img.Stream) {

	ab := make([]byte, 1024)

	for {
		n, err := c.Read(ab)
		if err != nil {
			break
		}
		if n == 5 && bytes.Compare(ab[:n], ctrlC) == 0 {
			break
		}
	}
	zj.J(`write`)

	c.Close()
}
