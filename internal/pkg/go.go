package pkg

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"runtime"
)

var gore = regexp.MustCompile(`go(\d+\.\d+\.\d+)`)

func ignoreOutput(cmd *exec.Cmd) {
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
}

func CollectGoVersion() string {
	cmd := exec.Command("go", "version")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	// 执行命令
	err := cmd.Run()
	if err != nil {
		return runtime.Version()
	}

	match := gore.FindStringSubmatch(out.String())
	if len(match) == 0 {
		return runtime.Version()
	}
	return match[1]
}

func GoModInit(module string) error {
	cmd := exec.Command("go", "mod", "init", module)
	ignoreOutput(cmd)
	return cmd.Run()
}

func GoModEditGo(gv string) error {
	cmd := exec.Command("go", "mod", "edit", fmt.Sprintf("-go=%s", gv))
	ignoreOutput(cmd)
	return cmd.Run()
}

func GoModTidy() error {
	cmd := exec.Command("go", "mod", "tidy")
	ignoreOutput(cmd)
	return cmd.Run()
}
