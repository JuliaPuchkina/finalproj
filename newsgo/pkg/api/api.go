package api

import (
	"encoding/json"
	storage "finalproj/newsgo/pkg/storage"
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
	// получить новости на странице n
	api.r.HandleFunc("/news", api.postsHandler).Queries("page", "{page}").Methods(http.MethodGet, http.MethodOptions)
	// веб-приложение
	api.r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./webapp"))))
	// добавление новости в базу
	api.r.HandleFunc("/news", api.addPostHandler).Methods(http.MethodPost, http.MethodOptions)
	// поиск по названию новости
	api.r.HandleFunc("/news", api.newsTitleHandler).Queries("title", "{title}").Methods(http.MethodGet, http.MethodOptions)
}

// postsHandler возвращает все публикации на странице
func (api *API) postsHandler(w http.ResponseWriter, r *http.Request) {
	s := mux.Vars(r)["page"]
	n, _ := strconv.Atoi(s)
	posts, err := api.db.Posts(n)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	bytes, err := json.Marshal(posts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
}

// addPostHandler добавляет публикацию в базу
func (api *API) addPostHandler(w http.ResponseWriter, r *http.Request) {
	var p storage.Post
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = api.db.AddPost(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// newsTitleHandler осуществляет поиск по названию публикации
// и возвращает публикации с соответствующим заголовком
func (api *API) newsTitleHandler(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"] // ??? почему тут это не работает, а работает в страницах? где ошибка?
	posts, err := api.db.NewsTitle(title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	bytes, err := json.Marshal(posts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
}

// журналирование запросов
// TODO как его вставить после обработчика???
func (api *API) logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

// добавление сквозного идентификатора запроса
// TODO как его вставить до обработчика???
func (api *API) requestIdMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}
