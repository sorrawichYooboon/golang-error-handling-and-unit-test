package mockRepo

import "github.com/stretchr/testify/mock"

func (m *ITaskRepository) ClearAll() {
	m.Mock = mock.Mock{}
}
