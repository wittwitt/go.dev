package iota1

import (
	"fmt"
	"testing"
)

func Test_t1(t *testing.T) {
	type SectorFileType int
	const (
		a SectorFileType = 1 << iota
		b
		c
		d
		e

		f = iota
	)

	// t.Log(a)
	// t.Log(b)
	// t.Log(f)

	var existing SectorFileType
	var allocate SectorFileType

	//cccc
	existing = c
	allocate = c

	if existing == allocate {
		fmt.Println("ok2")
	}

	if existing^allocate == 0 {
		fmt.Println("ok3")
	}

	// 1110011
	// 1110011
	//
	//
	// 1001111
	//

	if existing|allocate != existing^allocate {
		fmt.Println("ok")
	}

}

func Test_t2(t *testing.T) {

	type SectorFileType int

	const (
		a SectorFileType = iota
		b
		c
		d
		e

		f = iota
	)

	// t.Log(a)
	// t.Log(b)
	// t.Log(f)
	// 用法： 检测一个类型是不是规定的类型

}

func Test_t3(t *testing.T) {

	type SectorFileType int

	const (
		a SectorFileType = 3 + iota
		b
		c
		d
		e

		f = iota
	)

	t.Log(a)
	t.Log(b)
	t.Log(f)
}

func Test_t4(t *testing.T) {
	t.Log(1 << 0)
	t.Log(1 << 1)
	t.Log(1 << 2)
}
