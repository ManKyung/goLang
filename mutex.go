// example.go
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var data = []int{} // int 형 슬라이드 생성
	var mutex = new(sync.Mutex)

	go func() {
		for i := 0; i < 1000; i++ {
			mutex.Lock()
			data = append(data, i)
			mutex.Unlock()
			runtime.Gosched() // 다른 고루틴이 CPU를 사용할 수 있도록 양보
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			mutex.Lock()
			data = append(data, i)
			mutex.Unlock()
			runtime.Gosched()
		}
	}()

	time.Sleep(2 * time.Second)
	fmt.Println(len(data))
}
