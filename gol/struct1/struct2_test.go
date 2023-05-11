package struct1

import (
	"fmt"
	"testing"
)

// 字段 方法/指针方法，，全不能冲突，唯一

// type A struct {
// 	Some func()
// }

// func (p A) Some() {
// }

// func (p *A) Some() {
// }

type Cat struct {
	Name string
}

func (p Cat) Say() Cat { // new Cat copy from p
	p.Name = "saysay"
	return p
}

func (p *Cat) Hi() {
	p.Name = "HiHi"
}

func Test_t1(t *testing.T) {
	cat := Cat{Name: "mzw"}
	cat.Say()
	fmt.Println(cat.Name)

	cat.Hi()
	fmt.Println(cat.Name)
}

func Test_fnt2(t *testing.T) {
	c1 := &Cat{Name: "c1"}
	c2 := *c1
	c2.Name = "c2"

	t.Log(c1.Name != c2.Name)
}
