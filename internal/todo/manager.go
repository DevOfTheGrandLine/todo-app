package todo

import (
	"errors"
	"sort"
)

type TaskManager struct {
	tasks []Task
}

func (tm *TaskManager) Add(description string) {
	maxID := 0
	for _, task := range tm.tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}

	tm.tasks = append(tm.tasks, Task{
		ID:          maxID + 1,
		Description: description,
		Done:        false,
	})
}

func (tm *TaskManager) List(filter string) []Task {
	var result []Task

	switch filter {
	case "done":
		for _, t := range tm.tasks {
			if t.Done {
				result = append(result, t)
			}
		}
	case "pending":
		for _, t := range tm.tasks {
			if !t.Done {
				result = append(result, t)
			}
		}
	default: // "all"
		result = tm.tasks
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].ID < result[j].ID
	})
	return result
}

func (tm *TaskManager) Complete(id int) error {
	for i := range tm.tasks {
		if tm.tasks[i].ID == id {
			tm.tasks[i].Done = true
			return nil
		}
	}
	return errors.New("task not found")
}

func (tm *TaskManager) Delete(id int) error {
	for i, task := range tm.tasks {
		if task.ID == id {
			tm.tasks = append(tm.tasks[:i], tm.tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}

// Вспомогательные методы для работы с хранилищем
func (tm *TaskManager) GetTasks() []Task {
	return tm.tasks
}

func (tm *TaskManager) SetTasks(tasks []Task) {
	tm.tasks = tasks
}
