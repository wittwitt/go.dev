package ifelse

import (
	"fmt"
	"testing"
)

func return1Err() (int, error) {
	return 10, fmt.Errorf("xxxx")
}

func return1() (err error) {

	// i := 10

	i, err := return1Err()
	if err != nil {
		err = fmt.Errorf("some error: %v", err)
		return
	}

	fmt.Println(i)
	return nil
}

func Test_return1(t *testing.T) {

	t.Log(return1())

}
