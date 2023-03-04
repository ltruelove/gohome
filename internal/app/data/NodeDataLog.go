package data

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/ltruelove/gohome/internal/app/models"
)

func CreateNewLog(data models.NodeData, db *sql.DB) error {
	stmt, err := db.Prepare(`INSERT INTO nodesensorlog
	(nodeid, datelogged)
	VALUES ($1, $2) RETURNING id`)

	if err != nil {
		log.Println("Error preparing create data log sql")
		return err
	}

	currentTime := time.Now().UTC()

	lastInsertId := 0
	err = stmt.QueryRow(&data.NodeId,
		fmt.Sprintf(currentTime.Format("20060102150405"))).Scan(&lastInsertId)

	if err != nil {
		log.Println("Error creating data log entry: " + err.Error())
		return err
	}

	defer stmt.Close()

	sensors, err := FetchAllNodeSensorsByNode(data.NodeId, db)

	if err != nil {
		log.Println("Error getting sensors by node")
		return err
	}

	for _, sensor := range sensors {
		switch sensor.SensorTypeId {
		case 1:
			err = CreateNewTempLog(int(lastInsertId), data, db)

			if err != nil {
				log.Println("Error creating temperature log data")
				return err
			}
		case 2:
			err = CreateNewMoistureLog(int(lastInsertId), data, db)

			if err != nil {
				log.Println("Error creating moisture log data")
				return err
			}
		case 3:
			err = CreateNewMagneticLog(int(lastInsertId), data, db)

			if err != nil {
				log.Println("Error creating magnetic log data")
				return err
			}
		case 4:
			err = CreateNewResistorLog(int(lastInsertId), data, db)

			if err != nil {
				log.Println("Error creating resistor log data")
				return err
			}

		}
	}

	return nil
}

func CreateNewTempLog(logId int, data models.NodeData, db *sql.DB) error {
	stmt, err := db.Prepare(`INSERT INTO templog
	(nodesensorlogid, temperaturef, temperaturec, humidity)
	VALUES ($1, $2, $3, $4)`)

	if err != nil {
		log.Println("Error preparing create temperature data log sql")
		return err
	}

	_, err = stmt.Exec(&logId, data.TemperatureF, data.TemperatureC, data.Humidity)

	if err != nil {
		log.Println("Error creating temperature log entry")
		return err
	}

	defer stmt.Close()

	return err
}

func CreateNewMoistureLog(logId int, data models.NodeData, db *sql.DB) error {
	stmt, err := db.Prepare(`INSERT INTO moisturelog
	(nodesensorlogid, moisture)
	VALUES ($1, $2)`)

	if err != nil {
		log.Println("Error preparing create moisture data log sql")
		return err
	}

	_, err = stmt.Exec(&logId, data.Moisture)

	if err != nil {
		log.Println("Error creating moisture log entry")
		return err
	}

	defer stmt.Close()

	return err
}

func CreateNewMagneticLog(logId int, data models.NodeData, db *sql.DB) error {
	stmt, err := db.Prepare(`INSERT INTO magneticlog
	(nodesensorlogid, isclosed)
	VALUES ($1, $2)`)

	if err != nil {
		log.Println("Error preparing create magnetic data log sql")
		return err
	}

	isClosed := 0
	if data.MagneticValue {
		isClosed = 1
	}

	_, err = stmt.Exec(&logId, isClosed)

	if err != nil {
		log.Println("Error creating magnetic log entry: " + err.Error())
		return err
	}

	defer stmt.Close()

	return err

}

func CreateNewResistorLog(logId int, data models.NodeData, db *sql.DB) error {
	stmt, err := db.Prepare(`INSERT INTO resistorlog
	(nodesensorlogid, resistorvalue)
	VALUES ($1, $2)`)

	if err != nil {
		log.Println("Error preparing create resistor data log sql")
		return err
	}

	_, err = stmt.Exec(&logId, data.ResistorValue)

	if err != nil {
		log.Println("Error creating resistor log entry")
		return err
	}

	defer stmt.Close()

	return err
}

func GetSensorLogData(nodeId int, db *sql.DB, start time.Time, end time.Time) ([]models.NodeSensorLog, error) {
	stmt, err := db.Prepare(`SELECT
	 	id,
		nodeid,
		datelogged
		FROM nodesensorlog
		WHERE nodeid = $1
		AND datelogged >= $2
		AND datelogged <= $3
		ORDER BY datelogged`)

	if err != nil {
		log.Printf("Error preparing select node sensor data: %v", err)
		return nil, err
	}

	var sensorLogData []models.NodeSensorLog

	rows, err := stmt.Query(nodeId, fmt.Sprintf(start.Format("20060102150405")), fmt.Sprintf(end.Format("20060102150405")))

	if err != nil {
		log.Println("Error querying for node sensor data")
		return nil, err
	}

	defer stmt.Close()

	for rows.Next() {
		var sensorLog models.NodeSensorLog
		var dateLoggedString string

		err = rows.Scan(&sensorLog.Id,
			&sensorLog.NodeId,
			&dateLoggedString)

		if err != nil {
			log.Println("Error scanning node switch")
			return nil, err
		}

		sensorLog.DateLogged, err = time.Parse("20060102150405", dateLoggedString)

		if err != nil {
			log.Println("Error parsing log date")
			return nil, err
		}

		sensorLogData = append(sensorLogData, sensorLog)
	}

	return sensorLogData, nil
}

