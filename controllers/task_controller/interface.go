package taskcontroller

import "github.com/gin-gonic/gin"

type ITaskController interface {
	GetTasks(c *gin.Context)
	CreateTask(c *gin.Context)
	GetTaskById(c *gin.Context)
	UpdateTaskById(c *gin.Context)
	DeleteTaskById(c *gin.Context)
}
