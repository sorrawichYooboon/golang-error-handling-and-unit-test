package taskrepository

type Task struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ITaskRepository interface {
	GetTasks() ([]Task, error)
	CreateTask(task Task) (Task, error)
	GetTaskById(id string) (Task, error)
	UpdateTaskById(id string, task Task) (Task, error)
	DeleteTaskById(id string) error
}
