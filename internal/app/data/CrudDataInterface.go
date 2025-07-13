package data

import (
	"database/sql"

	"github.com/ltruelove/gohome/internal/app/data/statements"
	"github.com/ltruelove/gohome/internal/app/models"
)

type CrudDataInterface interface {
	DB() *sql.DB
	Stmt() statements.CrudStatement
	SelectAll() ([]models.Model, error)
	SelectByParentId(id int) ([]models.Model, error)
	SelectBySecondParentId(id int) ([]models.Model, error)
	SelectById(id int) (models.Model, error)
	Insert(data models.Model) (models.Model, error)
	Update(data models.Model) error
	Delete(id int) error
	DeleteAll() error
	DeleteByParentId(id int) error
	DeleteBySecondParentId(id int) error
}
