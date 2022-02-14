package cli

import (
	"strconv"
	"strings"
)

func parseSize(size string) (w, h int, ok bool) {

	s := strings.SplitN(size, `x`, 3)
	if len(s) < 2 {
		return
	}

	sw, err := strconv.ParseInt(strings.TrimSpace(s[0]), 10, 20)
	if err != nil {
		return
	}

	sh, err := strconv.ParseInt(strings.TrimSpace(s[1]), 10, 20)
	if err != nil {
		return
	}

	w = int(sw)
	h = int(sh)

	ok = true
	return
}
