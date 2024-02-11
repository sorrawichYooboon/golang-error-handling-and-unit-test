package taskservice

import (
	"github.com/pkg/errors"
	ce "github.com/sorrawichYooboon/golang-error-handling-and-unit-test/errors"
	"github.com/sorrawichYooboon/golang-error-handling-and-unit-test/models"
	taskrepository "github.com/sorrawichYooboon/golang-error-handling-and-unit-test/repositories/task_repository"
)

type taskService struct {
	taskRepo taskrepository.ITaskRepository
}

func NewTaskService(taskRepo taskrepository.ITaskRepository) ITaskService {
	return &taskService{
		taskRepo: taskRepo,
	}
}

func (ts *taskService) GetTasks() ([]models.Task, error) {
	tasks, err := ts.taskRepo.GetTasks()
	if err != nil {
		return nil, ce.ErrorInternal(err)
	}

	return tasks, nil
}

func (ts *taskService) CreateTask(task models.Task) (models.Task, error) {
	if task.Id == "" || task.Title == "" || task.Description == "" {
		return models.Task{}, ce.ErrorInvalidFormat(errors.New("task id, title, and description are required"))
	}

	task, err := ts.taskRepo.CreateTask(task)
	if err != nil {
		return models.Task{}, ce.ErrorInternal(err)
	}

	return task, nil
}

func (ts *taskService) GetTaskById(id string) (models.Task, error) {
	if id == "" {
		return models.Task{}, ce.ErrorInvalidFormat(errors.New("task id is required"))
	}

	task, err := ts.taskRepo.GetTaskById(id)
	if err != nil {
		return models.Task{}, ce.ErrorInternal(err)
	}

	return task, nil
}

func (ts *taskService) UpdateTaskById(id string, task models.Task) (models.Task, error) {
	if id == "" {
		return models.Task{}, ce.ErrorInvalidFormat(errors.New("task id is required"))
	}

	if task.Id == "" || task.Title == "" || task.Description == "" {
		return models.Task{}, ce.ErrorInvalidFormat(errors.New("task id, title, and description are required"))
	}

	taskById, err := ts.taskRepo.GetTaskById(id)
	if err != nil {
		return models.Task{}, ce.ErrorInternal(err)
	}

	task, err = ts.taskRepo.UpdateTaskById(taskById.Id, task)
	if err != nil {
		return models.Task{}, ce.ErrorInternal(err)
	}

	return task, nil
}

func (ts *taskService) DeleteTaskById(id string) error {
	if id == "" {
		return ce.ErrorInvalidFormat(errors.New("task id is required"))
	}

	taskById, err := ts.taskRepo.GetTaskById(id)
	if err != nil {
		return ce.ErrorInternal(err)
	}

	err = ts.taskRepo.DeleteTaskById(taskById.Id)
	if err != nil {
		return ce.ErrorInternal(err)
	}

	return nil
}
