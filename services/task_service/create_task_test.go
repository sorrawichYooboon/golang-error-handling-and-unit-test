package taskservice

import (
	"errors"
	"testing"

	"github.com/sorrawichYooboon/golang-error-handling-and-unit-test/models"
	mockRepo "github.com/sorrawichYooboon/golang-error-handling-and-unit-test/repositories/mock"
	"github.com/stretchr/testify/assert"
)

func Test_taskService_CreateTask(t *testing.T) {
	mockTaskRepo := mockRepo.NewITaskRepository(t)

	clearAllMock := func() {
		mockTaskRepo.ClearAll()
	}

	tests := []struct {
		name               string
		mockServiceRequest models.Task

		wantServiceOrRepoCallWithAndResponse func()
		wantServiceOrRepoCallTimes           map[string]map[string]int
		wantMainServiceError                 error
		wantMainServiceResponse              interface{}
	}{
		{
			name: "Test should return error if not provide task id",
			mockServiceRequest: models.Task{
				Id:          "",
				Title:       "task 1",
				Description: "task 1 description",
			},
			wantMainServiceError: errors.New("task id, title, and description are required"),
		},
		{
			name: "Test should return error if not provide task title",
			mockServiceRequest: models.Task{
				Id:          "1",
				Title:       "",
				Description: "task 1 description",
			},
			wantMainServiceError: errors.New("task id, title, and description are required"),
		},
		{
			name: "Test should return error if not provide task description",
			mockServiceRequest: models.Task{
				Id:          "1",
				Title:       "task 1",
				Description: "",
			},
			wantMainServiceError: errors.New("task id, title, and description are required"),
		},
		{
			name: "Test should return error if create task of task repo return error",
			mockServiceRequest: models.Task{
				Id:          "1",
				Title:       "task 1",
				Description: "task 1 description",
			},
			wantServiceOrRepoCallWithAndResponse: func() {
				mockTaskRepo.On("CreateTask", models.Task{
					Id:          "1",
					Title:       "task 1",
					Description: "task 1 description",
				}).Return(models.Task{}, errors.New("something went wrong"))
			},
			wantServiceOrRepoCallTimes: map[string]map[string]int{
				"taskRepository": {
					"CreateTask": 1,
				},
			},
			wantMainServiceError: errors.New("something went wrong"),
		},
		{
			name: "Test should return task if create task of task repo return task",
			mockServiceRequest: models.Task{
				Id:          "1",
				Title:       "task 1",
				Description: "task 1 description",
			},
			wantServiceOrRepoCallWithAndResponse: func() {
				mockTaskRepo.On("CreateTask", models.Task{
					Id:          "1",
					Title:       "task 1",
					Description: "task 1 description",
				}).Return(models.Task{
					Id:          "1",
					Title:       "task 1",
					Description: "task 1 description",
				}, nil)
			},
			wantServiceOrRepoCallTimes: map[string]map[string]int{
				"taskRepository": {
					"CreateTask": 1,
				},
			},
			wantMainServiceResponse: models.Task{
				Id:          "1",
				Title:       "task 1",
				Description: "task 1 description",
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
			response, err := taskService.CreateTask(test.mockServiceRequest)

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
