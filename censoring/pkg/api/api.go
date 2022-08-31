package api

import (
	"finalproj/censoring/pkg/censor"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// API приложения.
type API struct {
	r *mux.Router // маршрутизатор запросов
}

// Конструктор API.
func New() *API {
	api := API{}
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
	api.r.HandleFunc("/comments/{test}", api.censoreHandler).Methods(http.MethodGet, http.MethodOptions)
	// принятие текста выглядит неверным, но пока оставлю так
}

// censoreHandler возвращает отметку о цензурности комментария
func (api *API) censoreHandler(w http.ResponseWriter, r *http.Request) {
	text := mux.Vars(r)["text"]

	var resp string

	if censor.IsCommentBad(text) {
		resp = strconv.Itoa(http.StatusBadRequest)
	} else {
		resp = strconv.Itoa(http.StatusOK)
	}

	w.Write([]byte(resp))
}
