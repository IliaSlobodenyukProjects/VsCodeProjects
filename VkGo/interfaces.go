package main

import "fmt"

type Speaker interface {
	/*
		~ В именовании всегда постфикс -er (Sender, Reader, Closer, etc)
		~ Duck typing (неявная, латентная или утиная типизация)
		~ Интерфейс не может содержать значений, только методы
			В GO структура с методами будет удовлетворять интерфейсу просто самим фактом объявления метода
			Структура хранит данные, но не поведение - интерфейс хранит поведение, но не данные
	*/
	SayHello()
}

type Flyer interface {
	Fly()
}

type Bird struct {
	Name string
}

func (b Bird) Fly() {
	fmt.Println(b.Name + " is flying")
}

func DoFly(f Flyer) {
	f.Fly()
}

func main() {
	duck := Bird{"duck"}
	DoFly(duck)
}
