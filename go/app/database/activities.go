package database

import (
	"log"

	"server/app/models"
)

func (d *DB) GetActivities() []*models.Activity {
	rows, err := d.db.Query(getActivities)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var activities []*models.Activity
	for rows.Next() {
		a := &models.Activity{}
		err := rows.Scan(&a.ID, &a.Symbol, &a.Time, &a.Note)
		if err != nil {
			log.Fatal(err)
		}
		activities = append(activities, a)
	}

	return activities
}

func (d *DB) Create(a *models.Activity) {
	err := d.db.QueryRow(insertActivity, a.Symbol, a.Time, a.Note).Scan(&a.ID)
	if err != nil {
		log.Fatal(err)
	}
}
