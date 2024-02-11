package taskrepository

type taskRepository struct{}

func NewTaskRepository() ITaskRepository {
	return &taskRepository{}
}

func (tr *taskRepository) GetTasks() ([]Task, error) {
	return nil, nil
}

func (tr *taskRepository) CreateTask(task Task) (Task, error) {
	return Task{}, nil
}

func (tr *taskRepository) GetTaskById(id string) (Task, error) {
	return Task{}, nil
}

func (tr *taskRepository) UpdateTaskById(id string, task Task) (Task, error) {
	return Task{}, nil
}

func (tr *taskRepository) DeleteTaskById(id string) error {
	return nil
}
