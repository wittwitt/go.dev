package slice1

import (
	"fmt"
	"testing"
)

func TestSlice1F1(t *testing.T) {

	/*
		1.slice {指针(ptr),长度(len),容量(cap)}
	*/

	var grid [][]string

	var line_teplate []string = []string{"cat", "black"}

	for i := 0; i < 5; i++ {
		grid = append(grid, line_teplate)
		// b = append(b, make([]string, 2))
	}

	// b[2][1] = "b"
	grid[2] = []string{"dog", "white"}

	for i := 0; i < len(grid); i++ {

		t.Logf("%p, %v", &grid[i], grid[i])
	}
}

func TestSlice1F2(t *testing.T) {
	s1 := []string{"a", "b", "c"}

	s2 := s1[0:1]

	for i := 0; i < 12; i++ {
		s1 = append(s1, fmt.Sprintf("%d", i))
		t.Logf("%p,  %p, %p, %v", &s1, &s1[0], &s2, s2)
	}
}

func TestSlice1_append(t *testing.T) {

	// 如果原 数组长度够， 则copy进
	// 不够，新建数组，

	// s1 := make([]string, 8)
	// s1[0] = "a"
	// s1[1] = "b"
	s1 := []string{"a", "b"}
	s1 = append(s1, "c1")
	s1 = append(s1, "c2")
	s1 = append(s1, "c3")
	t.Logf("%p, %p", &s1[0], &s1[1])

	s2 := []string{"c", "d"}
	t.Logf("%p, %p", &s2[0], &s2[1])

	s3 := append(s1, s2...)
	// for i, v := range s3 {
	// 	t.Log(i, v)
	// }
	t.Logf("%p, %p, %p, %p", &s3[0], &s3[1], &s3[2], &s3[3])

	t.Logf("%p, %p", &s1[0], &s1[1])
	t.Logf("%p, %p", &s2[0], &s2[1])
}

func TestSlice1_append2(t *testing.T) {

	s1 := []int{1, 2, 3, 4}

	s2 := append(s1[:2], s1[3:]...)

	for i, item := range s1 {
		println(item, &s1[i])
	}

	for i, item := range s2 {
		println(item, &s2[i])
	}
}

func TestSlice1_append3(t *testing.T) {

	s1 := []int{1, 2, 3, 4}

	s2 := append(s1[:2], s1[3:]...)

	s2 = append(s1[:2], []int{9, 10, 11, 12}...)

	for i, item := range s1 {
		println(item, &s1[i])
	}

	for i, item := range s2 {
		println(item, &s2[i])
	}
}

func TestSlice1_append4(t *testing.T) {

	s1 := []int{1, 2, 3, 4}

	for i, item := range s1 {
		println(item, &s1[i])
	}

	s2 := append(s1[:2], s1[3:]...)

	fmt.Printf("%d", len(s2))

	for i, item := range s1 {
		println(item, &s1[i])
	}
}

func TestSlice1_new(t *testing.T) {
	s := []string{1: "a", 4: "b"}
	fmt.Println(s)
}

func TestSlice1_new2(t *testing.T) {

}

type SliceCat struct {
	Name string
	Age  int
}

func TestSlice1_nil(t *testing.T) {
	ss := make([]SliceCat, 2)
	fmt.Printf("xxx %v\n", ss[0].Name)
}
