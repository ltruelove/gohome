package setup

const StaticData string = `INSERT INTO "SensorType"
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
VALUES ('Toggle');`
