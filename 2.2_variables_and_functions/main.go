package main

import "log"

func main() {
	var something string
	var somethingElse string
	var world string
	var i int

	something, _ = saySomething("Hello")
	somethingElse, world = saySomething("Goodbye")

	log.Println(something)
	log.Println(somethingElse)
	log.Println(world)
	log.Println(i)
}

func saySomething(something string) (string, string) {
	return something, "World"
}
