// example.go
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

type Data struct {
	tag    string
	buffer []int // 데이터 저장용 슬라이스
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용

	pool := sync.Pool{ // 풀 할당
		New: func() interface{} { // GET함수를 사용했을 때 호출될 함수 정의
			data := new(Data)             // 새 메모리 할당
			data.tag = "new"              // 태그 설정
			data.buffer = make([]int, 10) // 슬라이스 공간 할당
			return data                   // 할당한 메모리(객체) 리턴
		},
	}

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get().(*Data)

			for index := range data.buffer {
				data.buffer[index] = rand.Intn(100) // 슬라이스에 랜덤 값 저장
			}

			fmt.Println(data)
			data.tag = "used"
			pool.Put(data)
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get().(*Data)

			n := 0
			for index := range data.buffer {
				data.buffer[index] = n
				n += 2
			}
			fmt.Println(data)
			data.tag = "used"
			pool.Put(data)
		}()
	}

	fmt.Scanln()
}
