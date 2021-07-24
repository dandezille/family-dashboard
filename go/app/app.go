package app

import (
	"github.com/gorilla/mux"

	"server/app/datastore"
)

type App interface {
	Close() error
	SetupRouter() (*mux.Router, error)
}

type app struct {
	activities datastore.Datastore
}

func New() (App, error) {
	activities, err := datastore.NewSqliteDatastore("./activities.db")
	return app{activities}, err
}

func (a app) Close() error {
	return a.activities.Close()
}
