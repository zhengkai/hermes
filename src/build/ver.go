package build

import (
	"fmt"
	"strings"
)

// Ver ...
func Ver() {

	var t string
	if len(BuildTime) > 10 {
		t = fmt.Sprintf(`build at %s `, BuildTime[0:10])
	}

	fmt.Println(BuildGit)
	fmt.Printf(
		"\n%sby %s\n",
		t,
		strings.TrimPrefix(BuildGoVersion, `go version `),
	)
}
