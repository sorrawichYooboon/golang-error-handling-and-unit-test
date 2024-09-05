package taskcontroller

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	internal "github.com/sorrawichYooboon/golang-error-handling-and-unit-test/internal"
	"github.com/sorrawichYooboon/golang-error-handling-and-unit-test/models"
	mockService "github.com/sorrawichYooboon/golang-error-handling-and-unit-test/services/mock"
	"github.com/stretchr/testify/assert"
)

func TestTaskController_GetTaskById(t *testing.T) {
	mockTaskSv := mockService.NewITaskService(t)

	clearAllMock := func() {
		mockTaskSv.ClearAll()
	}

	tests := []struct {
		name        string
		mockPayload interface{}
		id          string

		wantServiceOrRepoCallWithAndResponse func()
		wantServiceOrRepoCallTimes           map[string]map[string]int
		wantStatusCode                       int
		wantMainServiceResponse              interface{}
	}{
		{
			name: "Test should response error if get task by id of task service return error",
			id:   "1",
			wantServiceOrRepoCallWithAndResponse: func() {
				mockTaskSv.On("GetTaskById", "1").Return(models.Task{}, errors.New("something went wrong"))
			},
			wantServiceOrRepoCallTimes: map[string]map[string]int{
				"taskService": {
					"GetTaskById": 1,
				},
			},
			wantStatusCode: 500,
		},
		{
			name: "Test should response success",
			id:   "1",
			mockPayload: models.Task{
				Id:          "1",
				Title:       "task 1",
				Description: "task 1 description",
			},
			wantServiceOrRepoCallWithAndResponse: func() {
				mockTaskSv.On("GetTaskById", "1").Return(models.Task{
					Id:          "1",
					Title:       "task 1",
					Description: "task 1 description",
				}, nil)
			},
			wantServiceOrRepoCallTimes: map[string]map[string]int{
				"taskService": {
					"GetTaskById": 1,
				},
			},
			wantStatusCode: 200,
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
			ctx.Params = append(ctx.Params, gin.Param{Key: "id", Value: test.id})

			if test.wantServiceOrRepoCallWithAndResponse != nil {
				test.wantServiceOrRepoCallWithAndResponse()
			}

			taskController := NewTaskController(mockTaskSv)
			taskController.GetTaskById(ctx)

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
