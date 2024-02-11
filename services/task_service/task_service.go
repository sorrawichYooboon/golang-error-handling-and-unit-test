package taskservice

type taskService struct{}

func NewTaskService() ITaskService {
	return &taskService{}
}

func (ts *taskService) GetTasks() ([]Task, error) {
	return nil, nil
}

func (ts *taskService) CreateTask(task Task) (Task, error) {
	return Task{}, nil
}

func (ts *taskService) GetTaskById(id string) (Task, error) {
	return Task{}, nil
}

func (ts *taskService) UpdateTaskById(id string, task Task) (Task, error) {
	return Task{}, nil
}

func (ts *taskService) DeleteTaskById(id string) error {
	return nil
}
