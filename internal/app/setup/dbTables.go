package setup

const CreateTemperatureSensors string = `CREATE TABLE IF NOT EXISTS "TemperatureSensors" (
	"id"	TEXT NOT NULL UNIQUE,
	"name"	TEXT NOT NULL,
	"isGarage"	INTEGER NOT NULL,
	"ipAddress"	TEXT,
	PRIMARY KEY("id")
)`
