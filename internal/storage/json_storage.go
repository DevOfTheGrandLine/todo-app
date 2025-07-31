package storage

import (
	"encoding/json"
	"os"
	"todo-app/internal/todo"
)

func LoadJSON(filename string) ([]todo.Task, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return []todo.Task{}, nil
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var tasks []todo.Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func SaveJSON(tasks []todo.Task, filename string) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}
