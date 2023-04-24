package runtime1

import (
	"fmt"
	"runtime"
	"testing"
)

func TestCall(t *testing.T) {
	C1()
}

func C1() {
	C2()
}

func C2() {
	C3()
}

func C3() {
	if pc, _, _, ok := runtime.Caller(0); ok {
		pcName := runtime.FuncForPC(pc).Name()
		fmt.Println(pcName)
	}

	pc := make([]uintptr, 10)
	n := runtime.Callers(0, pc)
	for i := 0; i < n; i++ {
		fmt.Println(i, runtime.FuncForPC(pc[i]).Name())
	}
}
