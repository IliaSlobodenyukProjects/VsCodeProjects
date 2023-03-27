package main

import "fmt"

func main() {

}

func singlePanic() (err error) {
	fmt.Println("start test")
	defer catchPanic(err) // этот блок будет вызван сразу по выходе из функции, то есть после panic в том числе
	panic("at the disco")
	// r := recover() // при выполнении panic эта функция никогда не выполнится
}

func catchPanic(err error) {
	var ok bool
	//обработка паники
	if r := recover(); r != nil {
		if err, ok = r.(error); ok {
			fmt.Println()
		}
		fmt.Println("PANIC deffered")
	}
}
