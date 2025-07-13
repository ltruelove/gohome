/*
MySQL table:

CREATE TABLE IF NOT EXISTS Node (
	Id	INT NOT NULL UNIQUE AUTO_INCREMENT,
	Mac	CHAR(12) UNIQUE,
	NodeName	VARCHAR(255) UNIQUE,
	IpAddress VARCHAR(15) NOT NULL,
	PRIMARY KEY(Id)
);

postgresql table:

CREATE TABLE IF NOT EXISTS "Node" (
	"Id"	INTEGER NOT NULL UNIQUE,
	"Mac"	TEXT UNIQUE,
	"Name"	TEXT UNIQUE,
	PRIMARY KEY("Id" AUTOINCREMENT)
);
*/

package statements

import "github.com/ltruelove/gohome/config"

type NodeDataStatements struct {
	DbType string
}

func NewNodeDataStatements(config *config.Configuration) CrudStatement {
	return &NodeDataStatements{DbType: config.DbType}
}

func (statements *NodeDataStatements) SelectAll() string {
	if statements.DbType == "mysql" {
		return `SELECT
		Id,
		Mac,
		NodeName,
		IpAddress
		FROM Node`
	} else {
		return `SELECT
		id,
		mac,
		name,
		ipaddress
		FROM node`
	}
}

func (statements *NodeDataStatements) SelectById() string {
	if statements.DbType == "mysql" {
		return `SELECT
		Id,
		Mac,
		NodeName,
		IpAddress
		FROM Node
		WHERE Id = ?`
	} else {
		return `SELECT
		id,
		mac,
		name,
		ipaddress
		FROM node
		WHERE id = $1`
	}
}

func (statements *NodeDataStatements) SelectByParentId() string {
	return "Node table has no parent ID"
}

func (statements *NodeDataStatements) SelectBySecondParentId() string {
	return "Node table has no second parent ID"
}

func (statements *NodeDataStatements) Insert() string {
	if statements.DbType == "mysql" {
		return `INSERT INTO Node (Mac, NodeName, IpAddress) VALUES (?, ?, ?); SELECT LAST_INSERT_ID();`
	} else {
		return `INSERT INTO node (mac, name, ipaddress) VALUES ($1, $2, $3) RETURNING id`
	}
}

func (statements *NodeDataStatements) Update() string {
	if statements.DbType == "mysql" {
		return `UPDATE Node SET Mac = ?, NodeName = ?, IpAddress = ? WHERE Id = ?`
	} else {
		return `UPDATE node SET mac = $1, name = $2, ipaddress = $3 WHERE id = $4`
	}
}

func (statements *NodeDataStatements) Delete() string {
	if statements.DbType == "mysql" {
		return `DELETE FROM Node WHERE Id = ?`
	} else {
		return `DELETE FROM node WHERE id = $1`
	}
}

func (statements *NodeDataStatements) DeleteAll() string {
	if statements.DbType == "mysql" {
		return `DELETE FROM Node`
	} else {
		return `DELETE FROM node`
	}
}

func (statements *NodeDataStatements) DeleteByParentId() string {
	return "Node table has no parent ID"
}

func (statements *NodeDataStatements) DeleteBySecondParentId() string {
	return "Node table has no second parent ID"
}
