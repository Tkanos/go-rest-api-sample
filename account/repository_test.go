package account

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type mockedAccountRepository struct {
	mock.Mock
}

func (m *mockedAccountRepository) GetAccount(ctx context.Context, id string) (*Account, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Account), args.Error(1)
}

func (m *mockedAccountRepository) GetAccounts(ctx context.Context, filter Filter, pagination Pagination) ([]*Account, error) {
	args := m.Called(filter, pagination)
	return args.Get(0).([]*Account), args.Error(1)
}

func (m *mockedAccountRepository) UpdateAccount(ctx context.Context, a Account) error {
	args := m.Called(a)
	return args.Error(0)
}

func (m *mockedAccountRepository) CreateAccount(ctx context.Context, a Account) (string, error) {
	args := m.Called(a)
	return args.Get(0).(string), args.Error(1)
}

func (m *mockedAccountRepository) DeleteAccount(ctx context.Context, id string) error {
	args := m.Called(id)
	return args.Error(0)
}
