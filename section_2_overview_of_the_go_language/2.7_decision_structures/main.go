package main

import "log"

/*
	В этом уроке изучаются конструкции ветвления
	if-else и switch. Синтаксис switch очень приятный,
	а в остальном все стандартно.
*/
func main() {
	var isTrue bool
	isTrue = true

	if isTrue {
		log.Println("true")
	} else {
		log.Println("false")
	}

	cat := "cat"

	if cat == "cat" {
		log.Println("cat")
	} else {
		log.Println("not a cat")
	}

	myInt := 100

	switch myInt {
	case 99:
		log.Println("99")
	case 100:
		log.Println("100")
	case 101:
		log.Println("101")
	default:
		log.Println("else")
	}
}
