package struct1

import (
	"fmt"
	"testing"
)

type A struct {
	*B
	*C
}

// func (p *A) Do() {
// 	fmt.Println("a do")
// }

type B struct {
}

func (p *B) Do() {
	fmt.Println("b do")
}

type C struct {
}

func (p *C) Do() {
	fmt.Println("c do")
}

func TestStuct1_ma3(t *testing.T) {
	a := &A{
		&B{},
		&C{},
	}

	a.B.Do()
	a.C.Do()
	// a.Do() err: ambiguous selector
}
