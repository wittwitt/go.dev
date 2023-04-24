package json1

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_t1(t *testing.T) {
	type Cat struct {
		Name     string
		ByteData []byte // []byte 默认，，base64.StdEncoding.DecodeString
	}
	ppV := Cat{Name: "abc", ByteData: []byte("abc")}
	jsonData, err := json.Marshal(&ppV)
	require.NoError(t, err)
	require.Equal(t, `{"Name":"abc","ByteData":"YWJj"}`, string(jsonData))
}

func Test_stuct_defaut(t *testing.T) {
	t2 := struct {
		time.Time
		N int
	}{
		time.Date(2020, 12, 20, 0, 0, 0, 0, time.UTC),
		5,
	}
	t2Data, err := json.Marshal(t2) // 使用time.Time的Marshal ，丢失 int 5
	require.NoError(t, err)
	require.Equal(t, `"2020-12-20T00:00:00Z"`, string(t2Data))

	//
	t3 := Cat{
		5,
		time.Date(2020, 12, 20, 0, 0, 0, 0, time.UTC),
	}
	t3Data, err := json.Marshal(t3)
	require.NoError(t, err)
	require.Equal(t, `{"age":"100"}`, string(t3Data))

	//
	t4 := struct {
		*Cat
		*Dog
		N int
	}{
		&Cat{1, time.Date(2020, 12, 20, 0, 0, 0, 0, time.UTC)},
		&Dog{"wawa"},
		5,
	}
	t4Data, err := json.Marshal(t4)
	require.NoError(t, err)
	require.Equal(t, `"2020-12-20T00:00:00Z"`, string(t4Data))
}

type Cat struct {
	Age int
	time.Time
}

func (p Cat) MarshalJSON() ([]byte, error) {
	return []byte(`{"age":"100"}`), nil
}

type Dog struct {
	Name string
}

func (p *Dog) MarshalJSON() ([]byte, error) {
	return []byte("DogDogDog"), nil
}
