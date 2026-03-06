package repository

import "github.com/ltruelove/gohome/internal/app/models"

type CrudRepository interface {
	SelectAll() ([]models.Model, error)
	SelectByParentId(id int) ([]models.Model, error)
	SelectById(id int) (models.Model, error)
	Insert(data models.Model) (models.Model, error)
	Update(data models.Model) error
	Delete(id int) error
	DeleteByParentId(id int) error
	DeleteBySecondParentId(id int) error
}
