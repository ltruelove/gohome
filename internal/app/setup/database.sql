CREATE TABLE IF NOT EXISTS "SensorType" (
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
	"SensorTypeDataId"	INTEGER NOT NULL,
	"Name"	TEXT NOT NULL,
	UNIQUE("NodeId","ViewId","NodeSensorId","SensorTypeDataId"),
	FOREIGN KEY("NodeId") REFERENCES "Node"("Id"),
	FOREIGN KEY("ViewId") REFERENCES "View"("Id"),
	FOREIGN KEY("SensorTypeDataId") REFERENCES "SensorTypeData"("Id"),
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

INSERT INTO "SensorType"
("Name")
VALUES ('DHT');

INSERT INTO "SensorType"
("Name")
VALUES ('Moisture');

INSERT INTO "SensorType"
("Name")
VALUES ('Magnetic');

INSERT INTO "SensorType"
("Name")
VALUES ('Photoresistor');

INSERT INTO "SensorTypeData"
("SensorTypeId", "Name", "ValueType")
VALUES (1, 'TemperatureF', 'float');

INSERT INTO "SensorTypeData"
("SensorTypeId", "Name", "ValueType")
VALUES (1, 'TemperatureC', 'float');

INSERT INTO "SensorTypeData"
("SensorTypeId", "Name", "ValueType")
VALUES (1, 'Humidity', 'float');

INSERT INTO "SensorTypeData"
("SensorTypeId", "Name", "ValueType")
VALUES (2, 'Moisture', 'int');

INSERT INTO "SensorTypeData"
("SensorTypeId", "Name", "ValueType")
VALUES (3, 'IsClosed', 'int');

INSERT INTO "SensorTypeData"
("SensorTypeId", "Name", "ValueType")
VALUES (4, 'ResistorValue', 'int');

INSERT INTO "SwitchType"
("Name")
VALUES ('Momentary');

INSERT INTO "SwitchType"
("Name")
VALUES ('Toggle');