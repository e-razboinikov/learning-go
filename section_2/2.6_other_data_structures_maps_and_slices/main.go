package main

import (
	"log"
	"sort"
)

func main() {
	myMap := make(map[string]string)
	// Таким образом создается [Map] с динамическим типом value.
	// Этот способ не рекомендуется использовать
	// mySecondMap := make(map[string]interface{})

	myMap["dog"] = "Samson"

	var mySlice []int

	mySlice = append(mySlice, 4)
	mySlice = append(mySlice, 2)

	sort.Ints(mySlice)

	log.Println(mySlice)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	log.Println(numbers)

	log.Println(numbers[0:2])

	log.Println(numbers[6:9])
}
