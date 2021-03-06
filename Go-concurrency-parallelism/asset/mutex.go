package main

import (
	"sync"
	"fmt"
)

var id int

func generateId(mutex *sync.Mutex) int {
    // Lock()/Unlock()をペアで呼び出してロックする
    mutex.Lock()
    id++
	mutex.Unlock()
	return id
}

func main() {
	// sync.Mutex構造体の変数宣言
	// 次の宣言をしてもポインタ型になるだけで正常に動作します
	// mutex := new(sync.Mutex)
	var mutex sync.Mutex

	for i := 0; i < 100; i++ {
	    go func() {
	        fmt.Printf("id: %d\n", generateId(&mutex))
	    }()
	}
}