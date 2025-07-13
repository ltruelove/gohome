package statements

import "github.com/ltruelove/gohome/config"

type NodeSwitchDataStatements struct {
	DbType string
}

func NewNodeSwitchDataStatements(config *config.Configuration) CrudStatement {
	return &NodeSwitchDataStatements{DbType: config.DbType}
}

func (statements *NodeSwitchDataStatements) SelectAll() string {
	if statements.DbType == "mysql" {
		return `SELECT
		ns.Id,
		ns.NodeId,
		ns.SwitchTypeId,
		ns.SwitchName,
		ns.Pin,
		ns.MomentaryPressDuration,
		ns.IsClosedOn,
		st.Name AS SwitchTypeName
		FROM NodeSwitch AS ns
		INNER JOIN SwitchType AS st ON st.Id = ns.SwitchTypeId`
	} else {
		return `SELECT
		ns.id,
		ns.nodeid,
		ns.switchtypeid,
		ns.name,
		ns.pin,
		ns.momentarypressduration,
		ns.isclosedon,
		st.name AS switchtypename
		FROM nodeswitch AS ns
		INNER JOIN switchtype AS st ON st.id = ns.switchtypeid`
	}
}

func (statements *NodeSwitchDataStatements) SelectById() string {
	if statements.DbType == "mysql" {
		return `SELECT
		ns.Id,
		ns.NodeId,
		ns.SwitchTypeId,
		ns.SwitchName,
		ns.Pin,
		ns.MomentaryPressDuration,
		ns.IsClosedOn,
		st.Name AS SwitchTypeName
		FROM NodeSwitch AS ns
		INNER JOIN SwitchType AS st ON st.Id = ns.SwitchTypeId
		WHERE ns.Id = ?`
	} else {
		return `SELECT
		ns.id,
		ns.nodeid,
		ns.switchtypeid,
		ns.name,
		ns.pin,
		ns.momentarypressduration,
		ns.isclosedon,
		st.name AS switchtypename
		FROM nodeswitch AS ns
		INNER JOIN switchtype AS st ON st.id = ns.switchtypeid
		WHERE ns.id = $1`
	}
}

// SelectByParentId retrieves NodeSwitch records by NodeId
func (statements *NodeSwitchDataStatements) SelectByParentId() string {
	if statements.DbType == "mysql" {
		return `SELECT
		ns.Id,
		ns.NodeId,
		ns.SwitchTypeId,
		ns.SwitchName,
		ns.Pin,
		ns.MomentaryPressDuration,
		ns.IsClosedOn,
		st.Name AS SwitchTypeName
		FROM NodeSwitch AS ns
		INNER JOIN SwitchType AS st ON st.Id = ns.SwitchTypeId
		WHERE ns.NodeId = ?`
	} else {
		return `SELECT
		ns.id,
		ns.nodeid,
		ns.switchtypeid,
		ns.name,
		ns.pin,
		ns.momentarypressduration,
		ns.isclosedon,
		st.name AS switchtypename
		FROM nodeswitch AS ns
		INNER JOIN switchtype AS st ON st.id = ns.switchtypeid
		WHERE ns.nodeid = $1`
	}
}

// SelectBySecondParentId retrieves NodeSwitch records by SwitchTypeId
func (statements *NodeSwitchDataStatements) SelectBySecondParentId() string {
	if statements.DbType == "mysql" {
		return `SELECT
		ns.Id,
		ns.NodeId,
		ns.SwitchTypeId,
		ns.SwitchName,
		ns.Pin,
		ns.MomentaryPressDuration,
		ns.IsClosedOn,
		st.Name AS SwitchTypeName
		FROM NodeSwitch AS ns
		INNER JOIN SwitchType AS st ON st.Id = ns.SwitchTypeId
		WHERE ns.SwitchTypeId = ?`
	} else {
		return `SELECT
		ns.id,
		ns.nodeid,
		ns.switchtypeid,
		ns.name,
		ns.pin,
		ns.momentarypressduration,
		ns.isclosedon,
		st.name AS switchtypename
		FROM nodeswitch AS ns
		INNER JOIN switchtype AS st ON st.id = ns.switchtypeid
		WHERE ns.switchtypeid = $1`
	}
}

func (statements *NodeSwitchDataStatements) Insert() string {
	if statements.DbType == "mysql" {
		return `INSERT INTO NodeSwitch (
			NodeId,
			SwitchTypeId,
			SwitchName,
			Pin,
			MomentaryPressDuration,
			IsClosedOn
		) VALUES (?, ?, ?, ?, ?, ?)`
	} else {
		return `INSERT INTO nodeswitch (
			nodeid,
			switchtypeid,
			name,
			pin,
			momentarypressduration,
			isclosedon
		) VALUES ($1, $2, $3, $4, $5, $6)`
	}
}

func (statements *NodeSwitchDataStatements) Update() string {
	if statements.DbType == "mysql" {
		return `UPDATE NodeSwitch
			SET NodeId = ?, SwitchTypeId = ?, SwitchName = ?, Pin = ?, MomentaryPressDuration = ?, IsClosedOn = ?
			WHERE Id = ?`
	} else {
		return `UPDATE nodeswitch
			SET nodeid = $1, switchtypeid = $2, name = $3, pin = $4, momentarypressduration = $5, isclosedon = $6
			WHERE id = $7`
	}
}

func (statements *NodeSwitchDataStatements) Delete() string {
	if statements.DbType == "mysql" {
		return `DELETE FROM NodeSwitch WHERE Id = ?`
	} else {
		return `DELETE FROM nodeswitch WHERE id = $1`
	}
}

func (statements *NodeSwitchDataStatements) DeleteAll() string {
	if statements.DbType == "mysql" {
		return `DELETE FROM NodeSwitch`
	} else {
		return `DELETE FROM nodeswitch`
	}
}

func (statements *NodeSwitchDataStatements) DeleteByParentId() string {
	if statements.DbType == "mysql" {
		return `DELETE FROM NodeSwitch WHERE NodeId = ?`
	} else {
		return `DELETE FROM nodeswitch WHERE nodeid = $1`
	}
}
