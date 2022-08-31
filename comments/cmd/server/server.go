package main

import (
	api "finalproj/comments/pkg/api"
	storage "finalproj/comments/pkg/comdb"
	"log"
	"net/http"
)

// сервер newsgo
type server struct {
	db  storage.Interface
	api *api.API
}

func main() {
	// объект сервера
	var srv server

	// объект базы данных postgresql
	db, err := storage.New("postgres://postgres:postgres@127.0.0.1/newscomm")
	if err != nil {
		log.Fatal(err)
	}

	// Инициализируем хранилище сервера конкретной БД.
	srv.db = db

	// Создаём объект API и регистрируем обработчики.
	srv.api = api.New(srv.db)

	// запуск веб-сервера с API и приложением
	err = http.ListenAndServe(":81", srv.api.Router())
	if err != nil {
		log.Fatal(err)
	}

}

// TODO: введение айди и текста комментария и добавление этого в базу
// я хер знает как
