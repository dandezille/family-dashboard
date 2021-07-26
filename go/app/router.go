package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (a app) SetupRouter() (*mux.Router, error) {
	r := mux.NewRouter()
	r.Use(logRequests)

	a.setupRoutes(r)
	a.setupApiRoutes(r.PathPrefix("/api").Subrouter())
	setupStatic(r)
	setupJavascript(r)

	return r, nil
}

func (a app) setupRoutes(r *mux.Router) {
	r.HandleFunc("/", GetHome).Methods("GET")
	r.HandleFunc("/activities", a.GetActivities).Methods("GET")
	r.HandleFunc("/activities", a.PostActivities).Methods("POST")
	r.HandleFunc("/activities/new", GetNewActivity).Methods("GET")
	r.HandleFunc("/activities/{id}/edit", a.GetEditActivity).Methods("GET")
	r.HandleFunc("/activities/{id}", a.PostActivity).Methods("POST")
	r.HandleFunc("/activities/{id}", a.DeleteActivity).Methods("DELETE")
}

func (a app) setupApiRoutes(r *mux.Router) {
	r.HandleFunc("/activities", a.ApiGetActivities)
	r.HandleFunc("/weather", ApiGetWeather)
}

func setupStatic(r *mux.Router) {
	static := r.PathPrefix("/static/")
	static.Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./app/static"))))
}

func setupJavascript(r *mux.Router) {
	js := r.PathPrefix("/javascript/")
	js.Handler(http.StripPrefix("/javascript/", http.FileServer(http.Dir("./app/javascript"))))
}

func logRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s: %s\n", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
