// go run hello.go

package main

import (
	"fmt"
	"time"
	"sync"
	"sync/atomic"
)

var (
	shutdown int64
	wg sync.WaitGroup
)

func main() {
	wg.Add(2)

	go doWork("A")
	go doWork("B")

	time.Sleep(5 * time.Second)

	fmt.Println("shutdown now")
	atomic.StoreInt64(&shutdown, 1)

	wg.Wait()
}

func doWork(name string){
	defer wg.Done()

	for {
		fmt.Printf("doing %s work\n", name)
		time.Sleep(250 * time.Microsecond)

		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("shutting %s down\n", name)
			break
		}
	}
}