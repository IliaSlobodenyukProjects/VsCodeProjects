package main

import "fmt"

type example struct { // возвращается в других пакетах только функцией, Example - можно создать с помощью объявления
	flag    bool  //   только в этом пакете
	Counter int16 // доступен в других пакетах
	pi      float32
}

func main() {
	var e1 example
	fmt.Printf("%+v\n", e1)

	e2 := example{
		flag:    true,
		Counter: 10,
		pi:      3.14,
	}
	fmt.Println("flag", e2.flag)
	fmt.Println("Counter", e2.Counter)
	fmt.Println("pi", e2.pi)
}
