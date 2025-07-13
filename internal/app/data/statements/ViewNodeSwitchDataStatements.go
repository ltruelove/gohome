package statements

import "github.com/ltruelove/gohome/config"

type ViewNodeSwitchDataStatements struct {
	DbType string
}

func NewViewNodeSwitchDataStatements(config *config.Configuration) CrudStatement {
	return &ViewNodeSwitchDataStatements{DbType: config.DbType}
}

func (statements *ViewNodeSwitchDataStatements) SelectAll() string {
	if statements.DbType == "mysql" {
		return `SELECT
			Id,
			NodeId,
			ViewId,
			NodeSwitchId,
			Name
			FROM ViewNodeSwitchData`

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

func (statements *ViewNodeSwitchDataStatements) SelectById() string {
	if statements.DbType == "mysql" {
		return `SELECT
			Id,
			NodeId,
			ViewId,
			NodeSwitchId,
			Name
			FROM ViewNodeSwitchData
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

func (statements *ViewNodeSwitchDataStatements) SelectByParentId() string {
	if statements.DbType == "mysql" {
		return `SELECT
			Id,
			NodeId,
			ViewId,
			NodeSwitchId,
			Name
			FROM ViewNodeSwitchData
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

func (statements *ViewNodeSwitchDataStatements) SelectBySecondParentId() string {
	if statements.DbType == "mysql" {
		return `SELECT
		d.Id,
		d.NodeId,
		d.ViewId,
		d.NodeSwitchId,
		d.Name,
		n.Name As NodeName,
		ns.Name AS SwitchName,
		st.Name AS SwitchTypeName
		FROM ViewNodeSwitchData AS d
		INNER JOIN Node AS n ON n.Id = d.NodeId
		INNER JOIN NodeSwitch AS ns ON ns.Id = d.NodeSwitchId
		INNER JOIN SwitchType AS st on st.Id = ns.SwitchTypeId
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

func (statements *ViewNodeSwitchDataStatements) Insert() string {
	if statements.DbType == "mysql" {
		return `INSERT INTO ViewNodeSwitchData
			(NodeId, ViewId, NodeSwitchId, Name)
			VALUES (?, ?, ?, ?); SELECT LAST_INSERT_ID();`
	} else {
		return `INSERT INTO viewnodesensordata
			(nodeid, viewid, nodesensorid, name)
			VALUES ($1, $2, $3, $4) RETURNING id`
	}
}

func (statements *ViewNodeSwitchDataStatements) Update() string {
	if statements.DbType == "mysql" {
		return `UPDATE ViewNodeSwitchData
			SET NodeId = ?, ViewId = ?, NodeSwitchId = ?, Name = ?
			WHERE Id = ?`
	} else {
		return `UPDATE viewnodesensordata
			SET nodeid = $1, viewid = $2, nodesensorid = $3, name = $4
			WHERE id = $5`
	}
}

func (statements *ViewNodeSwitchDataStatements) Delete() string {
	if statements.DbType == "mysql" {
		return `DELETE FROM ViewNodeSwitchData
			WHERE Id = ?`
	} else {
		return `DELETE FROM viewnodesensordata
			WHERE id = $1`
	}
}

func (statements *ViewNodeSwitchDataStatements) DeleteByParentId() string {
	if statements.DbType == "mysql" {
		return `DELETE FROM ViewNodeSwitchData
			WHERE NodeId = ?`
	} else {
		return `DELETE FROM viewnodesensordata
			WHERE nodeid = $1`
	}
}

func (statements *ViewNodeSwitchDataStatements) DeleteAll() string {
	if statements.DbType == "mysql" {
		return `DELETE FROM ViewNodeSwitchData`
	} else {
		return `DELETE FROM viewnodesensordata`
	}
}

func (statements *ViewNodeSwitchDataStatements) DeleteBySecondParentId() string {
	if statements.DbType == "mysql" {
		return `DELETE FROM ViewNodeSwitchData
			WHERE ViewId = ?`
	} else {
		return `DELETE FROM viewnodesensordata
			WHERE viewid = $1`
	}
}
