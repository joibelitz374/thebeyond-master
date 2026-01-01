package mocks

import (
	"context"

	"github.com/quickpowered/thebeyond-master/internal/repositories/xraymanager/dto"
	"github.com/stretchr/testify/mock"
)

type MockXrayManagerRepo struct {
	mock.Mock
}

func (m *MockXrayManagerRepo) GetGroupedNodes(region dto.Region) (map[dto.GroupType][]*dto.Node, error) {
	args := m.Called(region)
	return args.Get(0).(map[dto.GroupType][]*dto.Node), args.Error(1)
}

func (m *MockXrayManagerRepo) AddClient(ctx context.Context, region dto.Region, id string, email string) error {
	args := m.Called(ctx, region, id, email)
	return args.Error(0)
}

func (m *MockXrayManagerRepo) RemoveClient(ctx context.Context, region dto.Region, email string) error {
	args := m.Called(ctx, region, email)
	return args.Error(0)
}
