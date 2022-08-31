package api

import (
	"encoding/json"
	storage "finalproj/comments/pkg/comdb"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// API приложения.
type API struct {
	r  *mux.Router       // маршрутизатор запросов
	db storage.Interface // база данных
}

// Конструктор API.
func New(db storage.Interface) *API {
	api := API{
		db: db,
	}
	api.r = mux.NewRouter()
	api.endpoints()
	return &api
}

// Router возвращает маршрутизатор запросов.
func (api *API) Router() *mux.Router {
	return api.r
}

// Регистрация методов API в маршрутизаторе запросов.
func (api *API) endpoints() {
	// получить комментарии
	api.r.HandleFunc("/comments/{n}", api.commentsHandler).Methods(http.MethodGet, http.MethodOptions)
	// добавление комментария в базу
	api.r.HandleFunc("/comments", api.addCommHandler).Methods(http.MethodPost, http.MethodOptions)
}

// commentsHandler возвращает комментарии по id новости
func (api *API) commentsHandler(w http.ResponseWriter, r *http.Request) {
	s := mux.Vars(r)["n"]
	n, _ := strconv.Atoi(s)
	comments, err := api.db.Comments(n)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	bytes, err := json.Marshal(comments)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
}

// addCommHandler добавляет комментарий в базу
func (api *API) addCommHandler(w http.ResponseWriter, r *http.Request) {
	var c storage.Comment
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = api.db.AddComm(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
