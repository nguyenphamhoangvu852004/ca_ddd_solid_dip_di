package main

import (
	"fmt"
	"sync"
	"time"
)

var count int
var mt sync.Mutex

func main() {

	go func() {
		mt.Lock()
		count++
		mt.Unlock()
	}()

	go func() {
		mt.Lock()
		count++
		mt.Unlock()
	}()

	time.Sleep(time.Duration(2) * time.Second)

	fmt.Println(count)
}
