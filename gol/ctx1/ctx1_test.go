package ctx1

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func Test_t1(t *testing.T) {

	pCtx1, pCancel1 := context.WithCancel(context.Background())

	ctx, cancel := context.WithTimeout(pCtx1, 3*time.Second)
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("end")
		cancel()
	}()

	time.Sleep(2 * time.Second)
	pCancel1()

	select {
	case <-ctx.Done():
		// 超时或者父 Context 被取消
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("超时")
		} else {
			fmt.Println("父 Context 被取消")
		}
	case <-time.After(8 * time.Second):
		// 执行需要在 8 秒内完成的操作
		fmt.Println("操作完成")
	}
}
