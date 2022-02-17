package cli

import (
	"fmt"
	"project/zj"
	"regexp"
	"strconv"
	"strings"
)

var regexpSeek = regexp.MustCompile(`^(-)?(((\d{1,3}):)?((\d{1,3}):))?(\d{1,8})(\.(\d{0,8}))?$`)

func parseSeek(seek string) (out string, ok bool) {

	seek = strings.TrimSpace(seek)
	if seek == `` {
		ok = true
		return
	}

	list := regexpSeek.FindStringSubmatch(seek)
	if len(list) == 0 {
		return
	}

	h, _ := strconv.Atoi(list[4])
	m, _ := strconv.Atoi(list[6])
	s, _ := strconv.Atoi(list[7])
	ms := (list[9] + `000`)[0:3]
	if ms == `000` {
		ms = ``
	} else {
		ms = `.` + ms
	}

	if s >= 60 {
		m += s / 60
		s %= 60
	}
	if m >= 60 {
		h += m / 60
		m %= 60
	}

	out = fmt.Sprintf(`%s%02d:%02d:%02d%s`, list[1], h, m, s, ms)
	if out == `00:00:00` || out == `-00:00:00` {
		out = ``
	}

	ok = true
	return
}

func parseSeekTest() {

	list := []string{
		`-00:00:12.123`,
		`0:00:112.123`,
		`00:112`,
		`112`,
		``,
		`0`,
		`0.0`,
		`-00:112.34567`,
		`-99:0.3`,
	}
	for _, v := range list {
		fmt.Println()
		s, ok := parseSeek(v)
		b := `X`
		if ok {
			b = `O`
		}
		zj.F(`%s "%s" "%s"`, b, s, v)
	}
}
