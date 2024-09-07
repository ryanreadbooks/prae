package pkg

import (
	"testing"
)

func TestExecCommand(t *testing.T) {
	err := ExecCommand("go", "ae")
	t.Log(err)
}
