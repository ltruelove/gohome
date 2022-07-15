package models

type model interface {
	IsValid(checkId bool) (bool, error)
}
