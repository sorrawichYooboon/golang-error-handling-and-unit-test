package taskservice

import (
	"errors"
	"testing"

	"github.com/sorrawichYooboon/golang-error-handling-and-unit-test/models"
	mockRepo "github.com/sorrawichYooboon/golang-error-handling-and-unit-test/repositories/mock"
	"github.com/stretchr/testify/assert"
)

func Test_taskService_GetTasks(t *testing.T) {
	mockTaskRepo := mockRepo.NewITaskRepository(t)

	clearAllMock := func() {
		mockTaskRepo.ClearAll()
	}

	tests := []struct {
		name               string
		mockServiceRequest interface{}

		wantServiceOrRepoCallWithAndResponse func()
		wantServiceOrRepoCallTimes           map[string]map[string]int
		wantMainServiceError                 error
		wantMainServiceResponse              interface{}
	}{
		{
			name: "Test should return error if get tasks of task repo return error",
			wantServiceOrRepoCallWithAndResponse: func() {
				mockTaskRepo.On("GetTasks").Return(nil, errors.New("something went wrong"))
			},
			wantServiceOrRepoCallTimes: map[string]map[string]int{
				"taskRepository": {
					"GetTasks": 1,
				},
			},
			wantMainServiceError: errors.New("something went wrong"),
		},
		{
			name: "Test should return tasks if get tasks of task repo return tasks",
			wantServiceOrRepoCallWithAndResponse: func() {
				mockTaskRepo.On("GetTasks").Return([]models.Task{
					{
						Id:          "1",
						Title:       "task 1",
						Description: "task 1 description",
					},
				}, nil)
			},
			wantServiceOrRepoCallTimes: map[string]map[string]int{
				"taskRepository": {
					"GetTasks": 1,
				},
			},
			wantMainServiceResponse: []models.Task{
				{
					Id:          "1",
					Title:       "task 1",
					Description: "task 1 description",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer clearAllMock()

			if test.wantServiceOrRepoCallWithAndResponse != nil {
				test.wantServiceOrRepoCallWithAndResponse()
			}

			taskService := NewTaskService(mockTaskRepo)
			response, err := taskService.GetTasks()

			if test.wantMainServiceError != nil {
				assert.Equal(t, test.wantMainServiceError.Error(), err.Error())
			}

			if test.wantMainServiceResponse != nil {
				assert.Equal(t, test.wantMainServiceResponse, response)
			}

			for serviceName, serviceCallTimes := range test.wantServiceOrRepoCallTimes {
				for methodName, times := range serviceCallTimes {
					switch serviceName {
					case "taskRepository":
						mockTaskRepo.AssertNumberOfCalls(t, methodName, times)
					default:
						t.Errorf("service %s or method %s not found", serviceName, methodName)
					}
				}
			}
		})
	}
}
