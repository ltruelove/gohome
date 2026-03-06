package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSensorType_IsValid(t *testing.T) {
	s := SensorType{TypeName: "Temp"}
	ok, err := s.IsValid(false)
	assert.True(t, ok)
	assert.NoError(t, err)

	s2 := SensorType{TypeName: ""}
	ok2, err2 := s2.IsValid(false)
	assert.False(t, ok2)
	assert.Error(t, err2)

	s3 := SensorType{Id: 0, TypeName: "T"}
	ok3, err3 := s3.IsValid(true)
	assert.False(t, ok3)
	assert.Error(t, err3)
}
