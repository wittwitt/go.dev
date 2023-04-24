package ifelse

import (
	"fmt"
	"testing"
)

func TestLabel13(t *testing.T) {
	i := 10
RS:
	{
		i = i - 1
		fmt.Println("ok==")
		if i > 5 {
			// break RS ，，break label 必须是对应for/select /switch
			goto RS
		}
		fmt.Println("ok== ==")
	}
	fmt.Println(i)
	if i > 2 {
		goto RS
	}
}

func TestLabel12(t *testing.T) {
	i := 1
FOR:
	for {
		i++
		fmt.Println(i)
		if i > 100 {
			break FOR // 单层for，直接break就可以了，多层嵌套，可以用Label
		}
	}
	fmt.Println("ccc")
}

func TestLabel1(t *testing.T) { // label: for/swith/select ,可以goto，break

	var a int = 10
	var j int = 1

Loop:
	for a < 20 {
		for j < a {
			if j > 10 {
				continue // 只能跳出一层
			}

			if j < 10 {
				continue Loop // continue Loop // 可以continue  loop对应for/swit/select
			}
		}
	}
}
