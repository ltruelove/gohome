package repository

import (
	"github.com/ltruelove/gohome/internal/app/data"
	"github.com/ltruelove/gohome/internal/app/models"
)

type crudRepositoryAdapter struct {
	d data.CrudDataInterface
}

func NewCrudRepositoryFromData(d data.CrudDataInterface) CrudRepository {
	return &crudRepositoryAdapter{d: d}
}

func (r *crudRepositoryAdapter) SelectAll() ([]models.Model, error) {
	return r.d.SelectAll()
}

func (r *crudRepositoryAdapter) SelectByParentId(id int) ([]models.Model, error) {
	return r.d.SelectByParentId(id)
}

func (r *crudRepositoryAdapter) SelectById(id int) (models.Model, error) {
	return r.d.SelectById(id)
}

func (r *crudRepositoryAdapter) Insert(data models.Model) (models.Model, error) {
	return r.d.Insert(data)
}

func (r *crudRepositoryAdapter) Update(data models.Model) error {
	return r.d.Update(data)
}

func (r *crudRepositoryAdapter) Delete(id int) error {
	return r.d.Delete(id)
}
func (r *crudRepositoryAdapter) DeleteByParentId(id int) error {
	return r.d.DeleteByParentId(id)
}
func (r *crudRepositoryAdapter) DeleteBySecondParentId(id int) error {
	return r.d.DeleteBySecondParentId(id)
}
