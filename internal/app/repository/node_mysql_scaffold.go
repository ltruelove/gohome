package repository

import (
	"database/sql"
	"log"
	"time"

	"github.com/ltruelove/gohome/internal/app/data"
	"github.com/ltruelove/gohome/internal/app/models"
)

type mysqlNodeRepository struct {
	db *sql.DB
}

func NewMySQLNodeRepository(db *sql.DB) NodeRepository {
	return &mysqlNodeRepository{db: db}
}

func (r *mysqlNodeRepository) FetchAll() ([]models.Node, error) {
	rows, err := r.db.Query("SELECT id, mac, name, ipaddress FROM node")
	if err != nil {
		log.Println("Error querying FetchAll nodes:", err)
		return nil, err
	}
	defer rows.Close()

	var nodes []models.Node
	for rows.Next() {
		var n models.Node
		if err := rows.Scan(&n.Id, &n.Mac, &n.Name, &n.IpAddress); err != nil {
			return nil, err
		}
		nodes = append(nodes, n)
	}
	return nodes, nil
}

func (r *mysqlNodeRepository) FetchById(id int) (models.Node, error) {
	var n models.Node
	err := r.db.QueryRow("SELECT id, mac, name, ipaddress FROM node WHERE id = ?", id).Scan(&n.Id, &n.Mac, &n.Name, &n.IpAddress)
	if err != nil {
		return n, err
	}
	return n, nil
}

func (r *mysqlNodeRepository) Create(node *models.Node) error {
	res, err := r.db.Exec("INSERT INTO node (mac, name, ipaddress) VALUES (?, ?, ?)", node.Mac, node.Name, node.IpAddress)
	if err != nil {
		log.Println("Error inserting node:", err)
		return err
	}
	id, err := res.LastInsertId()
	if err == nil {
		node.Id = int(id)
	}
	return nil
}

func (r *mysqlNodeRepository) VerifyIdIsNew(id int) (bool, error) {
	_, err := r.FetchById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func (r *mysqlNodeRepository) Update(node *models.Node) error {
	_, err := r.db.Exec("UPDATE node SET mac = ?, name = ?, ipaddress = ? WHERE id = ?", node.Mac, node.Name, node.IpAddress, node.Id)
	return err
}

func (r *mysqlNodeRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM node WHERE id = ?", id)
	return err
}

func (r *mysqlNodeRepository) FetchIndividual(id int) (models.Node, error) {
	return r.FetchById(id)
}

func (r *mysqlNodeRepository) FetchControlPointByNode(nodeId int) (models.ControlPoint, error) {
	var cp models.ControlPoint
	err := r.db.QueryRow(`SELECT cp.id, cp.name, cp.ipaddress, cp.mac FROM controlpointnodes AS cpn INNER JOIN controlpoint AS cp ON cp.id = cpn.controlpointid WHERE cpn.nodeid = ?`, nodeId).Scan(&cp.Id, &cp.Name, &cp.IpAddress, &cp.Mac)
	if err != nil {
		return cp, err
	}
	return cp, nil
}

func (r *mysqlNodeRepository) FetchNodeSwitches(nodeId int) ([]models.NodeSwitch, error) {
	rows, err := r.db.Query("SELECT id, nodeid, switchtypeid, name, pin, momentarypressduration, isclosedon FROM nodeswitch WHERE nodeid = ?", nodeId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.NodeSwitch
	for rows.Next() {
		var ns models.NodeSwitch
		if err := rows.Scan(&ns.Id, &ns.NodeId, &ns.SwitchTypeId, &ns.Name, &ns.Pin, &ns.MomentaryPressDuration, &ns.IsClosedOn); err != nil {
			return nil, err
		}
		list = append(list, ns)
	}
	return list, nil
}

func (r *mysqlNodeRepository) FetchNodeSwitch(id int) (models.NodeSwitch, error) {
	var ns models.NodeSwitch
	err := r.db.QueryRow("SELECT id, nodeid, switchtypeid, name, pin, momentarypressduration, isclosedon FROM nodeswitch WHERE id = ?", id).Scan(&ns.Id, &ns.NodeId, &ns.SwitchTypeId, &ns.Name, &ns.Pin, &ns.MomentaryPressDuration, &ns.IsClosedOn)
	if err != nil {
		return ns, err
	}
	return ns, nil
}

func (r *mysqlNodeRepository) GetSensorLogData(nodeId int, start time.Time, end time.Time) ([]models.NodeSensorLog, error) {
	// delegate to existing data helper (currently returns empty stub)
	return []models.NodeSensorLog{}, nil
}

func (r *mysqlNodeRepository) GetTempLogDataByLogId(logId int) ([]models.TempLogData, error) {
	return data.GetTempLogDataByLogId(logId, r.db)
}

func (r *mysqlNodeRepository) CreateNewLog(item models.NodeData) error {
	return nil
}

func (r *mysqlNodeRepository) GetMoistureLogDataByLogId(logId int) ([]models.MoistureLogData, error) {
	return data.GetMoistureLogDataByLogId(logId, r.db)
}

func (r *mysqlNodeRepository) GetResistorLogDataByLogId(logId int) ([]models.ResistorLogData, error) {
	return data.GetResistorLogDataByLogId(logId, r.db)
}

func (r *mysqlNodeRepository) GetMagneticLogDataByLogId(logId int) ([]models.MagneticLogData, error) {
	return data.GetMagneticLogDataByLogId(logId, r.db)
}

func (r *mysqlNodeRepository) CreateNodeSensor(sensor *models.NodeSensor) error {
	res, err := r.db.Exec("INSERT INTO nodesensor (nodeid, sensortypeid, name, pin, dhttype) VALUES (?, ?, ?, ?, ?)", sensor.NodeId, sensor.SensorTypeId, sensor.Name, sensor.Pin, sensor.DHTType)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err == nil {
		sensor.Id = int(id)
	}
	return nil
}

func (r *mysqlNodeRepository) CreateNodeSwitch(sw *models.NodeSwitch) error {
	isClosedOnInt := 0
	if sw.IsClosedOn {
		isClosedOnInt = 1
	}
	res, err := r.db.Exec("INSERT INTO nodeswitch (nodeid, switchtypeid, name, pin, momentarypressduration, isclosedon) VALUES (?, ?, ?, ?, ?, ?)", sw.NodeId, sw.SwitchTypeId, sw.Name, sw.Pin, sw.MomentaryPressDuration, isClosedOnInt)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err == nil {
		sw.Id = int(id)
	}
	return nil
}
