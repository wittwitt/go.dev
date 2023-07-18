package reflect1

import (
	"fmt"
	"reflect"
	"testing"
)

type A struct {
	name string
}

func (p *A) Some() {

}

func Test_t3332(t *testing.T) {
	a := A{name: "aa"}
	a2 := &A{name: "a2a2"}

	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(a2))

	fmt.Println(reflect.ValueOf(a))
	fmt.Println(reflect.ValueOf(a2))

	fmt.Println(reflect.TypeOf(a).Elem())
}
