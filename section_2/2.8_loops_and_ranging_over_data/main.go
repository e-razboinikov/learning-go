package main

import "log"

type User struct {
	FirstName string
	LastName  string
}

func main() {
	// for i := 0; i <= 10; i++ {
	// 	log.Println(i * i)
	// }

	mySlice := []string{
		"dog", "cat", "fish", "banana",
	}

	for _, x := range mySlice {
		log.Println(x)
	}

	myMap := make(map[string]string)
	myMap["dog"] = "dog"
	myMap["cat"] = "cat"
	myMap["fish"] = "fish"

	for i, x := range myMap {
		log.Println(i, x)
	}

	var mySecondSlice []User
	userOne := User{
		FirstName: "John",
		LastName:  "Doe",
	}
	userTwo := User{
		FirstName: "Max",
	}
	mySecondSlice = append(mySecondSlice, userOne, userTwo)

	for i, x := range mySecondSlice {
		log.Println(i, x.FirstName)
	}
}
