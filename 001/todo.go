package main

import "errors"

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) AddTask(title string) error {
	tasks, err := s.repo.LoadTasks()
	if err != nil {
		return err
	}

	newTask := Task{
		ID:        len(tasks) + 1,
		Title:     title,
		Completed: false,
	}

	tasks = append(tasks, newTask)
	return s.repo.saveTasks(tasks)
}

//to retreive all tasks

func (s *Service) ListTasks() ([]Task, error) {
	return s.repo.LoadTasks()
}

func (s *Service) CompleteTask(id int) error {
	tasks, err := s.repo.LoadTasks()
	if err != nil {
		return err
	}

	found := false
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Completed = true
			found = true
			break
		}
	}

	if !found {
		return errors.New("task not found")
	}

	return s.repo.saveTasks(tasks)
}
