package package1

import (

	// 1. 文件夹名和package 不一样，必须使用别名倒入
	// some "gol/demo/pk/pk1"
	//
	// 2. 文件夹不能有空格
	//
	// pkss "gol/demo/pk/'pk spacke'" // 不可以
	//
	// 3. package pk-x  不可以，如果文件夹是 pk-x，可以 package pkxxx
	// 但，要别名引入 pkxxxx "gol/demo/pk/pk-x"
	//
	// 4. 可以 "gol/demo/pk/pk_f"

	"fmt"
	"testing"
)

func Test_t1(t *testing.T) {

	fmt.Println("ok")
	// some.Cat()
	// pk_f.Cat()
	// pkxxxx.Cat()
}

// 1. 文件夹不能空格
