package map1

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

// key按hash存储，遍历的顺序都可能不一样(golang故意的).
func TestMap1_key_order(t *testing.T) {
	m := map[string]string{
		"1": "1",
		"2": "2",
		"3": "3",
	}
	for k, v := range m {
		println(k, v)
	}
}

func TestMap1_del(t *testing.T) {
	m := make(map[int]string)
	delete(m, 1)      // 可以删除不存在的key
	fmt.Println(m[1]) // 可以获取不存在的key value （默认值）
}

// 可以并发读
func TestMap1_mut_thread_read(t *testing.T) {
	m := map[string]int{
		"a": 101,
		"b": 102,
	}

	startCh := make(chan struct{})
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(index int) {
			<-startCh
			for j := 0; j < 30; j++ {
				require.Equal(t, 101, m["a"])
				require.Equal(t, 102, m["b"])
			}
			wg.Done()
		}(i)
	}
	close(startCh)
	wg.Wait()
}

// 也不能： 写加锁，读不加锁！！！
func TestMap1_mut_thread_read2(t *testing.T) {
	m := map[string]int{
		"a": 1,
		"b": 1,
	}

	var lock sync.Mutex

	startCh := make(chan struct{})

	writeWg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		writeWg.Add(1)
		go func(index int) {
			<-startCh
			for j := 0; j < 1000; j++ {
				lock.Lock()
				index = index + 1
				m["a"] = index
				lock.Unlock()
			}
			writeWg.Done()
		}(i)
	}

	readWg := sync.WaitGroup{}
	for i := 0; i < 30; i++ {
		readWg.Add(1)
		go func(index int) {
			<-startCh
			for j := 0; j < 10000; j++ {
				if k, exist := m["a"]; exist {
					fmt.Println(index, k)
				}
			}
		}(i)
		readWg.Done()
	}

	close(startCh)
	writeWg.Wait()
	readWg.Wait()
}
