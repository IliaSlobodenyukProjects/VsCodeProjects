package main

import (
	"fmt"
	"math/rand"
	"time"
)

// использование нескольких каналов для их объединения

func main() {
	c := funIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're both boring - I'm living")
}

// получаем 2 канала только для чтения -> возвращаем 1 только для чтения
func funIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

func funInPro(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case nsg := <-input1:
				c <- nsg
			case nsg := <-input2:
				c <- nsg
			}
		}
	}()
	return c
}

func boring(msg string) <-chan string { // возвращает канал строк только для чтения
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
	}()
	return c
}
