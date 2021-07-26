package datastore

import (
	"time"

	"server/app/models"
)

type Datastore interface {
	Close() error
	FindById(id int64) (*models.Activity, error)
	FindByTime(time.Time) ([]*models.Activity, error)
	Find() ([]*models.Activity, error)
	Create(activity *models.Activity) error
	Update(activity *models.Activity) error
	Delete(id int64) error
}
