package interface1

import (
	"fmt"
	"testing"
)

type Iface interface {
	F1() (int, error)
	F2() (int, error)
	F3() (int, error)
}

type Face struct {
}

func (p *Face) F1() (int, error) {
	return 1, nil
}

func (p *Face) F2() (int, error) {
	return 2, nil
}

func (p *Face) F3() (int, error) {
	return 3, nil
}

func IFn(name string, ff func(Iface) (int, error)) {
	f := &Face{}
	fmt.Println(ff(f))
}

func TestI1f1(t *testing.T) {
	IFn("", Iface.F1)
	IFn("", Iface.F2)
	IFn("", Iface.F3)

	//
	IFn("", SS)
}

func SS(i Iface) (int, error) {
	return i.F1()
	// return 100, nil
}
