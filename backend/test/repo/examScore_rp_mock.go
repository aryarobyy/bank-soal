package repo

import (
	"context"

	"github.com/stretchr/testify/mock"
	"latih.in-be/internal/model"
)

type ExamScoreRepoMock struct {
	mock.Mock
}

func (m *ExamScoreRepoMock) Create(ctx context.Context, e model.ExamScore) error {
	args := m.Called(ctx, e)
	return args.Error(0)
}

func (m *ExamScoreRepoMock) GetById(ctx context.Context, id int) (*model.ExamScore, error) {
	args := m.Called(ctx, id)
	return GetReturn[*model.ExamScore](args)
}

func (m *ExamScoreRepoMock) GetMany(ctx context.Context, examId int, limit int, offset int) ([]model.ExamScore, error) {
	args := m.Called(ctx, examId, limit, offset)

	var zero []model.ExamScore
	if args.Get(0) == nil {
		return zero, args.Error(1)
	}

	return args.Get(0).([]model.ExamScore), args.Error(1)
}

func (m *ExamScoreRepoMock) Update(ctx context.Context, e model.ExamScore, id int) (*model.ExamScore, error) {
	args := m.Called(ctx, e, id)
	return GetReturn[*model.ExamScore](args)
}

func (m *ExamScoreRepoMock) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *ExamScoreRepoMock) GetByUser(ctx context.Context, userId int, limit int, offset int) ([]model.ExamScore, int64, error) {
	args := m.Called(ctx, userId, limit, offset)

	var zero []model.ExamScore
	if args.Get(0) == nil {
		return zero, 0, args.Error(2)
	}

	return args.Get(0).([]model.ExamScore), args.Get(1).(int64), args.Error(2)
}

func (m *ExamScoreRepoMock) GetSpesificScore(ctx context.Context, userId int, examId int) (*model.ExamScore, error) {
	args := m.Called(ctx, userId, examId)
	return GetReturn[*model.ExamScore](args)
}
