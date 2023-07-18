package reflect1

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

type ICat interface {
	Eat(string) string
}

func Test_abc(t *testing.T) {

	t.Log("abc")

	var icat ICat
	require.Equal(t, nil, reflect.TypeOf(icat))

	v := reflect.ValueOf(icat)

	t.Log(v.Kind() == reflect.Invalid)

	tof := reflect.TypeOf(&icat).Elem()

	//	t.Log(tof.Elem())

	t.Log(" tof.NumMethod()", tof.NumMethod())

	for i := 0; i < tof.NumMethod(); i++ {
		method := tof.Method(i)
		t.Log("Method:", method.Name)
		t.Log(method.Func.Addr().Interface())
	}
}
