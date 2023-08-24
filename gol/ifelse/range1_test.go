package ifelse

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type range1S1 struct {
	name string
}

func Test_range(t *testing.T) {

	list := []range1S1{
		{"z1"},
		{"z2"},
		{"z3"},
	}

	var lock sync.Mutex

	lock.Lock()

	go func() {
		for _, item := range list {
			fmt.Println(item.name)
			fmt.Println("ok")
			lock.Unlock()
			lock.Lock()
		}
	}()

	lock.Lock()
	list[1].name = "haha"
	lock.Unlock()

	time.Sleep(10 * time.Second)

}

func TestRange1_range_slice(t *testing.T) {
	m := make([]int, 7)
	m[0] = 111
	m[1] = 100
	m[2] = 200
	m[3] = 300
	m[4] = 400
	m[5] = 500
	m[6] = 600

	/*
		m_len=len(m)
		for i:=0;i<m_len;i++{
			item=m[i]
			//
			...
		}
	*/

	for i, item := range m {

		fmt.Println(i, item)

		if i%2 == 0 {
			m = append(m[0:i], m[i+1:]...)
		}
	}
}

func TestRange1_range_map(t *testing.T) {
	m := make(map[string]int)
	m["1"] = 1
	m["2"] = 2
	m["3"] = 3
	m["4"] = 4
	m["5"] = 5
	m["6"] = 6

	/*
		map key hash存储，所以删除没有问题，
		但在for里添加新元素，不一定会别便利
	*/

	for k, v := range m {
		fmt.Println(k, v)
		if v%2 == 0 {
			delete(m, k)
		}
	}

	for k, v := range m {
		fmt.Println(k, v)
		if v == 2 {
			m["k"] = 10
		}
	}

}
