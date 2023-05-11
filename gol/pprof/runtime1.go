package main

import (
	"context"
	"runtime/pprof"
)

func do() {
	pprof.Do(context.TODO(), pprof.Labels("label-key", "label-value"), func(ctx context.Context) {
		// execute labeled code
	})
}
