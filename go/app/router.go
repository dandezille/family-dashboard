package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"server/app/api"
)

func (a app) SetupRouter() (*mux.Router, error) {
	r := mux.NewRouter()
	r.Use(logRequests)

	a.setupRoutes(r)
	setupStatic(r)
	setupJavascript(r)
	api.SetupRoutes(r.PathPrefix("/api").Subrouter())

	return r, nil
}

func (a app) setupRoutes(r *mux.Router) {
	r.HandleFunc("/", HandleHome).Methods("GET")
	r.HandleFunc("/activities", a.HandleGetActivities).Methods("GET")
	r.HandleFunc("/activities", a.HandlePostActivities).Methods("POST")
	r.HandleFunc("/activities/new", HandleGetNewActivity).Methods("GET")
	r.HandleFunc("/activities/{id}/edit", a.HandleGetEditActivity).Methods("GET")
	r.HandleFunc("/activities/{id}", a.HandlePostActivity).Methods("POST")
	r.HandleFunc("/activities/{id}", a.HandleDeleteActivity).Methods("DELETE")
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
