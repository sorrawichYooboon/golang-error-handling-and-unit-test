package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/sorrawichYooboon/golang-error-handling-and-unit-test/controllers/task_controller"
)

type TaskHttpRoutes struct {
	TaskController controllers.ITaskController
}

func ApplyTaskRoutes(router *gin.Engine, httpRoutes *TaskHttpRoutes) {
	router.GET("/tasks", httpRoutes.TaskController.GetTasks)
	router.POST("/tasks", httpRoutes.TaskController.CreateTask)
	router.GET("/tasks/:id", httpRoutes.TaskController.GetTaskById)
	router.PUT("/tasks/:id", httpRoutes.TaskController.UpdateTaskById)
	router.DELETE("/tasks/:id", httpRoutes.TaskController.DeleteTaskById)
}

// if have subcategory
// router.GET("/tasks/:id/subcategory", httpRoutes.TaskController.GetTaskSubCategory)
// router.POST("/tasks/:id/subcategory", httpRoutes.TaskController.CreateTaskSubCategory)
// router.GET("/tasks/:id/subcategory/:subid", httpRoutes.TaskController.GetTaskSubCategoryById)
// router.PUT("/tasks/:id/subcategory/:subid", httpRoutes.TaskController.UpdateTaskSubCategoryById)
// router.DELETE("/tasks/:id/subcategory/:subid", httpRoutes.TaskController.DeleteTaskSubCategoryById)

// router.GET("/tasks/:id/subcategory/:subid/subsub", httpRoutes.TaskController.GetTaskSubCategorySubSub)
// router.POST("/tasks/:id/subcategory/:subid/subsub", httpRoutes.TaskController.CreateTaskSubCategorySubSub)
// router.GET("/tasks/:id/subcategory/:subid/subsub/:subsubid", httpRoutes.TaskController.GetTaskSubCategorySubSubById)
// router.PUT("/tasks/:id/subcategory/:subid/subsub/:subsubid", httpRoutes.TaskController.UpdateTaskSubCategorySubSubById)
// router.DELETE("/tasks/:id/subcategory/:subid/subsub/:subsubid", httpRoutes.TaskController.DeleteTaskSubCategorySubSubById)
