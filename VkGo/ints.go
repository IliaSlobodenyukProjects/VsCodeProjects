package main

import "fmt"

type MyInt int

func (m MyInt) showYourSelf() { // копия, не меняет объект
	fmt.Println("m", m)
}

func (m *MyInt) add(i MyInt) { // меняет объект
	*m = *m + MyInt(i)
}
