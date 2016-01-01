package main

import (
	"math"
	"runtime"
	"sync"
	"time"
)

func testGo() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := new(sync.WaitGroup)
	wg.Add(2)
	start := time.Now().UnixNano()
	for i := 0; i < 2; i++ {
		go func(id int) {
			defer wg.Done()
			sum(id)
		}(i)
	}
	println("cost time:", time.Now().UnixNano()-start)
	wg.Wait()
}

func sum(id int) {
	var x int64
	for i := 0; i < math.MaxUint32; i++ {
		x += int64(i)
	}
	println(id, x)
}
