package main

import (
	"log"

	"github.com/e-razboinikov/learning-go/2.12_channels/helpers"
)

const numPool = 1000

func calculateValue(intChan chan int) {
	randomNubmer := helpers.RandomNumber(numPool)
	intChan <- randomNubmer
}

func main() {
	intChan := make(chan int)
	defer close(intChan)

	go calculateValue(intChan)

	num := <-intChan
	log.Println(num)
}
