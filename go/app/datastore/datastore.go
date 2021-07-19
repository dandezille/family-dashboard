package datastore

import (
	"server/app/models"
)

type Datastore interface {
	Close() error
	FindById(id int64) (*models.Activity, error)
	Find() ([]*models.Activity, error)
	Create(activity *models.Activity) error
	Update(activity *models.Activity) error
	Delete(id int64) error
}
