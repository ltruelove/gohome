package data

import "github.com/ltruelove/gohome/config"

type SwitchTypeStatements struct {
	DbType string
}

func NewSwitchTypeStatements(config *config.Configuration) *SwitchTypeStatements {
	return &SwitchTypeStatements{DbType: config.DbType}
}

func (statements *SwitchTypeStatements) SelectAllSwitchTypes() string {
	if statements.DbType == "mysql" {
		return "SELECT Id, TypeName FROM SwitchType"
	} else {
		return "SELECT id, name FROM switchtype"
	}

}

func (statements *SwitchTypeStatements) SelectSwitchTypeById() string {
	if statements.DbType == "mysql" {
		return "SELECT Id, TypeName FROM SwitchType WHERE Id = ?"
	} else {
		return "SELECT id, name FROM switchtype WHERE id = $1"
	}
}

func (statements *SwitchTypeStatements) SelectSwitchTypeData() string {
	if statements.DbType == "mysql" {
		return `SELECT
		Id, 
		TypeName, 
		ValueType 
		FROM SwitchTypeData 
		WHERE SwitchTypeId = ?`
	} else {
		return `SELECT
		id, 
		name, 
		valuetype 
		FROM switchtypedata 
		WHERE switchtypeid = $1`
	}
}
