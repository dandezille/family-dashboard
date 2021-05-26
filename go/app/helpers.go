package app

import (
	"encoding/json"
	"log"
	"net/http"
	"server/app/models"
)

func parse(w http.ResponseWriter, r *http.Request, data interface{}) error {
	return json.NewDecoder(r.Body).Decode(data)
}

func sendResponse(w http.ResponseWriter, _ *http.Request, data interface{}, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	if data == nil {
		return
	}

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("Cannot format json. err=%v\n", err)
	}
}

func mapActivityToJSON(a *models.Activity) models.JsonActivity {
	return models.JsonActivity{
		ID:     a.ID,
		Symbol: a.Symbol,
		Note:   a.Note,
	}
}