func GetTempLogDataByLogId(nodeSensorLogId int, db *sql.DB) ([]models.TempLogData, error) {
	stmt, err := db.Prepare(`SELECT
	 	id,
		nodesensorlogid,
		temperaturef,
		temperaturec,
		humidity
		FROM templog
		WHERE nodesensorlogid = $1`)

	if err != nil {
		log.Printf("Error preparing select temperature data: %v", err)
		return nil, err
	}

	rows, err := stmt.Query(nodeSensorLogId)

	if err != nil {
		log.Println("Error querying for temperature log data")
		return nil, err
	}

	defer stmt.Close()

	var logData []models.TempLogData
	var i int = 0

	for rows.Next() {
		i++
		var logRecord models.TempLogData

		err = rows.Scan(&logRecord.Id,
			&logRecord.NodeSensorLogId,
			&logRecord.TemperatureF,
			&logRecord.TemperatureC,
			&logRecord.Humidity)

		if err != nil {
			log.Println("Error scanning temperature log record")
			return nil, err
		}

		logData = append(logData, logRecord)
	}

	return logData, nil
}

func GetMagneticLogDataByLogId(nodeSensorLogId int, db *sql.DB) ([]models.MagneticLogData, error) {
	stmt, err := db.Prepare(`SELECT
	 	id,
		nodesensorlogid,
		isclosed
		FROM magneticlog
		WHERE nodesensorlogid = $1`)

	if err != nil {
		log.Printf("Error preparing select magnetic data: %v", err)
		return nil, err
	}

	rows, err := stmt.Query(nodeSensorLogId)

	if err != nil {
		log.Println("Error querying for magnetic log data")
		return nil, err
	}

	defer stmt.Close()

	var logData []models.MagneticLogData

	for rows.Next() {
		var logRecord models.MagneticLogData

		err = rows.Scan(&logRecord.Id,
			&logRecord.NodeSensorLogId,
			&logRecord.IsClosed)

		if err != nil {
			log.Println("Error scanning magnetic log record")
			return nil, err
		}

		logData = append(logData, logRecord)
	}

	return logData, nil
}

func GetResistorLogDataByLogId(nodeSensorLogId int, db *sql.DB) ([]models.ResistorLogData, error) {
	stmt, err := db.Prepare(`SELECT
	 	id,
		nodesensorlogid,
		resistorvalue
		FROM resistorlog
		WHERE nodesensorlogid = $1`)

	if err != nil {
		log.Printf("Error preparing select resistor data: %v", err)
		return nil, err
	}

	rows, err := stmt.Query(nodeSensorLogId)

	if err != nil {
		log.Println("Error querying for resistor log data")
		return nil, err
	}

	defer stmt.Close()

	var logData []models.ResistorLogData

	for rows.Next() {
		var logRecord models.ResistorLogData

		err = rows.Scan(&logRecord.Id,
			&logRecord.NodeSensorLogId,
			&logRecord.ResistorValue)

		if err != nil {
			log.Println("Error scanning resistor log record")
			return nil, err
		}

		logData = append(logData, logRecord)
	}

	return logData, nil
}

func GetMoistureLogDataByLogId(nodeSensorLogId int, db *sql.DB) ([]models.MoistureLogData, error) {
	stmt, err := db.Prepare(`SELECT
	 	id,
		nodesensorlogid,
		moisture
		FROM moisturelog
		WHERE nodesensorlogid = $1`)

	if err != nil {
		log.Printf("Error preparing select moisture data: %v", err)
		return nil, err
	}

	rows, err := stmt.Query(nodeSensorLogId)

	if err != nil {
		log.Println("Error querying for moisture log data")
		return nil, err
	}

	defer stmt.Close()

	var logData []models.MoistureLogData

	for rows.Next() {
		var logRecord models.MoistureLogData

		err = rows.Scan(&logRecord.Id,
			&logRecord.NodeSensorLogId,
			&logRecord.Moisture)

		if err != nil {
			log.Println("Error scanning moisture log record")
			return nil, err
		}

		logData = append(logData, logRecord)
	}

	return logData, nil
}
