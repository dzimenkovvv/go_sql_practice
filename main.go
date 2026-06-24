package main

import (
	"Study/http_server"
	"fmt"
)

func main() {
	fmt.Println("Запуск HTTP сервера")
	fmt.Println("Привет!")

	err := http_server.StartHttpServer()
	if err != nil {
		fmt.Println("Во время работы сервера произошла ошибка:", err)
	} else {
		fmt.Println("Сервер запустился успешно!")
	}
}
