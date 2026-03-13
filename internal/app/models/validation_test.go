package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTempLogData_IsValid(t *testing.T) {
	// valid
	t1 := TempLogData{Id: 1, NodeSensorLogId: 2, TemperatureF: 72.0, TemperatureC: 22.0, Humidity: 50.0}
	ok, err := t1.IsValid(true)
	assert.True(t, ok)
	assert.NoError(t, err)

	// invalid NodeSensorLogId
	t2 := TempLogData{Id: 2, NodeSensorLogId: 0, TemperatureF: 72.0, TemperatureC: 22.0, Humidity: 50.0}
	ok2, err2 := t2.IsValid(true)
	assert.False(t, ok2)
	assert.Error(t, err2)

	// invalid temperature (less than absolute zero)
	t3 := TempLogData{Id: 3, NodeSensorLogId: 4, TemperatureF: -1000.0, TemperatureC: -1000.0, Humidity: 10.0}
	ok3, err3 := t3.IsValid(true)
	assert.False(t, ok3)
	assert.Error(t, err3)
}

func TestNodeSensorLog_IsValid(t *testing.T) {
	now := time.Now()
	// valid
	n1 := NodeSensorLog{Id: 1, NodeId: 1, DateLogged: now}
	ok, err := n1.IsValid(true)
	assert.True(t, ok)
	assert.NoError(t, err)

	// missing node id
	n2 := NodeSensorLog{Id: 2, NodeId: 0, DateLogged: now}
	ok2, err2 := n2.IsValid(true)
	assert.False(t, ok2)
	assert.Error(t, err2)

	// zero date
	n3 := NodeSensorLog{Id: 3, NodeId: 1, DateLogged: time.Time{}}
	ok3, err3 := n3.IsValid(true)
	assert.False(t, ok3)
	assert.Error(t, err3)
}

func TestControlPoint_IsValid(t *testing.T) {
	// valid
	cp := ControlPoint{Id: 1, Name: "cp", IpAddress: "127.0.0.1", Mac: "AA:BB:CC"}
	ok, err := cp.IsValid(true)
	assert.True(t, ok)
	assert.NoError(t, err)

	// missing fields
	cp2 := ControlPoint{Id: 0, Name: "", IpAddress: "", Mac: ""}
	ok2, err2 := cp2.IsValid(true)
	assert.False(t, ok2)
	assert.Error(t, err2)
	assert.Contains(t, err2.Error(), "Name cannot be empty")
	assert.Contains(t, err2.Error(), "IpAddress")
}

func TestNodeSensor_IsValid(t *testing.T) {
	// valid
	s := NodeSensor{Id: 1, NodeId: 1, SensorTypeId: 1, Name: "s1", Pin: 4, DHTType: 0}
	ok, err := s.IsValid(true)
	assert.True(t, ok)
	assert.NoError(t, err)

	// missing NodeId
	s2 := NodeSensor{Id: 2, NodeId: 0, SensorTypeId: 1, Name: "s2", Pin: 2}
	ok2, err2 := s2.IsValid(true)
	assert.False(t, ok2)
	assert.Error(t, err2)

	// missing SensorTypeId
	s3 := NodeSensor{Id: 3, NodeId: 1, SensorTypeId: 0, Name: "s3", Pin: 2}
	ok3, err3 := s3.IsValid(true)
	assert.False(t, ok3)
	assert.Error(t, err3)

	// empty name
	s4 := NodeSensor{Id: 4, NodeId: 1, SensorTypeId: 1, Name: "", Pin: 2}
	ok4, err4 := s4.IsValid(true)
	assert.False(t, ok4)
	assert.Error(t, err4)

	// negative pin
	s5 := NodeSensor{Id: 5, NodeId: 1, SensorTypeId: 1, Name: "s5", Pin: -1}
	ok5, err5 := s5.IsValid(true)
	assert.False(t, ok5)
	assert.Error(t, err5)
}

func TestNodeSwitch_IsValid(t *testing.T) {
	// valid
	sw := NodeSwitch{Id: 1, NodeId: 1, SwitchTypeId: 1, Name: "sw1", Pin: 3, MomentaryPressDuration: 100, IsClosedOn: true}
	ok, err := sw.IsValid(true)
	assert.True(t, ok)
	assert.NoError(t, err)

	// missing NodeId
	sw2 := NodeSwitch{Id: 2, NodeId: 0, SwitchTypeId: 1, Name: "sw2", Pin: 3}
	ok2, err2 := sw2.IsValid(true)
	assert.False(t, ok2)
	assert.Error(t, err2)

	// missing SwitchTypeId
	sw3 := NodeSwitch{Id: 3, NodeId: 1, SwitchTypeId: 0, Name: "sw3", Pin: 3}
	ok3, err3 := sw3.IsValid(true)
	assert.False(t, ok3)
	assert.Error(t, err3)

	// empty name
	sw4 := NodeSwitch{Id: 4, NodeId: 1, SwitchTypeId: 1, Name: "", Pin: 3}
	ok4, err4 := sw4.IsValid(true)
	assert.False(t, ok4)
	assert.Error(t, err4)

	// negative pin
	sw5 := NodeSwitch{Id: 5, NodeId: 1, SwitchTypeId: 1, Name: "sw5", Pin: -2}
	ok5, err5 := sw5.IsValid(true)
	assert.False(t, ok5)
	assert.Error(t, err5)
}
