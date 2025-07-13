package statements

import "github.com/ltruelove/gohome/config"

type SwitchTypeStatements struct {
	DbType string
}

func NewSwitchTypeStatements(config *config.Configuration) CrudStatement {
	return &SwitchTypeStatements{DbType: config.DbType}
}

func (statements *SwitchTypeStatements) SelectAll() string {
	if statements.DbType == "mysql" {
		return "SELECT Id, TypeName FROM SwitchType"
	} else {
		return "SELECT id, name FROM switchtype"
	}

}

func (statements *SwitchTypeStatements) SelectById() string {
	if statements.DbType == "mysql" {
		return "SELECT Id, TypeName FROM SwitchType WHERE Id = ?"
	} else {
		return "SELECT id, name FROM switchtype WHERE id = $1"
	}
}

func (statements *SwitchTypeStatements) SelectByParentId() string {
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

func (statements *SwitchTypeStatements) SelectBySecondParentId() string {
	return "Not implemented for SwitchTypeStatements"
}

func (statements *SwitchTypeStatements) Insert() string {
	if statements.DbType == "mysql" {
		return "INSERT INTO SwitchType (TypeName) VALUES (?)"
	} else {
		return "INSERT INTO switchtype (name) VALUES ($1)"
	}
}

func (statements *SwitchTypeStatements) Update() string {
	if statements.DbType == "mysql" {
		return "UPDATE SwitchType SET TypeName = ? WHERE Id = ?"
	} else {
		return "UPDATE switchtype SET name = $1 WHERE id = $2"
	}
}

func (statements *SwitchTypeStatements) Delete() string {
	if statements.DbType == "mysql" {
		return "DELETE FROM SwitchType WHERE Id = ?"
	} else {
		return "DELETE FROM switchtype WHERE id = $1"
	}
}

func (statements *SwitchTypeStatements) DeleteByParentId() string {
	if statements.DbType == "mysql" {
		return "DELETE FROM SwitchType WHERE Id = ?"
	} else {
		return "DELETE FROM switchtype WHERE id = $1"
	}
}

func (statements *SwitchTypeStatements) DeleteAll() string {
	if statements.DbType == "mysql" {
		return "DELETE FROM SwitchType"
	} else {
		return "DELETE FROM switchtype"
	}
}
