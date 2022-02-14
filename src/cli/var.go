package cli

// var
var (
	FileName string
	Size     string
	Verbose  bool

	sizeW   int
	sizeH   int
	sizeSet bool
)

const (
	exitFilename = iota + 1
	exitFlag
	exitSize
	exitFFmpeg
)
