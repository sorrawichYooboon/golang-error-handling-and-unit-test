package taskcontroller

import (
	"github.com/gin-gonic/gin"
	ce "github.com/sorrawichYooboon/golang-error-handling-and-unit-test/errors"
	"github.com/sorrawichYooboon/golang-error-handling-and-unit-test/models"
	taskservice "github.com/sorrawichYooboon/golang-error-handling-and-unit-test/services/task_service"
)

type TaskController struct {
	taskSv taskservice.ITaskService
}

func NewTaskController(taskSv taskservice.ITaskService) ITaskController {
	return &TaskController{
		taskSv: taskSv,
	}
}

func (tc *TaskController) GetTasks(c *gin.Context) {
	response, err := tc.taskSv.GetTasks()
	if err != nil {
		ce.HandleErrorResponse(c, err)
		return
	}

	c.JSON(200, response)
}

func (tc *TaskController) CreateTask(c *gin.Context) {
	var request models.Task

	if err := c.ShouldBindJSON(&request); err != nil {
		ce.HandleErrorResponse(c, ce.ErrorInvalidFormat(err))

		return
	}

	response, err := tc.taskSv.CreateTask(request)
	if err != nil {
		ce.HandleErrorResponse(c, err)
		return
	}

	c.JSON(201, response)
}

func (tc *TaskController) GetTaskById(c *gin.Context) {
	id := c.Param("id")

	response, err := tc.taskSv.GetTaskById(id)
	if err != nil {
		ce.HandleErrorResponse(c, err)
		return
	}

	c.JSON(200, response)
}

func (tc *TaskController) UpdateTaskById(c *gin.Context) {
	var request models.Task
	id := c.Param("id")

	if err := c.ShouldBindJSON(&request); err != nil {
		ce.HandleErrorResponse(c, ce.ErrorInvalidFormat(err))
		return
	}

	response, err := tc.taskSv.UpdateTaskById(id, request)
	if err != nil {
		ce.HandleErrorResponse(c, err)
		return
	}

	c.JSON(200, response)
}

func (tc *TaskController) DeleteTaskById(c *gin.Context) {
	id := c.Param("id")

	err := tc.taskSv.DeleteTaskById(id)
	if err != nil {
		ce.HandleErrorResponse(c, err)
		return
	}

	c.JSON(204, "")
}
