package repository

import (
	"github.com/ltruelove/gohome/internal/app/data"
	"github.com/ltruelove/gohome/internal/app/viewModels"
)

type compoundRepositoryAdapter struct {
	d *data.CompoundData
}

func NewCompoundRepositoryFromData(d *data.CompoundData) CompoundRepository {
	return &compoundRepositoryAdapter{d: d}
}

func (r *compoundRepositoryAdapter) FetchViewNodeSensorDataByViewId(viewId int) ([]viewModels.ViewNodeSensorVM, error) {
	return r.d.FetchViewNodeSensorDataByViewId(viewId)
}

func (r *compoundRepositoryAdapter) FetchViewNodeSwitchDataByViewId(viewId int) ([]viewModels.ViewNodeSwitchVM, error) {
	return r.d.FetchViewNodeSwitchDataByViewId(viewId)
}
