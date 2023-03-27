package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {
		for {
			select { // default
			case nsg1 := <-c1:
				fmt.Println(nsg1)
			case nsg2 := <-c2:
				fmt.Println(nsg2)
			case <-time.After(time.Second):
				fmt.Println("timeout")
			}
		}
	}()

	fmt.Scanln()
}
