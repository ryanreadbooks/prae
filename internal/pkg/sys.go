package pkg

import "os"

func EnsureDir(s string) error {
	err := os.MkdirAll(s, 0755)
	return err
}
