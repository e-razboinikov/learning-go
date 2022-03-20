package main

import (
	"log"
	"time"
)

// Если название начинается с большой буквы, то эта переменная (функция и тд)
// видна за пределами пакета. Если с маленькой, то является приватной и видна только в пределах пакета.
var S = "seven"

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	// Ещё один способ объяления переменной
	user := User{
		FirstName: "Trevor",
		LastName:  "Sawler",
	}

	log.Println(user.BirthDate)
}
