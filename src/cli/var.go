package cli

// var
var (
	FileName  string
	Size      string
	Verbose   bool
	Version   bool
	Seek      string
	finalSeek string

	FirstFrames int

	sizeW   int
	sizeH   int
	sizeSet bool
)

const (
	exitFilename = iota + 1
	exitFlag
	exitSize
	exitFFmpeg
	exitSeek
)
