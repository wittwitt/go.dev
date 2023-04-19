package any1

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_any1(t *testing.T) {

	type Cat struct {
		Name string
		Age  float64 // int
	}

	type Res1 struct {
		ID     int64           `json:"id"`
		Result json.RawMessage `json:"result,omitempty"`
	}

	type Res2 struct {
		ID     int64 `json:"id"`
		Result any   `json:"result,omitempty"`
	}

	type tData struct {
		data          []byte
		r1ResultIsNil bool
		r2ResultIsNil bool
	}

	testDatas := []tData{
		{[]byte(`{"id":100}`), true, true},
		{[]byte(`{"id":100,"result":null}`), false, true}, // json.Rawmsess, any ,不一样
		{[]byte(`{"id":100,"result":12345}`), false, false},
		{[]byte(`{"id":100,"result": {"name":"zs","age":99}}`), false, false},
	}

	for i := 0; i < len(testDatas); i++ {
		rs1 := &Res1{}
		require.NoError(t, json.Unmarshal(testDatas[i].data, rs1))
		require.Equal(t, rs1.Result == nil, testDatas[i].r1ResultIsNil, i)

		rs2 := &Res2{}
		require.NoError(t, json.Unmarshal(testDatas[i].data, rs2))
		require.Equal(t, rs2.Result == nil, testDatas[i].r2ResultIsNil, i)

		if rs2.Result != nil {
			switch val := rs2.Result.(type) {
			case float64:
				fmt.Println("float64", val)
			case *Cat:
				fmt.Println("cat", val.Age, val.Name)
			case map[string]string:
				fmt.Println("map", val)
			case map[string]json.RawMessage:
				fmt.Println("map[string]json", val)
			case map[string]any:
				fmt.Println("map[string]any", val)
			default:
				fmt.Println("no case type", reflect.TypeOf(rs2.Result))
			}
		}
	}
}
