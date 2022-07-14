package data

import (
	"database/sql"
	"log"

	"github.com/ltruelove/gohome/internal/app/models"
	"github.com/ltruelove/gohome/internal/app/setup"
)

func VerifyViewIdIsNew(viewId int, db *sql.DB) (bool, error) {
	view, err := FetchView(viewId, db)

	if err != nil {
		log.Println("Error fetching view")
		return false, err
	}

	return view.Id > 0, nil
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
	stmt, err := db.Prepare("SELECT Id, Name FROM View WHERE id = ?")
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

func FetchViewSensorData(viewId int, db *sql.DB) ([]models.ViewNodeSensorData, error) {
	stmt, err := db.Prepare(`SELECT
		Id,
		NodeId,
		ViewId,
		NodeSensorId,
		SensorTypeDataId,
		Name FROM ViewNodeSensorData WHERE ViewId = ?`)

	if err != nil {
		log.Println("Error preparing fetch view sql")
		return nil, err
	}

	var sensorDataList []models.ViewNodeSensorData

	rows, err := stmt.Query()

	if err != nil {
		log.Println("Error querying for view")
		return nil, err
	}

	defer stmt.Close()

	for rows.Next() {
		var sensorDataItem models.ViewNodeSensorData

		err := rows.Scan(&sensorDataItem.Id,
			&sensorDataItem.NodeId,
			&sensorDataItem.ViewId,
			&sensorDataItem.NodeSensorId,
			&sensorDataItem.SensorTypeDataId,
			&sensorDataItem.Name)

		if err != nil {
			log.Println("Error scanning view")
			return nil, err
		}

		sensorDataList = append(sensorDataList, sensorDataItem)
	}

	return sensorDataList, nil
}

func FetchViewSwitchData(viewId int, db *sql.DB) ([]models.ViewNodeSwitchData, error) {
	stmt, err := db.Prepare(`SELECT
		Id,
		NodeId,
		ViewId,
		NodeSwitchId,
		Name FROM ViewNodeSwitchData WHERE ViewId = ?`)

	if err != nil {
		log.Println("Error preparing fetch view sql")
		return nil, err
	}

	var switchDataList []models.ViewNodeSwitchData

	rows, err := stmt.Query()

	if err != nil {
		log.Println("Error querying for view")
		return nil, err
	}

	defer stmt.Close()
	setup.CheckErr(err)

	for rows.Next() {
		var switchDataItem models.ViewNodeSwitchData

		err := rows.Scan(&switchDataItem.Id,
			&switchDataItem.NodeId,
			&switchDataItem.ViewId,
			&switchDataItem.NodeSwitchId,
			&switchDataItem.Name)

		if err != nil {
			log.Println("Error scanning view")
			return nil, err
		}

		switchDataList = append(switchDataList, switchDataItem)
	}

	return switchDataList, nil
}

func CreateView(view *models.View, db *sql.DB) error {
	stmt, err := db.Prepare(`INSERT INTO View
	(Id, Name)
	VALUES (?, ?)`)

	if err != nil {
		log.Println("Error preparing create view sql")
		return err
	}

	_, err = stmt.Exec(view.Id,
		view.Name)

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
