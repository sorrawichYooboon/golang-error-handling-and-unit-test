package taskcontroller

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
	internal "github.com/sorrawichYooboon/golang-error-handling-and-unit-test/internal"
	"github.com/sorrawichYooboon/golang-error-handling-and-unit-test/models"
	mockService "github.com/sorrawichYooboon/golang-error-handling-and-unit-test/services/mock"
)

func TestTaskController_GetTasks(t *testing.T) {
	mockTaskSv := mockService.NewITaskService(t)

	clearAllMock := func() {
		mockTaskSv.ClearAll()
	}

	tests := []struct {
		name        string
		mockPayload interface{}

		wantServiceOrRepoCallWithAndResponse func()
		wantServiceOrRepoCallTimes           map[string]map[string]int
		wantStatusCode                       int
		wantMainServiceResponse              interface{}
	}{
		{
			name:        "Test should return error if get tasks of task service return error",
			mockPayload: nil,
			wantServiceOrRepoCallWithAndResponse: func() {
				mockTaskSv.On("GetTasks").Return([]models.Task{}, errors.New("something went wrong"))
			},
			wantServiceOrRepoCallTimes: map[string]map[string]int{
				"taskService": {
					"GetTasks": 1,
				},
			},
			wantStatusCode: http.StatusInternalServerError,
		},
		{
			name:        "Test should return tasks if get tasks of task service return tasks",
			mockPayload: nil,
			wantServiceOrRepoCallWithAndResponse: func() {
				mockTaskSv.On("GetTasks").Return([]models.Task{
					{
						Id:          "1",
						Title:       "task 1",
						Description: "task 1 description",
					},
				}, nil)
			},
			wantServiceOrRepoCallTimes: map[string]map[string]int{
				"taskService": {
					"GetTasks": 1,
				},
			},
			wantStatusCode: http.StatusOK,
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

			jsonByte := internal.MarshalJSONData(test.mockPayload)
			request := internal.CreateHTTPRequest(http.MethodPost, "/mock-endpoint", jsonByte)

			response := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(response)
			ctx.Request = request

			if test.wantServiceOrRepoCallWithAndResponse != nil {
				test.wantServiceOrRepoCallWithAndResponse()
			}

			taskController := NewTaskController(mockTaskSv)
			taskController.GetTasks(ctx)

			if test.wantStatusCode != 0 {
				assert.Equal(t, test.wantStatusCode, response.Code)
			}

			if test.wantMainServiceResponse != nil {
				var gotResponse []models.Task
				internal.UnmarshalJSONData(response.Body.Bytes(), &gotResponse)
				assert.Equal(t, test.wantMainServiceResponse, gotResponse)
			}

			for serviceName, serviceCallTimes := range test.wantServiceOrRepoCallTimes {
				for methodName, times := range serviceCallTimes {
					switch serviceName {
					case "taskService":
						mockTaskSv.AssertNumberOfCalls(t, methodName, times)
					default:
						t.Errorf("service %s or method %s not found", serviceName, methodName)
					}
				}
			}
		})
	}
}
