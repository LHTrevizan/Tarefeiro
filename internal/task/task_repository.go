package task

import (
	"encoding/json"
	"os"
)

type Repository struct {
	File string
}

func (r *Repository) Load() ([]Task, error) {
	if _, err := os.Stat(r.File); os.IsNotExist(err) {
		return []Task{}, nil
	}

	data, err := os.ReadFile(r.File)
	if err != nil {
		return nil, err
	}

	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	return tasks, err
}

func (r *Repository) Save(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(r.File, data, 0644)
}
