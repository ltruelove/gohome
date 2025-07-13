package viewModels

import "github.com/ltruelove/gohome/internal/app/models"

type ViewModel interface {
	ImportModel(model *models.Model)
}
