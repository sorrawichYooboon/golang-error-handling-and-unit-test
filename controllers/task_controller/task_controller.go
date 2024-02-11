package taskcontroller

import "github.com/gin-gonic/gin"

type taskController struct{}

func NewTaskController() ITaskController {
	return &taskController{}
}

func (tc *taskController) GetTasks(c *gin.Context) {
}

func (tc *taskController) CreateTask(c *gin.Context) {
}

func (tc *taskController) GetTaskById(c *gin.Context) {
}

func (tc *taskController) UpdateTaskById(c *gin.Context) {
}

func (tc *taskController) DeleteTaskById(c *gin.Context) {
}
