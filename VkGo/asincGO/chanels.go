package main

import "fmt"

var c chan int

// Каналы нужны для синхронизации данных, когда есть n горутин

func main() {
	//Создаем канал
	c := make(chan string)
	// стартуем пишущую горутину
	go greet(c)
	for i := 0; i < 5; i++ {
		fmt.Println(<-c, " , ", <-c)
	}

	stuff := make(chan int, 7) // создание буферезированного канала
	for i := 0; i < 19; i = i + 3 {
		stuff <- i
	}
	close(stuff)
	fmt.Println("Res", process(stuff))
}

func greet(c chan<- string) {
	//Запускаем бесконечный цикл
	for {
		// пишем в канал пару строк
		// подпрограмма будет заблокирована до того, как кто-то захочет прочитать
		// значения в канале, в этом месте получается синхронизация
		c <- fmt.Sprintf("Lord")
		c <- fmt.Sprintf("Stormtrooper")
	}
}

func process(input <-chan int) (res int) {
	for r := range input {
		res += r
	}
	return
}
