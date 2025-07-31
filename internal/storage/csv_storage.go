package storage

import (
	"encoding/csv"
	"os"
	"strconv"
	"todo-app/internal/todo"
)

func LoadCSV(filename string) ([]todo.Task, error) {
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return []todo.Task{}, nil
		}
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var tasks []todo.Task
	for i, record := range records {
		if i == 0 { // Пропуск заголовка
			continue
		}
		if len(record) < 3 {
			continue
		}

		id, _ := strconv.Atoi(record[0])
		done, _ := strconv.ParseBool(record[2])

		tasks = append(tasks, todo.Task{
			ID:          id,
			Description: record[1],
			Done:        done,
		})
	}
	return tasks, nil
}

func SaveCSV(tasks []todo.Task, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Заголовок
	if err := writer.Write([]string{"id", "description", "done"}); err != nil {
		return err
	}

	for _, task := range tasks {
		record := []string{
			strconv.Itoa(task.ID),
			task.Description,
			strconv.FormatBool(task.Done),
		}
		if err := writer.Write(record); err != nil {
			return err
		}
	}
	return nil
}
