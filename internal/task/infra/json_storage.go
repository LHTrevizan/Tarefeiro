package infra

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"tarefeiro/internal/task/model"
)

type JSONStorage struct {
	FilePath string
}

func NewJSONStorage(filePath string) (*JSONStorage, error) {
	if filePath == "" {
		return nil, errors.New("file path não pode ser vazio")
	}

	dir := filepath.Dir(filePath)

	// cria diretório se não existir
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}

	// cria arquivo se não existir
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			return nil, err
		}
		defer file.Close()
	}

	return &JSONStorage{
		FilePath: filePath,
	}, nil
}

func (r *JSONStorage) Load() ([]model.Task, error) {
	data, err := os.ReadFile(r.FilePath)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return []model.Task{}, nil
	}

	var tasks []model.Task
	err = json.Unmarshal(data, &tasks)
	return tasks, err
}

func (r *JSONStorage) Save(tasks []model.Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(r.FilePath, data, 0644)
}
