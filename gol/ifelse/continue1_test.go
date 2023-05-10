package ifelse

import (
	"fmt"
	"testing"
)

func TestFor1(t *testing.T) {
	i := 100

	switch {
	case i > 10:
		fmt.Println(">10")
		// auto break
	case i > 50:
		fmt.Println(">50")
	case i > 101:
		fmt.Println(">101")
	case i > 200:
		fmt.Println(">200")
	}
}

func TestFor2(t *testing.T) {
	{
		i := 100
		switch {
		case i > 10:
			fmt.Println(">10")
			fallthrough
		case i > fn_TestFor2(150):
			fmt.Println(">150")
		case i > 101:
			fmt.Println(">101")
		default:
			fmt.Println(">200")
		}
	}

	{
		i := 2
		switch i {
		case 2:
			fmt.Println("2")
			fallthrough
		case 3:
			fmt.Println("3")
		case 4:
			fmt.Println("4")
		default:
			fmt.Println("not 2,3,4")
		}
	}
}

func fn_TestFor2(i int) int {
	println("fn_TestFor2", i)
	return i
}

func TestTcontinue(t *testing.T) {
	for i := 0; i < 10; i++ {
		k := i + 1
		switch k {
		case 1:
			fmt.Println("k", 1)
		case 2:
			fmt.Println("k", 2)
			continue
		default:
		}
		fmt.Println("===", i, k)
	}

}

func TestTbreak(t *testing.T) {

	i := 0
	for {
		i = i + 1

		switch i {

		case 2:
			i = 100
			fmt.Println("222")
			break // break switch not for
		default:
		}
		fmt.Println("===", i)
	}
}
