package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Repository struct {
	FilePath string
}

func NewRepository(filePath string) *Repository {
	return &Repository{FilePath: filePath}
}

func (r *Repository) LoadTasks() ([]Task, error) {
	file, err := os.Open(r.FilePath)
	if err != nil {
		// If file doesn't exist, return an empty slice
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var tasks []Task
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tasks); err != nil {
		// Handle empty file case
		if err.Error() == "EOF" {
			return []Task{}, nil
		}
		return nil, err
	}
	return tasks, nil
}

func (r *Repository) saveTasks(tasks []Task) error {
	file, err := os.Create(r.FilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")
	if err := encoder.Encode(tasks); err != nil {
		return err
	}

	fmt.Println("Tasks saved succefully.")
	return nil
}
