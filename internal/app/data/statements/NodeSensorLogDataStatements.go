package statements

import "github.com/ltruelove/gohome/config"

type NodeSensorLogDataStatements struct {
	DbType string
}

func NewNodeSensorLogDataStatements(config *config.Configuration) CrudStatement {
	return &NodeSensorLogDataStatements{DbType: config.DbType}
}

func (statements *NodeSensorLogDataStatements) SelectAll() string {
	if statements.DbType == "mysql" {
		return `SELECT
		Id,
		NodeId,
		DateLogged
		FROM NodeSensorLog`
	} else {
		return `SELECT
		id,
		nodeid,
		datelogged
		FROM nodesensorlog`
	}
}

func (statements *NodeSensorLogDataStatements) SelectById() string {
	if statements.DbType == "mysql" {
		return `SELECT
		Id,
		NodeId,
		DateLogged
		FROM NodeSensorLog
		WHERE Id = ?`
	} else {
		return `SELECT
		id,
		nodeid,
		datelogged
		FROM nodesensorlog
		WHERE id = $1`
	}
}

func (statements *NodeSensorLogDataStatements) SelectByParentId() string {
	if statements.DbType == "mysql" {
		return `SELECT
		Id,
		NodeId,
		DateLogged
		FROM NodeSensorLog
		WHERE NodeId = ?`
	} else {
		return `SELECT
		id,
		nodeid,
		datelogged
		FROM nodesensorlog
		WHERE nodeid = $1`
	}
}

func (statements *NodeSensorLogDataStatements) SelectBySecondParentId() string {
	return "NodeSensorLog table has no second parent ID"
}

func (statements *NodeSensorLogDataStatements) Insert() string {
	if statements.DbType == "mysql" {
		return `INSERT INTO NodeSensorLog
		(NodeId, DateLogged)
		VALUES (?, ?); SELECT LAST_INSERT_ID();`
	} else {
		return `INSERT INTO nodesensorlog
		(nodeid, datelogged)
		VALUES ($1, $2) RETURNING id`
	}
}

func (statements *NodeSensorLogDataStatements) Update() string {
	if statements.DbType == "mysql" {
		return `UPDATE NodeSensorLog
		SET NodeId = ?, DateLogged = ?
		WHERE Id = ?`
	} else {
		return `UPDATE nodesensorlog
		SET nodeid = $1, datelogged = $2
		WHERE id = $3`
	}
}

func (statements *NodeSensorLogDataStatements) Delete() string {
	if statements.DbType == "mysql" {
		return `DELETE FROM NodeSensorLog WHERE Id = ?`
	} else {
		return `DELETE FROM nodesensorlog WHERE id = $1`
	}
}

func (statements *NodeSensorLogDataStatements) DeleteAll() string {
	if statements.DbType == "mysql" {
		return `DELETE FROM NodeSensorLog`
	} else {
		return `DELETE FROM nodesensorlog`
	}
}

func (statements *NodeSensorLogDataStatements) DeleteByParentId() string {
	if statements.DbType == "mysql" {
		return `DELETE FROM NodeSensorLog WHERE NodeId = ?`
	} else {
		return `DELETE FROM nodesensorlog WHERE nodeid = $1`
	}
}

func (statements *NodeSensorLogDataStatements) DeleteBySecondParentId() string {
	return "NodeSensorLog table has no second parent ID deletion"
}
