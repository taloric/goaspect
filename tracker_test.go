package goaspect

import (
	"fmt"
	"testing"
)

func TestAspect_LogInfo(t *testing.T) {
	DefaultAspect.LogInfo("1").LogDebug("2").RetryOnce().Watch().Execute(func() {
		fmt.Println("Hello go")
	})
}
