package data

import (
	"database/sql"
	"log"

	"github.com/ltruelove/gohome/internal/app/models"
	"github.com/ltruelove/gohome/internal/app/setup"
	"github.com/ltruelove/gohome/internal/app/viewModels"
)

func VerifyViewIdIsNew(viewId int, db *sql.DB) (bool, error) {
	view, err := FetchView(viewId, db)

	log.Printf("view found with id: %d, and name: %s", view.Id, view.Name)

	if err != nil {
		log.Println("Error fetching view")
		return false, err
	}

	return view.Id < 1, nil
}

func FetchAllViews(db *sql.DB) ([]models.View, error) {
	stmt, err := db.Prepare(`SELECT id, name FROM view`)

	if err != nil {
		log.Println("Error preparing all views sql")
		return nil, err
	}

	var views []models.View

	rows, err := stmt.Query()
	if err != nil {
		log.Println("Error querying for all views")
		return nil, err
	}

	defer stmt.Close()

	for rows.Next() {
		var view models.View
		err := rows.Scan(&view.Id,
			&view.Name)

		if err != nil {
			log.Println("Error scanning view")
			return nil, err
		}

		views = append(views, view)
	}

	return views, nil
}

func FetchView(viewId int, db *sql.DB) (models.View, error) {
	var item models.View
	log.Printf("Fetching view for id: %d", viewId)
	stmt, err := db.Prepare("SELECT id, name FROM view WHERE id = $1")

	if err != nil {
		log.Println("Error preparing fetch view sql")
		return item, err
	}

	err = stmt.QueryRow(viewId).Scan(&item.Id,
		&item.Name)

	if err != nil {
		log.Println("Error querying for view")
		return item, err
	}

	defer stmt.Close()

	return item, nil
}

func FetchViewSensorData(viewId int, db *sql.DB) ([]viewModels.ViewNodeSensorVM, error) {
	stmt, err := db.Prepare(`SELECT
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
		WHERE d.viewid = $1`)

	if err != nil {
		log.Println("Error preparing fetch view sql")
		return nil, err
	}

	var sensorDataList []viewModels.ViewNodeSensorVM

	rows, err := stmt.Query(viewId)

	if err != nil {
		log.Println("Error querying for view")
		return nil, err
	}

	defer stmt.Close()

	for rows.Next() {
		var sensorDataItem viewModels.ViewNodeSensorVM

		err := rows.Scan(&sensorDataItem.Id,
			&sensorDataItem.NodeId,
			&sensorDataItem.ViewId,
			&sensorDataItem.NodeSensorId,
			&sensorDataItem.Name,
			&sensorDataItem.NodeName,
			&sensorDataItem.SensorName,
			&sensorDataItem.SensorTypeName)

		if err != nil {
			log.Println("Error scanning view")
			return nil, err
		}

		sensorDataList = append(sensorDataList, sensorDataItem)
	}

	return sensorDataList, nil
}

func FetchViewSwitchData(viewId int, db *sql.DB) ([]viewModels.ViewNodeSwitchVM, error) {
	log.Printf("Finding all switches for view with id %d", viewId)

	stmt, err := db.Prepare(`SELECT
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
		WHERE d.viewid = $1`)

	if err != nil {
		log.Println("Error preparing fetch view switch data sql")
		return nil, err
	}

	var switchDataList []viewModels.ViewNodeSwitchVM

	rows, err := stmt.Query(viewId)

	if err != nil {
		log.Println("Error querying for view")
		return nil, err
	}

	defer stmt.Close()
	setup.CheckErr(err)

	for rows.Next() {
		var switchDataItem viewModels.ViewNodeSwitchVM

		err := rows.Scan(&switchDataItem.Id,
			&switchDataItem.NodeId,
			&switchDataItem.ViewId,
			&switchDataItem.NodeSwitchId,
			&switchDataItem.Name,
			&switchDataItem.NodeName,
			&switchDataItem.SwitchName,
			&switchDataItem.SwitchTypeName)

		if err != nil {
			log.Println("Error scanning view")
			return nil, err
		}

		switchDataList = append(switchDataList, switchDataItem)
	}

	log.Printf("Found %d switches", len(switchDataList))
	return switchDataList, nil
}

func CreateView(view *models.View, db *sql.DB) error {
	stmt, err := db.Prepare(`INSERT INTO view
	(name)
	VALUES ($1)`)

	if err != nil {
		log.Println("Error preparing create view sql")
		return err
	}

	_, err = stmt.Exec(view.Name)

	if err != nil {
		log.Println("Error creating view")
		return err
	}

	defer stmt.Close()

	return nil
}

func UpdateView(view *models.View, db *sql.DB) error {
	stmt, err := db.Prepare(`UPDATE view
	SET name = $1
	WHERE id = $2`)

	if err != nil {
		log.Println("Error preparing update view sql")
		return err
	}

	_, err = stmt.Exec(view.Name,
		view.Id)

	if err != nil {
		log.Println("Error updating view")
		return err
	}

	defer stmt.Close()

	return nil
}

func DeleteView(viewId int, db *sql.DB) error {
	stmt, err := db.Prepare(`DELETE FROM view
	WHERE id = $1`)

	if err != nil {
		log.Println("Error preparing delete view sql")
		return err
	}

	_, err = stmt.Exec(viewId)

	if err != nil {
		log.Println("Error deleting view")
		return err
	}

	defer stmt.Close()

	return nil
}
