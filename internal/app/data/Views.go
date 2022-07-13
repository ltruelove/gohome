package data

import (
	"database/sql"

	"github.com/ltruelove/gohome/internal/app/models"
	"github.com/ltruelove/gohome/internal/app/setup"
)

func VerifyViewIdIsNew(viewId int, db *sql.DB) bool {
	view := FetchView(viewId, db)
	return view.Id > 0
}

func FetchAllViews(db *sql.DB) []models.View {
	stmt, err := db.Prepare(`SELECT Id, Name FROM View`)

	setup.CheckErr(err)
	var views []models.View

	rows, err := stmt.Query()
	setup.CheckErr(err)

	for rows.Next() {
		var view models.View
		rows.Scan(&view.Id,
			&view.Name)
		views = append(views, view)
	}
	defer stmt.Close()

	return views
}

func FetchView(viewId int, db *sql.DB) models.View {
	stmt, err := db.Prepare("SELECT Id, Name FROM View WHERE id = ?")
	setup.CheckErr(err)
	defer stmt.Close()

	var view models.View

	err = stmt.QueryRow(viewId).Scan(&view.Id,
		&view.Name)

	if err != sql.ErrNoRows {
		setup.CheckErr(err)
	}

	return view
}

func FetchViewSensorData(viewId int, db *sql.DB) []models.ViewNodeSensorData {
	stmt, err := db.Prepare(`SELECT
		Id,
		NodeId,
		ViewId,
		NodeSensorId,
		SensorTypeDataId,
		Name FROM ViewNodeSensorData WHERE ViewId = ?`)
	setup.CheckErr(err)
	defer stmt.Close()

	var sensorDataList []models.ViewNodeSensorData

	rows, err := stmt.Query()
	setup.CheckErr(err)

	for rows.Next() {
		var sensorDataItem models.ViewNodeSensorData

		rows.Scan(&sensorDataItem.Id,
			&sensorDataItem.NodeId,
			&sensorDataItem.ViewId,
			&sensorDataItem.NodeSensorId,
			&sensorDataItem.SensorTypeDataId,
			&sensorDataItem.Name)

		sensorDataList = append(sensorDataList, sensorDataItem)
	}

	return sensorDataList
}

func FetchViewSwitchData(viewId int, db *sql.DB) []models.ViewNodeSwitchData {
	stmt, err := db.Prepare(`SELECT
		Id,
		NodeId,
		ViewId,
		NodeSwitchId,
		Name FROM ViewNodeSwitchData WHERE ViewId = ?`)
	setup.CheckErr(err)
	defer stmt.Close()

	var switchDataList []models.ViewNodeSwitchData

	rows, err := stmt.Query()
	setup.CheckErr(err)

	for rows.Next() {
		var switchDataItem models.ViewNodeSwitchData

		rows.Scan(&switchDataItem.Id,
			&switchDataItem.NodeId,
			&switchDataItem.ViewId,
			&switchDataItem.NodeSwitchId,
			&switchDataItem.Name)

		switchDataList = append(switchDataList, switchDataItem)
	}

	return switchDataList
}

func CreateView(view *models.View, db *sql.DB) {
	stmt, err := db.Prepare(`INSERT INTO View
	(Id, Name)
	VALUES (?, ?)`)

	setup.CheckErr(err)

	_, err = stmt.Exec(view.Id,
		view.Name)

	defer stmt.Close()

	setup.CheckErr(err)
}

func UpdateView(view *models.View, db *sql.DB) {
	stmt, err := db.Prepare(`UPDATE View
	SET Name = ?
	WHERE id = ?`)

	setup.CheckErr(err)

	_, err = stmt.Exec(view.Name,
		view.Id)

	defer stmt.Close()

	setup.CheckErr(err)
}

func DeleteView(viewId int, db *sql.DB) {
	stmt, err := db.Prepare(`DELETE FROM View
	WHERE id = ?`)

	setup.CheckErr(err)

	_, err = stmt.Exec(viewId)

	defer stmt.Close()

	setup.CheckErr(err)
}
