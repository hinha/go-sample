package main

import (
	"fmt"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

func waitMicrosecond() {
	time.Sleep(10000 * time.Microsecond)
	fmt.Println("print myself")
}

func say(s string) {
	defer waitGroup.Done()
	for i := 0; i < 10; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
	}

}

func main() {
	fmt.Println("Started...")

	// go waitMicrosecond()

	// go func() {
	// 	fmt.Println("An anonymous function!")
	// }()
	waitGroup.Add(1)
	go say("heloo")
	waitGroup.Add(1)
	go say("not hello")
	waitGroup.Wait()

}
