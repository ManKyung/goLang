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

	var data int = 0
	var rwMutex = new(sync.RWMutex) // 읽기, 쓰기 뮤텍스 생성

	go func() {
		for i := 0; i < 3; i++ {
			rwMutex.Lock()
			data += 1
			fmt.Println("write : ", data)
			time.Sleep(10 * time.Millisecond)
			rwMutex.Unlock()
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			rwMutex.RLock()
			fmt.Println("read : ", data)
			time.Sleep(1 * time.Second)
			rwMutex.RUnlock()
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			rwMutex.RLock()
			fmt.Println("read2 : ", data)
			time.Sleep(2 * time.Second)
			rwMutex.RUnlock()
		}
	}()
	

	time.Sleep(10 * time.Second)
}
