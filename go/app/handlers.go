package app

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"

	"server/app/datastore"
	"server/app/models"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	renderPage("app/views/pages/home.html", w, nil)
}

func HandleGetActivities(w http.ResponseWriter, r *http.Request) {
	db := datastore.Open("data.db")
	defer db.Close()

	renderPage("app/views/activities/index.html", w, struct {
		Activities []*models.Activity
	}{
		Activities: db.GetActivities(),
	})
}

func HandlePostActivities(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Params: %v\n", r.PostForm)

	start, err := time.Parse("15:04", r.PostForm.Get("start"))
	if err != nil {
		log.Fatal(err)
	}

	activity := &models.Activity{
		Symbol: r.PostForm.Get("symbol"),
		Start:  start,
		Note:   r.PostForm.Get("note"),
	}

	log.Printf("Create: %v\n", activity)

	db := datastore.Open("data.db")
	defer db.Close()
	db.CreateActivity(activity)

	log.Printf("Created: %v\n", activity)
	http.Redirect(w, r, "/activities", 303)
}

func HandleGetNewActivity(w http.ResponseWriter, r *http.Request) {
	renderPage("app/views/activities/new.html", w, nil)
}

func HandleGetEditActivity(w http.ResponseWriter, r *http.Request) {
	db := datastore.Open("data.db")
	defer db.Close()

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	activity := db.GetActivity(id)
	log.Printf("Edit: %v\n", activity)

	renderPage("app/views/activities/edit.html", w, activity)
}

func HandlePostActivity(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Params: %v\n", r.PostForm)

	db := datastore.Open("data.db")
	defer db.Close()

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	activity := db.GetActivity(id)
	log.Printf("Update: %v\n", activity)

	activity.Symbol = r.PostForm.Get("symbol")
	activity.Start, err = time.Parse("15:04", r.PostForm.Get("start"))
	if err != nil {
		log.Fatal(err)
	}
	activity.Note = r.PostForm.Get("note")

	log.Printf("Updated: %v\n", activity)

	db.UpdateActivity(activity)
	http.Redirect(w, r, "/activities", 303)
}

func HandleDeleteActivity(w http.ResponseWriter, r *http.Request) {
	db := datastore.Open("data.db")
	defer db.Close()

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	db.DeleteActivity(id)
	http.Redirect(w, r, "/activities", 303)
}

func renderPage(page string, w http.ResponseWriter, data interface{}) {
	log.Printf("Render page %s\n", page)

	// partials
	partialsRoot := strings.LastIndex(page, "/")
	partialsGlob := page[:partialsRoot] + "/_*.html"
	log.Printf("Partials glob %s\n", partialsGlob)
	templates := template.Must(template.ParseGlob(partialsGlob))

	// layouts
	templates = template.Must(templates.ParseGlob("app/views/layouts/*.html"))

	// page
	templates = template.Must(templates.ParseFiles(page))

	err := templates.ExecuteTemplate(w, "layout", data)
	if err != nil {
		log.Fatal(err)
	}
}
