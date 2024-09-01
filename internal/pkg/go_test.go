package pkg

import (
	"runtime"
	"testing"
)

func TestCollectGoVersion(t *testing.T) {
	t.Log(runtime.Version())
	t.Log(CollectGoVersion())
}
