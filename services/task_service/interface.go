package taskservice

import "github.com/sorrawichYooboon/golang-error-handling-and-unit-test/models"

type ITaskService interface {
	GetTasks() ([]models.Task, error)
	CreateTask(task models.Task) (models.Task, error)
	GetTaskById(id string) (models.Task, error)
	UpdateTaskById(id string, task models.Task) (models.Task, error)
	DeleteTaskById(id string) error
}
