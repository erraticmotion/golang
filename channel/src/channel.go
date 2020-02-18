// go run channel.go 

package main

import (
	"fmt"
	"time"
)

func printCount(c chan int) {
	num := 0
	for num >= 0 {
		num = <-c
		fmt.Print(num, " ")
	}
}

func main() {
	c := make(chan int)
	go printCount(c)

	a := []int{8, 6, 7, 5, 3, 0, 9, 1}

	for _, v := range a {
		time.Sleep(time.Second)
		c <- v
	}

	time.Sleep(time.Second)
	fmt.Println("End of main")
}
