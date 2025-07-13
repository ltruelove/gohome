/*
MySQL:

CREATE TABLE IF NOT EXISTS ControlPointNodes (

	Id	INT NOT NULL UNIQUE AUTO_INCREMENT,
	ControlPointId	INT NOT NULL,
	NodeId	INT NOT NULL,
	FOREIGN KEY(NodeId) REFERENCES Node(Id),
	FOREIGN KEY(ControlPointId) REFERENCES ControlPoint(Id),
	PRIMARY KEY(Id)

);
*/
package statements

import "github.com/ltruelove/gohome/config"

type ControlPointNodeDataStatements struct {
	DbType string
}

func NewControlPointNodeDataStatements(config *config.Configuration) CrudStatement {
	return &ControlPointNodeDataStatements{DbType: config.DbType}
}

func (statements *ControlPointNodeDataStatements) SelectAll() string {
	if statements.DbType == "mysql" {
		return `SELECT
		Id,
		ControlPointId,
		NodeId
		FROM ControlPointNodes`
	} else {
		return `SELECT
		id,
		controlpointid,
		nodeid
		FROM controlpointnodes`
	}
}

func (statements *ControlPointNodeDataStatements) SelectById() string {
	if statements.DbType == "mysql" {
		return `SELECT
		Id,
		ControlPointId,
		NodeId
		FROM ControlPointNodes
		WHERE Id = ?`
	} else {
		return `SELECT
		id,
		controlpointid,
		nodeid
		FROM controlpointnodes
		WHERE id = $1`
	}
}

func (statements *ControlPointNodeDataStatements) SelectByParentId() string {
	if statements.DbType == "mysql" {
		return `SELECT
		Id,
		ControlPointId,
		NodeId
		FROM ControlPointNodes
		WHERE ControlPointId = ?`
	} else {
		return `SELECT
		id,
		controlpointid,
		nodeid
		FROM controlpointnodes
		WHERE controlpointid = $1`
	}
}

func (statements *ControlPointNodeDataStatements) SelectBySecondParentId() string {
	if statements.DbType == "mysql" {
		return `SELECT
		Id,
		ControlPointId,
		NodeId
		FROM ControlPointNodes
		WHERE NodeId = ?`
	} else {
		return `SELECT
		id,
		controlpointid,
		nodeid
		FROM controlpointnodes
		WHERE nodeid = $1`
	}
}

func (statements *ControlPointNodeDataStatements) Insert() string {
	if statements.DbType == "mysql" {
		return `INSERT INTO ControlPointNodes (ControlPointId, NodeId)
		VALUES (?, ?); SELECT LAST_INSERT_ID()`
	} else {
		return `INSERT INTO controlpointnodes (controlpointid, nodeid)
		VALUES ($1, $2) RETURNING id`
	}
}

func (statements *ControlPointNodeDataStatements) Update() string {
	if statements.DbType == "mysql" {
		return `UPDATE ControlPointNodes
		SET ControlPointId = ?, NodeId = ?
		WHERE Id = ?`
	} else {
		return `UPDATE controlpointnodes
		SET controlpointid = $1, nodeid = $2
		WHERE id = $3`
	}
}

func (statements *ControlPointNodeDataStatements) Delete() string {
	if statements.DbType == "mysql" {
		return `DELETE FROM ControlPointNodes WHERE Id = ?`
	} else {
		return `DELETE FROM controlpointnodes WHERE id = $1`
	}
}

func (statements *ControlPointNodeDataStatements) DeleteAll() string {
	if statements.DbType == "mysql" {
		return `DELETE FROM ControlPointNodes`
	} else {
		return `DELETE FROM controlpointnodes`
	}
}

func (statements *ControlPointNodeDataStatements) DeleteByParentId() string {
	if statements.DbType == "mysql" {
		return `DELETE FROM ControlPointNodes WHERE ControlPointId = ?`
	} else {
		return `DELETE FROM controlpointnodes WHERE controlpointid = $1`
	}
}

func (statements *ControlPointNodeDataStatements) DeleteBySecondParentId() string {
	if statements.DbType == "mysql" {
		return `DELETE FROM ControlPointNodes WHERE NodeId = ?`
	} else {
		return `DELETE FROM controlpointnodes WHERE nodeid = $1`
	}
}
