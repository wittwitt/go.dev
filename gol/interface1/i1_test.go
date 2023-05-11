package interface1

import (
	"testing"
)

type Animal interface {
	Eat(f string) error
}

type Cat struct {
	Name string
}

// 指针接收器
func (p *Cat) Eat(f string) error {
	return nil
}

type Dog struct {
	Age int
}

// 值接收器
func (p Dog) Eat(f string) error {
	return nil
}

// 接口实现，接收器可以是指针，或者值
func Test_t1(t *testing.T) {
	var cat Animal = &Cat{}
	var dog Animal = &Dog{}
	cat.Eat("")
	dog.Eat("")
}
