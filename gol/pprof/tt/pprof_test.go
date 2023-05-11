package tt

import (
	"context"
	"fmt"
	"runtime/pprof"
	"testing"
	"time"
)

const url = "Hello World!"

func TestAdd(t *testing.T) {
	s := Add(context.Background(), "url")
	if s == "" {
		t.Errorf("Test.Add error!")
	}
}

func BenchmarkAdd(b *testing.B) {
	ctx := context.Background()
	for i := 0; i < b.N; i++ {
		labels := pprof.Labels("workerxxxxxxxxxxxx", "purge", "number", fmt.Sprint(i%10))
		pprof.Do(ctx, labels, func(ctx context.Context) {
			Add(ctx, url)
		})
	}
}

func Add(ctx context.Context, url string) string {

	for i := 0; i < 1000; i++ {

		ss := make([]byte, 1024*100)
		for j := 0; j < len(ss); j++ {
			ss[j] = 1
		}
		ss2 := make([]byte, 1024*100)
		copy(ss2[0:], ss[:])
		fmt.Printf("%v", ss2[100])
		time.Sleep(100 * time.Microsecond)
	}

	return ""
}

// go test -bench=. -cpuprofile=cpu.prof

// go tool pprof -http=:8080 cpu.prof
