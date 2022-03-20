package main

import (
	"log"

	"github.com/e-razboinikov/learning-go/2.11_packages/structs"
)

/*
	Для создания модуля используется следующая команда:
	go mod init MODULE_NAME
	Пример: go mod init github.com/e-razboinikov/learning-go/2.11_packages
	В результате генерируется файл go.mod и можно создавать папки
*/

func main() {
	firstUser := structs.User{
		FirstName: "John",
		LastName:  "Doe",
	}

	log.Println(firstUser.FirstName, firstUser.LastName)
}
