package internal

import (
	"encoding/json"
	"os"

	"github.com/sorrawichYooboon/golang-error-handling-and-unit-test/models"
)

func ReadTasksFromFile(filename string) ([]models.Task, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var tasks []models.Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}
