package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Функция возвращает канал
	c := boring("boring!")
	for i := 0; i < 5; i++ {
		// Читаем из канала
		fmt.Printf("You say %q\n", <-c)
	}
	fmt.Println("You're boring - I'm living")
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
