package repository

import "github.com/ltruelove/gohome/internal/app/viewModels"

type CompoundRepository interface {
	FetchViewNodeSensorDataByViewId(viewId int) ([]viewModels.ViewNodeSensorVM, error)
	FetchViewNodeSwitchDataByViewId(viewId int) ([]viewModels.ViewNodeSwitchVM, error)
}
