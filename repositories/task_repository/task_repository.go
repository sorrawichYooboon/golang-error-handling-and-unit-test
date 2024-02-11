package taskrepository

import (
	"github.com/pkg/errors"

	"github.com/sorrawichYooboon/golang-error-handling-and-unit-test/models"
)

type taskRepository struct {
	testingTask []models.Task
}

func NewTaskRepository(testingTask []models.Task) ITaskRepository {
	return &taskRepository{
		testingTask: testingTask,
	}
}

func (tr *taskRepository) GetTasks() ([]models.Task, error) {
	return tr.testingTask, nil
}

func (tr *taskRepository) CreateTask(task models.Task) (models.Task, error) {
	tr.testingTask = append(tr.testingTask, task)
	return task, nil
}

func (tr *taskRepository) GetTaskById(id string) (models.Task, error) {
	var task models.Task

	isFound := false
	for _, t := range tr.testingTask {
		if t.Id == id {
			task = t
			isFound = true
			break
		}
	}

	if !isFound {
		return task, errors.New("task not found")
	}

	return task, nil
}

func (tr *taskRepository) UpdateTaskById(id string, task models.Task) (models.Task, error) {
	for i, t := range tr.testingTask {
		if t.Id == id {
			tr.testingTask[i] = task
			break
		}
	}

	return task, nil
}

func (tr *taskRepository) DeleteTaskById(id string) error {
	for i, t := range tr.testingTask {
		if t.Id == id {
			tr.testingTask = append(tr.testingTask[:i], tr.testingTask[i+1:]...)
			break
		}
	}

	return nil
}
