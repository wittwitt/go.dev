package reflect1

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_t2(t *testing.T) {
	fmt.Println(call1(foo, "hello"))
	fmt.Println(call2(foo, "hello"))
}

func foo(x string) string {
	return x + "!"
}

func call1(f, x interface{}) interface{} {
	fValue := reflect.ValueOf(f)
	inValue := reflect.New(reflect.TypeOf(x)).Elem()
	inValue.Set(reflect.ValueOf(x))

	inValue.Set(fValue.Call([]reflect.Value{inValue})[0])

	return inValue.Interface()
}

func call2(f, x interface{}) interface{} {
	fValue := reflect.ValueOf(f)
	xCopy := x
	inValue := reflect.ValueOf(&xCopy).Elem()

	inValue.Set(fValue.Call([]reflect.Value{inValue})[0])

	return inValue.Interface()
}
