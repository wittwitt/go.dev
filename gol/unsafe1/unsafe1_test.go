package unsafe1

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

type Da struct {
}

func Test_t1(t *testing.T) {

	var a *Da = &Da{}

	pa := unsafe.Pointer(a)

	fmt.Println(reflect.ValueOf(a))
	fmt.Println(pa)

}
