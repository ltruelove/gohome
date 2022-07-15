package models

type Model interface {
	IsValid(checkId bool) (bool, error)
}
