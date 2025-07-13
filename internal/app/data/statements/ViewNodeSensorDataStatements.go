package statements

import "github.com/ltruelove/gohome/config"

type ViewNodeSensorDataStatements struct {
	DbType string
}

func NewViewNodeSensorDataStatements(config *config.Configuration) CrudStatement {
	return &ViewNodeSensorDataStatements{DbType: config.DbType}
}

func (statements *ViewNodeSensorDataStatements) SelectAll() string {
	if statements.DbType == "mysql" {
		return `SELECT
			Id,
			NodeId,
			ViewId,
			NodeSensorId,
			Name
			FROM ViewNodeSensorData`

	} else {
		return `SELECT
			id,
			nodeid,
			viewid,
			nodesensorid,
			name
			FROM viewnodesensordata`
	}

}

func (statements *ViewNodeSensorDataStatements) SelectById() string {
	if statements.DbType == "mysql" {
		return `SELECT
			Id,
			NodeId,
			ViewId,
			NodeSensorId,
			Name
			FROM ViewNodeSensorData
			WHERE id = ?`

	} else {
		return `SELECT
			id,
			nodeid,
			viewid,
			nodesensorid,
			name
			FROM viewnodesensordata
			WHERE id = $1`
	}

}

func (statements *ViewNodeSensorDataStatements) SelectByParentId() string {
	if statements.DbType == "mysql" {
		return `SELECT
			Id,
			NodeId,
			ViewId,
			NodeSensorId,
			Name
			FROM ViewNodeSensorData
			WHERE NodeId = ?`

	} else {
		return `SELECT
			id,
			nodeid,
			viewid,
			nodesensorid,
			name
			FROM viewnodesensordata
			WHERE nodeid = $1`
	}
}

func (statements *ViewNodeSensorDataStatements) SelectBySecondParentId() string {
	if statements.DbType == "mysql" {
		return `SELECT
		d.Id,
		d.NodeId,
		d.ViewId,
		d.NodeSensorId,
		d.Name,
		n.Name As NodeName,
		ns.Name AS SensorName,
		st.Name AS SensorTypeName
		FROM ViewNodeSensorData AS d
		INNER JOIN Node AS n ON n.Id = d.NodeId
		INNER JOIN NodeSensor AS ns ON ns.Id = d.NodeSensorId
		INNER JOIN SensorType AS st on st.Id = ns.SensorTypeId
		WHERE d.ViewId = ?`

	} else {
		return `SELECT
		d.id,
		d.nodeid,
		d.viewid,
		d.nodesensorid,
		d.name,
		n.name As nodename,
		ns.name AS sensorname,
		st.name AS sensortypename
		FROM viewnodesensordata AS d
		INNER JOIN node AS n ON n.id = d.nodeid
		INNER JOIN nodesensor AS ns ON ns.id = d.nodesensorid
		INNER JOIN sensortype AS st on st.id = ns.sensortypeid
		WHERE d.viewid = $1`
	}
}

func (statements *ViewNodeSensorDataStatements) Insert() string {
	if statements.DbType == "mysql" {
		return `INSERT INTO ViewNodeSensorData
			(NodeId, ViewId, NodeSensorId, Name)
			VALUES (?, ?, ?, ?); SELECT LAST_INSERT_ID();`
	} else {
		return `INSERT INTO viewnodesensordata
			(nodeid, viewid, nodesensorid, name)
			VALUES ($1, $2, $3, $4) RETURNING id`
	}
}

func (statements *ViewNodeSensorDataStatements) Update() string {
	if statements.DbType == "mysql" {
		return `UPDATE ViewNodeSensorData
			SET NodeId = ?, ViewId = ?, NodeSensorId = ?, Name = ?
			WHERE Id = ?`
	} else {
		return `UPDATE viewnodesensordata
			SET nodeid = $1, viewid = $2, nodesensorid = $3, name = $4
			WHERE id = $5`
	}
}

func (statements *ViewNodeSensorDataStatements) Delete() string {
	if statements.DbType == "mysql" {
		return `DELETE FROM ViewNodeSensorData
			WHERE Id = ?`
	} else {
		return `DELETE FROM viewnodesensordata
			WHERE id = $1`
	}
}

func (statements *ViewNodeSensorDataStatements) DeleteByParentId() string {
	if statements.DbType == "mysql" {
		return `DELETE FROM ViewNodeSensorData
			WHERE NodeId = ?`
	} else {
		return `DELETE FROM viewnodesensordata
			WHERE nodeid = $1`
	}
}

func (statements *ViewNodeSensorDataStatements) DeleteAll() string {
	if statements.DbType == "mysql" {
		return `DELETE FROM ViewNodeSensorData`
	} else {
		return `DELETE FROM viewnodesensordata`
	}
}

func (statements *ViewNodeSensorDataStatements) DeleteBySecondParentId() string {
	if statements.DbType == "mysql" {
		return `DELETE FROM ViewNodeSensorData
			WHERE ViewId = ?`
	} else {
		return `DELETE FROM viewnodesensordata
			WHERE viewid = $1`
	}
}
