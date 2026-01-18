package model

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

type Status string
type Priority string

const (
	StatusPending   Status = "pending"
	StatusCompleted Status = "done"

	PriorityLow    Priority = "low"
	PriorityMedium Priority = "medium"
	PriorityHigh   Priority = "high"
)

var validate *validator.Validate

func init() {
	validate = validator.New()

	validate.RegisterValidation("priority", validatePriority)
}

type Task struct {
	ID          string     `json:"id"`
	Title       string     `json:"title" validate:"required"`
	Status      Status     `json:"status"`
	Priority    Priority   `json:"priority" validate:"priority"`
	Tags        []string   `json:"tags,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}

func ValidateTask(t *Task) error {
	if err := validate.Struct(t); err != nil {
		return formatValidationError(err)
	}
	return nil
}

func validatePriority(fl validator.FieldLevel) bool {
	priority := fl.Field().String()
	return priority == string(PriorityLow) || priority == string(PriorityMedium) || priority == string(PriorityHigh)
}

func formatValidationError(err error) error {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			field := e.Field()
			tag := e.Tag()

			switch tag {
			case "required":
				return fmt.Errorf("o campo %s é obrigatório", field)
			case "priority":
				return fmt.Errorf("o campo %s deve ser 'low', 'medium' ou 'high'", field)
			default:
				return fmt.Errorf("erro de validação no campo %s: %s", field, tag)
			}
		}
	}
	return err
}
