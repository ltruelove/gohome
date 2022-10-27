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
	stmt, err := db.Prepare(`SELECT Id, Name FROM View`)

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
	stmt, err := db.Prepare("SELECT Id, Name FROM View WHERE Id = ?")

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
		WHERE d.ViewId = ?`)

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
		INNER JOIN NodeSwitch AS ns on ns.Id = d.NodeSwitchId
		INNER JOIN SwitchType AS st on st.Id = ns.SwitchTypeId
		WHERE d.ViewId = ?`)

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
	stmt, err := db.Prepare(`INSERT INTO View
	(Name)
	VALUES (?)`)

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
	stmt, err := db.Prepare(`UPDATE View
	SET Name = ?
	WHERE id = ?`)

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
	stmt, err := db.Prepare(`DELETE FROM View
	WHERE id = ?`)

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
