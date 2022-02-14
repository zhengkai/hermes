package video

import "time"

// Frame ...
type Frame struct {
	Serial   int
	Data     []byte
	Skip     int
	Duration time.Duration
}
