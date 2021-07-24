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
	db datastore.Datastore
}

func New() (App, error) {
	db, err := datastore.NewSqliteDatastore("./data.db")
	return app{db}, err
}

func (a app) Close() error {
	return a.db.Close()
}
