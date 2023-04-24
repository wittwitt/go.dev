// Package err1 err1_test

package err1

import (
	"fmt"
	"testing"

	"golang.org/x/xerrors"
)

// type CTime struct {
// 	// time.Time
// 	Na
// }

// type Na struct {
// }

// func (p Na) Error() string {
// 	return "33333"
// }

// func (p Na) GoString() string {
// 	return "xxxxx"
// }

type SubErr struct {
	str string
}

func (p SubErr) Error() string {
	return p.str
}

type MyErr struct {
	str    string
	SubErr *SubErr
}

// reflect.ValueOf  ，反射对MyErr的影响
// fmt.Fprintf ，// 管道输入，提前return了

// func (p *MyErr) Error() string {    // ok
func (p MyErr) Error() string {
	return p.SubErr.Error()
}

func Test_t2(t *testing.T) {

	e1 := MyErr{str: "mzw"}

	{
		xe2 := xerrors.Errorf("========: %v", e1) // no panic, and no print ========
		fmt.Printf("%v\n", xe2)
	}

	{
		xe2 := fmt.Errorf("========: %v", e1)
		fmt.Printf("%v\n", xe2)
	}
}
