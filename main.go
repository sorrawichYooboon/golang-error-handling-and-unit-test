package main

import (
	"log"

	"github.com/gin-gonic/gin"
	taskcontroller "github.com/sorrawichYooboon/golang-error-handling-and-unit-test/controllers/task_controller"
	"github.com/sorrawichYooboon/golang-error-handling-and-unit-test/internal"
	taskrepository "github.com/sorrawichYooboon/golang-error-handling-and-unit-test/repositories/task_repository"
	routes "github.com/sorrawichYooboon/golang-error-handling-and-unit-test/routes"
	taskservice "github.com/sorrawichYooboon/golang-error-handling-and-unit-test/services/task_service"
)

func main() {
	router := gin.Default()

	tasks, err := internal.ReadTasksFromFile("mock_tasks.json")
	if err != nil {
		log.Fatalf("Failed to read tasks from file: %v", err)
	}

	taskRepo := taskrepository.NewTaskRepository(tasks)
	taskSv := taskservice.NewTaskService(taskRepo)
	taskController := taskcontroller.NewTaskController(taskSv)
	routes.ApplyTaskRoutes(router, &routes.TaskHttpRoutes{
		TaskController: taskController,
	})

	router.Run(":8080")
}
