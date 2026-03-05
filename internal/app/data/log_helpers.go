package data

import (
	"database/sql"
	"time"

	"github.com/ltruelove/gohome/internal/app/models"
)

// Minimal stub implementations to satisfy controller usages. Full
// implementations can replace these later.
func CreateNewLog(item models.NodeData, db *sql.DB) error {
	return nil
}

func GetSensorLogData(nodeId int, db *sql.DB, start time.Time, end time.Time) ([]models.NodeSensorLog, error) {
	return []models.NodeSensorLog{}, nil
}

func GetTempLogDataByLogId(logId int, db *sql.DB) ([]models.TempLogData, error) {
	return []models.TempLogData{}, nil
}

func GetMoistureLogDataByLogId(logId int, db *sql.DB) ([]models.MoistureLogData, error) {
	return []models.MoistureLogData{}, nil
}

func GetResistorLogDataByLogId(logId int, db *sql.DB) ([]models.ResistorLogData, error) {
	return []models.ResistorLogData{}, nil
}

func GetMagneticLogDataByLogId(logId int, db *sql.DB) ([]models.MagneticLogData, error) {
	return []models.MagneticLogData{}, nil
}
