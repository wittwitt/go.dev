package reflect1

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_t2(t *testing.T) {
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
