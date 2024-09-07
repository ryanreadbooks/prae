package pkg

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func EnsureDir(s string) error {
	err := os.MkdirAll(s, 0755)
	return err
}

func ExecCommand(name string, cmds ...string) error {
	cmd := exec.Command(name, cmds...)
	extractOutput(cmd)
	err := cmd.Run()
	if err != nil {
		if errbuf, ok := cmd.Stderr.(*bytes.Buffer); ok {
			return fmt.Errorf("%w: %s", err, errbuf.String())
		}
	}

	return nil
}
