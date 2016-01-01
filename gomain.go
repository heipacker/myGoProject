package main

import (
	"fmt"
	"os"
)
import (
	log "github.com/cihub/seelog"
)
import "myGoProject/stringutil"
import "strconv"

var a int = 1
var b string = "bstring"

const cConst = 10

//常量例子
const (
	ca = 1 << iota
	cb = 1 << iota
	cc
)

type human struct {
	Sex int
}

type teacher struct {
	human
	Name string
	Age  int
}

type student struct {
	human
	Name string
	Age  int
}

func main() {
	fmt.Println("hello world")
	fmt.Println(a)
	fmt.Println(stringutil.Reverse(b))
	var a int = 65
	ab := string(a)
	fmt.Println(ab)
	fmt.Println(strconv.Itoa(a))
	bb, _ := strconv.Atoi("65")

	fmt.Println(bb)
	fmt.Println(cConst)
	fmt.Println(ca)
	fmt.Println(cb)
	fmt.Println(cc)

	if a := 1; a > 0 {
		fmt.Println(a)
	}
	testArray()
	testMap()
	testStruct()
	testGetwd()
	testJson()
	testLog()
}

func testGetwd() {
	wd, err := os.Getwd()
	if err != nil {
		log.Criticalf("Getwd error")
	}
	fmt.Println(wd)
}

func testArray() {
	a := [2]int{1, 2}

	fmt.Println(a)

	b := [2][3]int{
		{1, 2, 3}, {3, 4, 5},
	}
	fmt.Println(b)
	c := []int{1, 2, 4}
	fmt.Println(c)
}

func testMap() {
	var m map[int]string
	m = make(map[int]string)
	var mm map[int]string = make(map[int]string)
	mmm := make(map[int]map[int]string)
	m[1] = "aa"
	m[2] = "bb"
	delete(m, 1)
	fmt.Println(m)
	mmm[1] = make(map[int]string)
	mmm[1][1] = "ok!"

	fmt.Println(mm)
	fmt.Println(mmm)
}

func testStruct() {
	a := teacher{Name: "joe", Age: 19, human: human{Sex: 0}}
	b := student{Name: "joe1", Age: 20, human: human{Sex: 1}}
	fmt.Println(a)
	fmt.Println(b)
}
