package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNodeSensor_IsValid_SuccessAndFailures(t *testing.T) {
	// valid without id check
	ns := NodeSensor{NodeId: 1, SensorTypeId: 2, Name: "S1", Pin: 0, DHTType: 11}
	ok, err := ns.IsValid(false)
	assert.True(t, ok)
	assert.NoError(t, err)

	// missing name
	ns2 := NodeSensor{NodeId: 1, SensorTypeId: 2, Name: " ", Pin: 0}
	ok2, err2 := ns2.IsValid(false)
	assert.False(t, ok2)
	assert.Error(t, err2)

	// check id required but invalid
	ns3 := NodeSensor{Id: 0, NodeId: 1, SensorTypeId: 2, Name: "S"}
	ok3, err3 := ns3.IsValid(true)
	assert.False(t, ok3)
	assert.Error(t, err3)
}
