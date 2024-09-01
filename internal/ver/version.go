package ver

import "strings"

var version string

func SetVersion(v string) {
	version = v
}

func Major() string {
	r := strings.Split(version, ".")
	return r[0]
}
