package main

import "log"

type myStruct struct {
	FirstName string
}

/*
	При помощи указателя * на тип [myStruct] можно расширять
	функционал типа путем добавления функций. Что-то вроде методов
	внутри класса.
*/

func (m *myStruct) printFirstName() string {
	return m.FirstName
}

func main() {
	var myVar myStruct
	myVar.FirstName = "John"

	myVar2 := myStruct{
		FirstName: "Mary",
	}

	log.Println("myVar is set to", myVar.FirstName)
	log.Println("myVar2 is set to", myVar2.printFirstName())
}
