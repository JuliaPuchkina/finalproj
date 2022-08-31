package main

import (
	api "finalproj/censoring/pkg/api"
	"log"
	"net/http"
)

// сервер newsgo
type server struct {
	api *api.API
}

func main() {
	// объект сервера
	var srv server

	// Создаём объект API и регистрируем обработчики.
	srv.api = api.New()

	// запуск веб-сервера с API и приложением
	err := http.ListenAndServe(":82", srv.api.Router())
	if err != nil {
		log.Fatal(err)
	}

}

// кажется, тут что-то недоделано...
