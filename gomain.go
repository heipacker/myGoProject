package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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
	testGoLang()
	testArray()
	testMap()
	testStruct()
	testType()
	testBreakContinue()
	testArgument("1", 1, 12, 2, 3, 4)
	testGo()
	testSelect()
	testPath()
	testGetwd()
	testJson()
	testLog()
	dbQuery()
}

func testSelect() {
	a, b := make(chan int, 3), make(chan int)
	go func() {
		v, ok, s := 0, false, ""
		for {
			select {
			// 随机选择可用用 channel,接收数据。
			case v, ok = <-a:
				s = "a"
			case v, ok = <-b:
				s = "b"
			}
			if ok {
				fmt.Println(s, v)
			} else {
				os.Exit(0)
			}
		}
	}()

	for i := 0; i < 5; i++ {
		select {
		// 随机选择可用用 channel,发送数据。
		case a <- i:
		case b <- i:
		}
	}
	close(a)
	select {} // 没有可用用 channel,阻塞 main goroutine。
}

func arrayTest() {
	a, j := [4]int{1, 2, 3, 4}, 2
	for i := 0; i < len(a); i++ {
		log.Info(a[i])
	}
	log.Info(j)
}

func testPath() {
	if lpath, err := exec.LookPath(os.Args[0]); err == nil {
		if pathName, err := filepath.Abs(lpath); err == nil {
			println(pathName)
			return
		}
	}
	println("get error.")
}

func testArgument(s string, s1 int, a ...int) {
	println(s)
	println(s1)
	for i, length := 0, len(a); i < length; i++ {
		println(a[i])
	}
	for _, i := range a {
		println(i)
	}
}

func testBreakContinue() {
L2:
	for i := 0; i < 3; i++ {
	L1:
		for j := 0; j < 5; j++ {
			if j > 2 {
				break L1
			}
			if i > 1 {
				break L2
			}
			print(i, ":", j, "	")
		}
		println()
	}
}

func testType() {
	u := "电脑"
	println(u)

	us := []rune(u)
	us[1] = '话'

	println(string(us))
}

func testGoLang() {
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
