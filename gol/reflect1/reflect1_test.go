package reflect1

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_t12(t *testing.T) {
	str := "str1"
	require.Equal(t, "string", anyFn(str).Name())
	require.Equal(t, "string", interfaceFn(str).Name())

	require.Equal(t, "", anyFn(&str).Name())
	require.Equal(t, "", interfaceFn(&str).Name())
}

func anyFn(r any) reflect.Type {
	return reflect.TypeOf(r)
}

func interfaceFn(r interface{}) reflect.Type {
	return reflect.TypeOf(r)
}

//

func Test_f1(t *testing.T) {
	var num float64 = 1.2345

	pointer := reflect.ValueOf(&num)
	value := reflect.ValueOf(num)

	// 可以理解为“强制转换”，但是需要注意的时候，转换的时候，如果转换的类型不完全符合，则直接panic
	// Golang 对类型要求非常严格，类型一定要完全符合
	// 如下两个，一个是*float64，一个是float64，如果弄混，则会panic
	convertPointer := pointer.Interface().(*float64)
	convertValue := value.Interface().(float64)

	fmt.Println(convertPointer)
	fmt.Println(convertValue)
}

type Cat struct {
	Name string
	Say  func(string) error
}

func (p *Cat) Hi(name string) error {
	return nil
}

func (p *Cat) String() string {
	return "cat"
}

func Test_f2(t *testing.T) {
	cat := &Cat{Name: "mzw"}

	typ := reflect.TypeOf(cat)
	{
		require.Equal(t, "*reflect1.Cat", typ.String())
		require.Equal(t, "", typ.Name())
		require.Equal(t, "ptr", typ.Kind().String())
		require.Equal(t, "reflect1.Cat", typ.Elem().String())

		// t.Log(typ.NumField()) // only struct
		require.Equal(t, 2, typ.NumMethod())
	}

	val := reflect.ValueOf(cat)
	{
		require.Equal(t, "<*reflect1.Cat Value>", val.String())
		require.Equal(t, "ptr", val.Kind().String())
		require.Equal(t, "<reflect1.Cat Value>", val.Elem().String())

		// t.Log(typ.NumField()) // only struct
		require.Equal(t, 2, val.NumMethod())
	}
}
