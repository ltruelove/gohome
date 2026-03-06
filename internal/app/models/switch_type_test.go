package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwitchType_IsValid(t *testing.T) {
	s := SwitchType{Name: "Toggle"}
	ok, err := s.IsValid(false)
	assert.True(t, ok)
	assert.NoError(t, err)

	s2 := SwitchType{Name: ""}
	ok2, err2 := s2.IsValid(false)
	assert.False(t, ok2)
	assert.Error(t, err2)

	s3 := SwitchType{Id: 0, Name: "X"}
	ok3, err3 := s3.IsValid(true)
	assert.False(t, ok3)
	assert.Error(t, err3)
}
