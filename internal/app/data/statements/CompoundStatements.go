package statements

import "github.com/ltruelove/gohome/config"

type CompoundStatements struct {
	DbType string
}

func NewCompoundStatements(config *config.Configuration) *CompoundStatements {
	return &CompoundStatements{DbType: config.DbType}
}

func (statements *CompoundStatements) SelectViewNodeSensorDataByViewId() string {
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

func (statements *CompoundStatements) SelectViewNodeSwitchDataByViewId() string {
	if statements.DbType == "mysql" {
		return `SELECT
		d.Id,
		d.NodeId,
		d.ViewId,
		d.NodeSwitchId,
		d.Name,
		n.Name AS NodeName,
		ns.Name AS SwitchName,
		st.Name AS SwitchTypeName
		FROM ViewNodeSwitchData AS d
		INNER JOIN Node AS n ON n.Id = d.NodeId
		INNER JOIN NodeSwitch AS ns ON ns.Id = d.NodeSwitchId
		INNER JOIN SwitchType AS st ON st.Id = ns.SwitchTypeId
		WHERE d.ViewId = ?`
	} else {
		return `SELECT
		d.id,
		d.nodeid,
		d.viewid,
		d.nodeswitchid,
		d.name,
		n.name AS nodename,
		ns.name AS switchname,
		st.name AS switchtypename
		FROM viewnodeswitchdata AS d
		INNER JOIN node AS n ON n.id = d.nodeid
		INNER JOIN nodeswitch AS ns ON ns.id = d.nodeswitchid
		INNER JOIN switchtype AS st ON st.id = ns.switchtypeid
		WHERE d.viewid = $1`
	}
}
