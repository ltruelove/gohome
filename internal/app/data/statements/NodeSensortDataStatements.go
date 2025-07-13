package statements

import "github.com/ltruelove/gohome/config"

type NodeSensorDataStatements struct {
	DbType string
}

func NewNodeSensorDataStatements(config *config.Configuration) CrudStatement {
	return &NodeSensorDataStatements{DbType: config.DbType}
}

func (statements *NodeSensorDataStatements) SelectAll() string {
	if statements.DbType == "mysql" {
		return `SELECT
		Id,
		NodeId,
		SensorTypeId,
		SensorName,
		Pin,
		DHTType
		FROM NodeSensor`
	} else {
		return `SELECT
		id,
		nodeid,
		sensortypeid,
		name,
		pin,
		dhttype
		FROM nodesensor`
	}
}

func (statements *NodeSensorDataStatements) SelectById() string {
	if statements.DbType == "mysql" {
		return `SELECT
		Id,
		NodeId,
		SensorTypeId,
		SensorName,
		Pin,
		DHTType
		FROM NodeSensor
		WHERE Id = ?`
	} else {
		return `SELECT
		id,
		nodeid,
		sensortypeid,
		name,
		pin,
		dhttype
		FROM nodesensor
		WHERE id = $1`
	}
}

// SelectByParentId retrieves NodeSensor records by NodeId
func (statements *NodeSensorDataStatements) SelectByParentId() string {
	if statements.DbType == "mysql" {
		return `SELECT
			Id,
			NodeId,
			SensorTypeId,
			SensorName,
			Pin,
			DHTType
			FROM NodeSensor
			WHERE NodeId = ?`
	} else {
		return `SELECT
			id,
			nodeid,
			sensortypeid,
			name,
			pin,
			dhttype
			FROM nodesensor
			WHERE nodeid = $1`
	}
}

// SelectBySecondParentId retrieves NodeSensor records by SensorTypeId
func (statements *NodeSensorDataStatements) SelectBySecondParentId() string {
	if statements.DbType == "mysql" {
		return `SELECT
			Id,
			NodeId,
			SensorTypeId,
			SensorName,
			Pin,
			DHTType
			FROM NodeSensor
			WHERE SensorTypeId = ?`
	} else {
		return `SELECT
			id,
			nodeid,
			sensortypeid,
			name,
			pin,
			dhttype
			FROM nodesensor
			WHERE sensortypeid = $1`
	}
}

func (statements *NodeSensorDataStatements) Insert() string {
	if statements.DbType == "mysql" {
		return `INSERT INTO NodeSensor
		(NodeId, SensorTypeId, SensorName, Pin, DHTType)
		VALUES (?, ?, ?, ?, ?); SELECT LAST_INSERT_ID();`
	} else {
		return `INSERT INTO nodesensor
		(nodeid, sensortypeid, name, pin, dhttype)
		VALUES ($1, $2, $3, $4, $5) RETURNING id`
	}
}

func (statements *NodeSensorDataStatements) Update() string {
	if statements.DbType == "mysql" {
		return `UPDATE NodeSensor
		SET NodeId = ?, SensorTypeId = ?, SensorName = ?, Pin = ?, DHTType = ?
		WHERE Id = ?`
	} else {
		return `UPDATE nodesensor
		SET nodeid = $1, sensortypeid = $2, name = $3, pin = $4, dhttype = $5
		WHERE id = $6`
	}
}

func (statements *NodeSensorDataStatements) Delete() string {
	if statements.DbType == "mysql" {
		return `DELETE FROM NodeSensor WHERE Id = ?`
	} else {
		return `DELETE FROM nodesensor WHERE id = $1`
	}
}

func (statements *NodeSensorDataStatements) DeleteAll() string {
	if statements.DbType == "mysql" {
		return `DELETE FROM NodeSensor`
	} else {
		return `DELETE FROM nodesensor`
	}
}

func (statements *NodeSensorDataStatements) DeleteByParentId() string {
	if statements.DbType == "mysql" {
		return `DELETE FROM NodeSensor WHERE NodeId = ?`
	} else {
		return `DELETE FROM nodesensor WHERE nodeid = $1`
	}
}

func (statements *NodeSensorDataStatements) DeleteBySecondParentId() string {
	if statements.DbType == "mysql" {
		return `DELETE FROM NodeSensor WHERE SensorTypeId = ?`
	} else {
		return `DELETE FROM nodesensor WHERE sensortypeid = $1`
	}
}
