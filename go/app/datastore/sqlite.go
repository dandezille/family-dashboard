package datastore

import (
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"

	"server/app/models"
)

const schema = `
CREATE TABLE IF NOT EXISTS activities
(
	id integer primary key,
	symbol text,
	start timestamp,
	note text
)
`

type datastore struct {
	db *sqlx.DB
}

func NewSqliteDatastore(path string) (Datastore, error) {
	db := sqlx.MustConnect("sqlite3", path)
	db.MustExec(schema)
	return &datastore{db}, nil
}

func (s *datastore) Close() error {
	return s.db.Close()
}

func (s *datastore) FindById(id int64) (*models.Activity, error) {
	activity := models.Activity{}
	query := "SELECT id, symbol, start, note FROM activities WHERE id=?"

	err := s.db.Get(&activity, query, id)
	if err != nil {
		return nil, err
	}

	return &activity, nil
}

func (s *datastore) FindByTime(time.Time) ([]*models.Activity, error) {
	activities := []*models.Activity{}
	query := "SELECT id, symbol, start, note FROM activities ORDER BY start"

	err := s.db.Select(&activities, query)
	if err != nil {
		return nil, err
	}

	return activities, nil
}

func (s *datastore) Find() ([]*models.Activity, error) {
	activities := []*models.Activity{}
	query := "SELECT id, symbol, start, note FROM activities ORDER BY start"

	err := s.db.Select(&activities, query)
	if err != nil {
		return nil, err
	}

	return activities, nil
}

func (s *datastore) Create(a *models.Activity) error {
	query := "INSERT INTO activities(symbol, start, note) VALUES (?, ?, ?) RETURNING id"
	return s.db.QueryRow(query, a.Symbol, a.Start, a.Note).Scan(&a.ID)
}

func (s *datastore) Update(a *models.Activity) error {
	query := "UPDATE activities SET symbol=?, start=?, note=? WHERE id=?"
	_, err := s.db.Exec(query, a.Symbol, a.Start, a.Note, a.ID)
	return err
}

func (s *datastore) Delete(id int64) error {
	query := "DELETE FROM activities WHERE id=?"
	_, err := s.db.Exec(query, id)
	return err
}
