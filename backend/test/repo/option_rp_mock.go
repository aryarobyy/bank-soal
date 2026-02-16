package repo

import (
	"context"

	"github.com/stretchr/testify/mock"
	"latih.in-be/internal/model"
)

type OptionRepoMock struct {
	mock.Mock
}

func (m *OptionRepoMock) Create(ctx context.Context, o model.Option) error {
	args := m.Called(ctx, o)
	return args.Error(0)
}

func (m *OptionRepoMock) GetById(ctx context.Context, id int) (*model.Option, error) {
	args := m.Called(ctx, id)
	return GetReturn[*model.Option](args)
}

func (m *OptionRepoMock) GetMany(ctx context.Context, qId int, limit int, offset int) ([]model.Option, error) {
	args := m.Called(ctx, qId, limit, offset)

	var zero []model.Option
	if args.Get(0) == nil {
		return zero, args.Error(1)
	}

	return args.Get(0).([]model.Option), args.Error(1)
}

func (m *OptionRepoMock) Update(ctx context.Context, o model.Option, id int) (*model.Option, error) {
	args := m.Called(ctx, o, id)
	return GetReturn[*model.Option](args)
}

func (m *OptionRepoMock) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *OptionRepoMock) DeleteByQuestionId(ctx context.Context, qId int) error {
	args := m.Called(ctx, qId)
	return args.Error(0)
}

func (m *OptionRepoMock) CheckCorrectAnswer(ctx context.Context, qId int, answer string) (bool, error) {
	args := m.Called(ctx, qId, answer)

	if args.Get(0) == nil {
		return false, args.Error(1)
	}

	return args.Get(0).(bool), args.Error(1)
}
