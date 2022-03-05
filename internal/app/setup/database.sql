CREATE TABLE "TemperatureSensors" (
	"id"	TEXT NOT NULL UNIQUE,
	"name"	TEXT NOT NULL,
	"isGarage"	INTEGER NOT NULL,
	"ipAddress"	TEXT,
	PRIMARY KEY("id")
)