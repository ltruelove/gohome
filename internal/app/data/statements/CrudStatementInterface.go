package statements

type CrudStatement interface {
	SelectAll() string
	SelectById() string
	SelectByParentId() string
	SelectBySecondParentId() string
	Insert() string
	Update() string
	Delete() string
	DeleteAll() string
	DeleteByParentId() string
}
