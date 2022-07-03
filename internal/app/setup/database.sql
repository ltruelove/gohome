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
	FOREIGN KEY("SensorTypeId") REFERENCES "SensorTypes"("Id")
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
	FOREIGN KEY("NodeId") REFERENCES "Node"("Id"),
	FOREIGN KEY("SensorTypeId") REFERENCES "SensorTypes"("Id"),
	PRIMARY KEY("Id" AUTOINCREMENT)
);

CREATE TABLE IF NOT EXISTS "NodeSwitch" (
	"Id"	INTEGER NOT NULL UNIQUE,
	"NodeId"	INTEGER NOT NULL,
	"SwitchTypeId"	INTEGER NOT NULL,
	"Name"	TEXT NOT NULL,
	"Pin"	INTEGER NOT NULL,
	"MomentaryPressDuration"	INTEGER,
	"IsClosedOn"	INTEGER,
	FOREIGN KEY("NodeId") REFERENCES "Node"("Id"),
	FOREIGN KEY("SwitchTypeId") REFERENCES "SwitchType"("Id"),
	PRIMARY KEY("Id" AUTOINCREMENT)
);

CREATE TABLE IF NOT EXISTS "NodeController" (
	"Id"	INTEGER NOT NULL UNIQUE,
	"Name"	TEXT NOT NULL,
	"IpAddress"	TEXT NOT NULL,
	PRIMARY KEY("Id" AUTOINCREMENT)
);

CREATE TABLE IF NOT EXISTS "ControllerNodes" (
	"Id"	INTEGER NOT NULL UNIQUE,
	"ControllerId"	INTEGER NOT NULL,
	"NodeId"	INTEGER NOT NULL,
	FOREIGN KEY("NodeId") REFERENCES "Node"("Id"),
	FOREIGN KEY("ControllerId") REFERENCES "NodeController"("Id"),
	PRIMARY KEY("Id" AUTOINCREMENT)
);

INSERT INTO "main"."SensorTypes"
("Name")
VALUES ('DHT');

INSERT INTO "main"."SensorTypes"
("Name")
VALUES ('Moisture');

INSERT INTO "main"."SensorTypes"
("Name")
VALUES ('Magnetic');

INSERT INTO "main"."SensorTypes"
("Name")
VALUES ('Photoresistor');

INSERT INTO "main"."SensorTypeData"
("SensorTypeId", "Name", "ValueType")
VALUES (1, 'TemperatureF', 'float');

INSERT INTO "main"."SensorTypeData"
("SensorTypeId", "Name", "ValueType")
VALUES (1, 'TemperatureC', 'float');

INSERT INTO "main"."SensorTypeData"
("SensorTypeId", "Name", "ValueType")
VALUES (1, 'Humidity', 'float');

INSERT INTO "main"."SensorTypeData"
("SensorTypeId", "Name", "ValueType")
VALUES (2, 'Moisture', 'int');

INSERT INTO "main"."SensorTypeData"
("SensorTypeId", "Name", "ValueType")
VALUES (3, 'IsClosed', 'int');

INSERT INTO "main"."SensorTypeData"
("SensorTypeId", "Name", "ValueType")
VALUES (4, 'ResistorValue', 'int');

INSERT INTO "main"."SwitchTypes"
("Name")
VALUES ('Momentary');

INSERT INTO "main"."SwitchTypes"
("Name")
VALUES ('Toggle');