package main

import (
	"encoding/json"
	"log"
)

// В ковычках - JsonKey
type Person struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	HairsColor string `json:"hairs_color"`
	HasDog     bool   `json:"has_dog"`
}

func main() {
	// Моковый json
	myJson := `
	[
		{
			"first_name": "Clark",
			"last_name": "Kent",
			"hairs_color": "black",
			"has_dog": true
		},
		{
			"first_name": "Bruce",
			"last_name": "Wayne",
			"hairs_color": "black",
			"has_dog": false
		}
	]`

	// fromJson - Unmarshal
	var unmarshalled []Person
	// Парсинг json-a (принимает массив [byte]),
	// возвращает массив - резульат и, возможно, ошибки
	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	// Обработка ошибок
	if err != nil {
		log.Println("Error unmarshalled json", err)
	}
	// Вывод результата в консоль
	log.Printf("Unmarshalled: %v", unmarshalled)
	log.Println(unmarshalled[0].FirstName)

	//toJson - Marshal
	var marshalled []Person
	firstPerson := Person{
		FirstName:  "John",
		LastName:   "Doe",
		HairsColor: "white",
		HasDog:     false,
	}
	secondPerson := Person{
		FirstName:  "Bill",
		LastName:   "Ingrew",
		HairsColor: "ginger",
		HasDog:     true,
	}
	marshalled = append(marshalled, firstPerson, secondPerson)
	// toJson. Возвращает [[]byte] и, возможно, ошибки
	newJson, err := json.Marshal(marshalled)
	// Обработка ошибок
	if err != nil {
		log.Println("error marshalled: ", err)
	}
	// Вывод результата в консоль
	log.Println(string(newJson))
}
