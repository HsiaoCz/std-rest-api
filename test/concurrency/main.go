package main

import (
	"fmt"
	"time"
)

// some struct we use to make an channel
type foo struct {
	val string
}

// concurrency
// fetchResources
func fetchResources(n int) string {
	time.Sleep(time.Second * 2)
	return fmt.Sprintf("result %d", n)
}

func main() {
	// result := fetchResources()
	// fmt.Println("the result :", result)
	// fetchResources()
	// fetchResources()
	// fetchResources()
	// go fetchResources(1) // async

	// this like go fetchResources
	go func() {
		result := fetchResources(1)
		fmt.Println(result)
	}()

	resultch := make(chan string)
	// put the foo in the channel
	resultch <- "foo"
	// read from the channel
	resultcc := <-resultch
	fmt.Println(resultcc)

	// 1. unbuffered channel
	// unbuffered channel always locked
	// till the recevier receive the value
	// 2. buffered channel
	fooch := make(chan foo)
	fooch <- foo{
		val: resultcc,
	}
	fmt.Println(fooch)
}
