package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestValidateTask_Success(t *testing.T) {
	task := &Task{
		ID:        "1",
		Title:     "Test Task",
		Status:    StatusPending,
		Priority:  PriorityMedium,
		CreatedAt: time.Now(),
	}

	err := ValidateTask(task)
	assert.NoError(t, err)
}

func TestValidateTask_Fail_RequiredTitle(t *testing.T) {
	task := &Task{
		ID:        "1",
		Status:    StatusPending,
		Priority:  PriorityMedium,
		CreatedAt: time.Now(),
	}

	err := ValidateTask(task)
	assert.Error(t, err)
	assert.Equal(t, "o campo Title é obrigatório", err.Error())
}

func TestValidateTask_Fail_InvalidPriority(t *testing.T) {
	task := &Task{
		ID:        "1",
		Title:     "Test Task",
		Status:    StatusPending,
		Priority:  "invalid",
		CreatedAt: time.Now(),
	}

	err := ValidateTask(task)
	assert.Error(t, err)
	assert.Equal(t, "o campo Priority deve ser 'low', 'medium' ou 'high'", err.Error())
}

func TestValidateTask_Fail_EmptyPriority(t *testing.T) {
	task := &Task{
		ID:        "1",
		Title:     "Test Task",
		Status:    StatusPending,
		Priority:  "",
		CreatedAt: time.Now(),
	}

	err := ValidateTask(task)
	assert.Error(t, err)
	assert.Equal(t, "o campo Priority é obrigatório", err.Error())
}
