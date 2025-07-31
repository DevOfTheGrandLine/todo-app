package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"todo-app/internal/storage"
	"todo-app/internal/todo"
)

const defaultFile = "tasks.json"

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	// Инициализация менеджера задач
	tm := &todo.TaskManager{}

	// Загрузка задач из файла
	tasks, err := storage.LoadJSON(defaultFile)
	if err != nil {
		fmt.Printf("Ошибка загрузки задач: %v\n", err)
		os.Exit(1)
	}
	tm.SetTasks(tasks)

	// Обработка команд
	switch os.Args[1] {
	case "add":
		handleAdd(tm)
	case "list":
		handleList(tm)
	case "complete":
		handleComplete(tm)
	case "delete":
		handleDelete(tm)
	case "export":
		handleExport(tm)
	case "load":
		handleImport(tm)
	default:
		printUsage()
		os.Exit(1)
	}

	// Сохранение задач
	if err := storage.SaveJSON(tm.GetTasks(), defaultFile); err != nil {
		fmt.Printf("Ошибка сохранения задач: %v\n", err)
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Использование: todo [команда]")
	fmt.Println("Команды:")
	fmt.Println("  add      --desc='описание'   - Добавить задачу")
	fmt.Println("  list     [--filter=all|done|pending] - Вывести задачи")
	fmt.Println("  complete --id=ID             - Отметить задачу выполненной")
	fmt.Println("  delete   --id=ID             - Удалить задачу")
	fmt.Println("  export   --format=json|csv --out=ФАЙЛ - Экспорт задач")
	fmt.Println("  load     --file=ФАЙЛ         - Импорт задач")
}

func handleAdd(tm *todo.TaskManager) {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	desc := addCmd.String("desc", "", "Описание задачи")
	addCmd.Parse(os.Args[2:])

	if *desc == "" {
		fmt.Println("Ошибка: необходимо указать описание задачи (--desc)")
		os.Exit(1)
	}

	tm.Add(*desc)
	fmt.Println("Задача добавлена.")
}

func handleList(tm *todo.TaskManager) {
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	filter := listCmd.String("filter", "all", "Фильтр задач (all|done|pending)")
	listCmd.Parse(os.Args[2:])

	tasks := tm.List(*filter)
	if len(tasks) == 0 {
		fmt.Println("Нет задач.")
		return
	}

	for _, task := range tasks {
		status := " "
		if task.Done {
			status = "✓"
		}
		fmt.Printf("%d [%s] %s\n", task.ID, status, task.Description)
	}
}

func handleComplete(tm *todo.TaskManager) {
	completeCmd := flag.NewFlagSet("complete", flag.ExitOnError)
	id := completeCmd.Int("id", 0, "ID задачи")
	completeCmd.Parse(os.Args[2:])

	if *id == 0 {
		fmt.Println("Ошибка: необходимо указать ID задачи (--id)")
		os.Exit(1)
	}

	if err := tm.Complete(*id); err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Задача отмечена как выполненная.")
}

func handleDelete(tm *todo.TaskManager) {
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	id := deleteCmd.Int("id", 0, "ID задачи")
	deleteCmd.Parse(os.Args[2:])

	if *id == 0 {
		fmt.Println("Ошибка: необходимо указать ID задачи (--id)")
		os.Exit(1)
	}

	if err := tm.Delete(*id); err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Задача удалена.")
}

func handleExport(tm *todo.TaskManager) {
	exportCmd := flag.NewFlagSet("export", flag.ExitOnError)
	format := exportCmd.String("format", "json", "Формат экспорта (json|csv)")
	outFile := exportCmd.String("out", "", "Выходной файл")
	exportCmd.Parse(os.Args[2:])

	if *outFile == "" {
		fmt.Println("Ошибка: необходимо указать выходной файл (--out)")
		os.Exit(1)
	}

	var err error
	switch strings.ToLower(*format) {
	case "json":
		err = storage.SaveJSON(tm.GetTasks(), *outFile)
	case "csv":
		err = storage.SaveCSV(tm.GetTasks(), *outFile)
	default:
		fmt.Println("Ошибка: неизвестный формат. Используйте json или csv")
		os.Exit(1)
	}

	if err != nil {
		fmt.Printf("Ошибка экспорта: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Задачи экспортированы в %s\n", *outFile)
}

func handleImport(tm *todo.TaskManager) {
	loadCmd := flag.NewFlagSet("load", flag.ExitOnError)
	file := loadCmd.String("file", "", "Файл для импорта")
	loadCmd.Parse(os.Args[2:])

	if *file == "" {
		fmt.Println("Ошибка: необходимо указать файл (--file)")
		os.Exit(1)
	}

	var tasks []todo.Task
	var err error

	if strings.HasSuffix(*file, ".json") {
		tasks, err = storage.LoadJSON(*file)
	} else if strings.HasSuffix(*file, ".csv") {
		tasks, err = storage.LoadCSV(*file)
	} else {
		fmt.Println("Ошибка: неизвестный формат файла. Используйте .json или .csv")
		os.Exit(1)
	}

	if err != nil {
		fmt.Printf("Ошибка импорта: %v\n", err)
		os.Exit(1)
	}

	tm.SetTasks(tasks)
	fmt.Printf("Задачи импортированы из %s\n", *file)
}
