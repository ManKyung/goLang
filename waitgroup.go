// example.go
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용

	wg := new(sync.WaitGroup)

	for i := 0; i < 10; i++ {
		wg.Add(1)        // 반복할 때마다 wg.add함수로 1씩 추가
		go func(n int) { // 고루틴 10개 생성
			fmt.Println(n)
			wg.Done() // 고루틴 완료를 알려줌
		}(i)
	}

	wg.Wait() // 고루틴이 끝날 때 까지 대기
	fmt.Println("the end")
}
