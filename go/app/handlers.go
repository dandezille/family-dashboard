package app

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"

	"server/app/models"
)

func GetHome(w http.ResponseWriter, r *http.Request) {
	err := renderPage("app/views/pages/home.html", w, nil)
	if handleError(w, err) {
		return
	}
}

func (a app) GetActivities(w http.ResponseWriter, r *http.Request) {
	activities, err := a.activities.Find()
	if handleError(w, err) {
		return
	}

	err = renderPage("app/views/activities/index.html", w, struct {
		Activities []*models.Activity
	}{
		Activities: activities,
	})

	if handleError(w, err) {
		return
	}
}

func (a app) PostActivities(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if handleError(w, err) {
		return
	}

	log.Printf("Params: %v\n", r.PostForm)

	start, err := time.Parse("15:04", r.PostForm.Get("start"))
	if handleError(w, err) {
		return
	}

	activity := &models.Activity{
		Symbol: r.PostForm.Get("symbol"),
		Start:  start,
		Note:   r.PostForm.Get("note"),
	}

	log.Printf("Create: %v\n", activity)

	err = a.activities.Create(activity)
	if handleError(w, err) {
		return
	}

	log.Printf("Created: %v\n", activity)
	http.Redirect(w, r, "/activities", 303)
}

func GetNewActivity(w http.ResponseWriter, r *http.Request) {
	err := renderPage("app/views/activities/new.html", w, nil)
	if handleError(w, err) {
		return
	}
}

func (a app) GetEditActivity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if handleError(w, err) {
		return
	}

	activity, err := a.activities.FindById(id)
	if handleError(w, err) {
		return
	}

	log.Printf("Edit: %v\n", activity)
	err = renderPage("app/views/activities/edit.html", w, activity)
	if handleError(w, err) {
		return
	}
}

func (a app) PostActivity(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if handleError(w, err) {
		return
	}

	log.Printf("Params: %v\n", r.PostForm)

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if handleError(w, err) {
		return
	}

	activity, err := a.activities.FindById(id)
	if handleError(w, err) {
		return
	}

	log.Printf("Update: %v\n", activity)

	activity.Symbol = r.PostForm.Get("symbol")
	activity.Start, err = time.Parse("15:04", r.PostForm.Get("start"))
	if handleError(w, err) {
		return
	}

	activity.Note = r.PostForm.Get("note")

	log.Printf("Updated: %v\n", activity)

	a.activities.Update(activity)
	http.Redirect(w, r, "/activities", 303)
}

func (a app) DeleteActivity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if handleError(w, err) {
		return
	}

	err = a.activities.Delete(id)
	if handleError(w, err) {
		return
	}

	http.Redirect(w, r, "/activities", 303)
}

func (a app) ApiGetActivities(w http.ResponseWriter, r *http.Request) {
	data := models.Activities{
		Current: models.Activity{
			Symbol: "A",
			Start:  time.Now().Local(),
		},
		Next: models.Activity{
			Symbol: "B",
			Start:  time.Now().Local().Add(time.Minute * time.Duration(1)),
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func ApiGetWeather(w http.ResponseWriter, r *http.Request) {
	data := models.Temperature{
		Min: 12,
		Now: 13,
		Max: 14,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
func handleError(w http.ResponseWriter, err error) bool {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error:", err)
	}

	return err != nil
}

func renderPage(page string, w http.ResponseWriter, data interface{}) error {
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

	return templates.ExecuteTemplate(w, "layout", data)
}
