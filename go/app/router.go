package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"server/app/api"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()
	r.Use(logRequests)

	setupRoutes(r)
	setupStatic(r)
	setupJavascript(r)
	api.SetupRoutes(r.PathPrefix("/api").Subrouter())

	return r
}

func setupRoutes(r *mux.Router) {
	r.HandleFunc("/", HandleHome)
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
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
