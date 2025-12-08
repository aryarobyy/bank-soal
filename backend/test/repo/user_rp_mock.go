package repo

import (
	"context"

	"github.com/stretchr/testify/mock"
	"latih.in-be/internal/model"
)

type UserRepoMock struct {
	mock.Mock
}

func (m *UserRepoMock) Register(ctx context.Context, user model.User) (*model.User, error) {
	args := m.Called(ctx, user)
	return GetReturn[*model.User](args)
}

func (m *UserRepoMock) GetById(ctx context.Context, id int) (*model.User, error) {
	args := m.Called(ctx, id)
	return GetReturn[*model.User](args)
}

func (m *UserRepoMock) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	args := m.Called(ctx, email)
	return GetReturn[*model.User](args)
}

func (m *UserRepoMock) GetByNim(ctx context.Context, nim string) (*model.User, error) {
	args := m.Called(ctx, nim)
	return GetReturn[*model.User](args)
}

func (m *UserRepoMock) GetByNip(ctx context.Context, nip string) (*model.User, error) {
	args := m.Called(ctx, nip)
	return GetReturn[*model.User](args)
}

func (m *UserRepoMock) GetByUsn(ctx context.Context, username string) (*model.User, error) {
	args := m.Called(ctx, username)
	return GetReturn[*model.User](args)
}

func (m *UserRepoMock) GetByName(ctx context.Context, name string, limit int, offset int) ([]model.User, int64, error) {
	args := m.Called(ctx, name, limit, offset)

	var zero []model.User
	if args.Get(0) == nil {
		return zero, 0, args.Error(2)
	}

	return args.Get(0).([]model.User), args.Get(1).(int64), args.Error(2)
}

func (m *UserRepoMock) GetByRole(ctx context.Context, role model.Role, limit int, offset int) ([]model.User, int64, error) {
	args := m.Called(ctx, role, limit, offset)

	var zero []model.User
	if args.Get(0) == nil {
		return zero, 0, args.Error(2)
	}

	return args.Get(0).([]model.User), args.Get(1).(int64), args.Error(2)
}

func (m *UserRepoMock) Update(ctx context.Context, user model.UpdateUser, id int) (*model.User, error) {
	args := m.Called(ctx, user, id)
	return GetReturn[*model.User](args)
}

func (m *UserRepoMock) GetMany(ctx context.Context, limit int, offset int) ([]model.User, int64, error) {
	args := m.Called(ctx, limit, offset)

	var zero []model.User
	if args.Get(0) == nil {
		return zero, 0, args.Error(2)
	}

	return args.Get(0).([]model.User), args.Get(1).(int64), args.Error(2)
}

func (m *UserRepoMock) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *UserRepoMock) ChangePassword(ctx context.Context, id int, password string) error {
	args := m.Called(ctx, id, password)
	return args.Error(0)
}

func (m *UserRepoMock) ChangeRole(ctx context.Context, id int, role model.Role) error {
	args := m.Called(ctx, id, role)
	return args.Error(0)
}

func (m *UserRepoMock) BulkInsert(ctx context.Context, users []model.User) ([]model.User, error) {
	args := m.Called(ctx, users)
	return GetReturn[[]model.User](args)
}

type TokenGenerator interface {
	Generate(user *model.User) (string, error)
}
