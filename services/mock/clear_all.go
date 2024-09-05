package mockService

import mock "github.com/stretchr/testify/mock"

func (m *ITaskService) ClearAll() {
	m.Mock = mock.Mock{}
}
