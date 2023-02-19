package setup

const CreateTables string = `CREATE TABLE IF NOT EXISTS "SensorType" (
	"Id"	INTEGER NOT NULL UNIQUE,
	"Name"	TEXT NOT NULL UNIQUE,
	PRIMARY KEY("Id" AUTOINCREMENT)
);

CREATE TABLE IF NOT EXISTS "SensorTypeData" (
	"Id"	INTEGER NOT NULL UNIQUE,
	"SensorTypeId"	INTEGER NOT NULL,
	"Name"	TEXT,
	"ValueType"	TEXT,
	PRIMARY KEY("Id" AUTOINCREMENT),
	FOREIGN KEY("SensorTypeId") REFERENCES "SensorType"("Id")
);

CREATE TABLE IF NOT EXISTS "SwitchType" (
	"Id"	INTEGER NOT NULL UNIQUE,
	"Name"	TEXT UNIQUE,
	PRIMARY KEY("Id" AUTOINCREMENT)
);

CREATE TABLE IF NOT EXISTS "Node" (
	"Id"	INTEGER NOT NULL UNIQUE,
	"Mac"	TEXT UNIQUE,
	"Name"	TEXT UNIQUE,
	PRIMARY KEY("Id" AUTOINCREMENT)
);

CREATE TABLE IF NOT EXISTS "NodeSensor" (
	"Id"	INTEGER NOT NULL UNIQUE,
	"NodeId"	INTEGER NOT NULL,
	"SensorTypeId"	INTEGER NOT NULL,
	"Name"	TEXT NOT NULL,
	"Pin"	INTEGER NOT NULL,
	"DHTType"	INTEGER,
	FOREIGN KEY("NodeId") REFERENCES "Node"("Id"),
	FOREIGN KEY("SensorTypeId") REFERENCES "SensorType"("Id"),
	PRIMARY KEY("Id" AUTOINCREMENT)
);

CREATE TABLE IF NOT EXISTS "NodeSwitch" (
	"Id"	INTEGER NOT NULL UNIQUE,
	"NodeId"	INTEGER NOT NULL,
	"SwitchTypeId"	INTEGER NOT NULL,
	"Name"	TEXT NOT NULL,
	"Pin"	INTEGER NOT NULL,
	"MomentaryPressDuration"	INTEGER DEFAULT 100,
	"IsClosedOn"	INTEGER DEFAULT 1,
	FOREIGN KEY("NodeId") REFERENCES "Node"("Id"),
	FOREIGN KEY("SwitchTypeId") REFERENCES "SwitchType"("Id"),
	PRIMARY KEY("Id" AUTOINCREMENT)
);

CREATE TABLE IF NOT EXISTS "ControlPoint" (
	"Id"	INTEGER NOT NULL UNIQUE,
	"Name"	TEXT NOT NULL,
	"IpAddress"	TEXT NOT NULL,
	"Mac"	TEXT NOT NULL,
	PRIMARY KEY("Id" AUTOINCREMENT)
);

CREATE TABLE IF NOT EXISTS "ControlPointNodes" (
	"Id"	INTEGER NOT NULL UNIQUE,
	"ControlPointId"	INTEGER NOT NULL,
	"NodeId"	INTEGER NOT NULL,
	FOREIGN KEY("NodeId") REFERENCES "Node"("Id"),
	FOREIGN KEY("ControlPointId") REFERENCES "ControlPoint"("Id"),
	PRIMARY KEY("Id" AUTOINCREMENT)
);

CREATE TABLE IF NOT EXISTS "View" (
	"Id"	INTEGER NOT NULL UNIQUE,
	"Name"	TEXT NOT NULL,
	PRIMARY KEY("Id" AUTOINCREMENT)
);

CREATE TABLE IF NOT EXISTS "ViewNodeSensorData" (
	"Id"	INTEGER NOT NULL UNIQUE,
	"NodeId"	INTEGER NOT NULL,
	"ViewId"	INTEGER NOT NULL,
	"NodeSensorId"	INTEGER NOT NULL,
	"Name"	TEXT NOT NULL,
	UNIQUE("NodeId","ViewId","NodeSensorId"),
	FOREIGN KEY("NodeId") REFERENCES "Node"("Id"),
	FOREIGN KEY("ViewId") REFERENCES "View"("Id"),
	FOREIGN KEY("NodeSensorId") REFERENCES "NodeSensor"("Id"),
	PRIMARY KEY("Id" AUTOINCREMENT)
);

CREATE TABLE IF NOT EXISTS "ViewNodeSwitchData" (
	"Id"	INTEGER NOT NULL UNIQUE,
	"NodeId"	INTEGER NOT NULL,
	"ViewId"	INTEGER NOT NULL,
	"NodeSwitchId"	INTEGER NOT NULL,
	"Name"	TEXT NOT NULL,
	UNIQUE("NodeId","ViewId","NodeSwitchId"),
	PRIMARY KEY("Id" AUTOINCREMENT),
	FOREIGN KEY("ViewId") REFERENCES "View"("Id"),
	FOREIGN KEY("NodeSwitchId") REFERENCES "NodeSwitch"("Id"),
	FOREIGN KEY("NodeId") REFERENCES "Node"("Id")
);

CREATE TABLE IF NOT EXISTS "NodeSensorLog" (
	"Id"	INTEGER NOT NULL UNIQUE,
	"NodeId"	INTEGER NOT NULL,
	"DateLogged"	TEXT NOT NULL,
	PRIMARY KEY("Id" AUTOINCREMENT),
	FOREIGN KEY("NodeId") REFERENCES "Node"("Id")
);

CREATE TABLE IF NOT EXISTS "TempLog" (
	"Id"	INTEGER NOT NULL UNIQUE,
	"NodeSensorLogId"	INTEGER NOT NULL,
	"TemperatureF"	REAL NOT NULL,
	"TemperatureC"	REAL NOT NULL,
	"Humidity"	REAL NOT NULL,
	FOREIGN KEY("NodeSensorLogId") REFERENCES "NodeSensorLog"("Id"),
	PRIMARY KEY("Id" AUTOINCREMENT)
);

CREATE TABLE IF NOT EXISTS "MoistureLog" (
	"Id"	INTEGER NOT NULL UNIQUE,
	"NodeSensorLogId"	INTEGER NOT NULL,
	"Moisture"	INTEGER NOT NULL,
	FOREIGN KEY("NodeSensorLogId") REFERENCES "NodeSensorLog"("Id"),
	PRIMARY KEY("Id" AUTOINCREMENT)
);

CREATE TABLE IF NOT EXISTS "MagneticLog" (
	"Id"	INTEGER NOT NULL UNIQUE,
	"NodeSensorLogId"	INTEGER NOT NULL,
	"IsClosed"	INTEGER NOT NULL DEFAULT 0,
	FOREIGN KEY("NodeSensorLogId") REFERENCES "NodeSensorLog"("Id"),
	PRIMARY KEY("Id" AUTOINCREMENT)
);

CREATE TABLE IF NOT EXISTS "ResistorLog" (
	"Id"	INTEGER NOT NULL UNIQUE,
	"NodeSensorLogId"	INTEGER NOT NULL,
	"ResistorValue"	INTEGER NOT NULL,
	FOREIGN KEY("NodeSensorLogId") REFERENCES "NodeSensorLog"("Id"),
	PRIMARY KEY("Id" AUTOINCREMENT)
);
`
