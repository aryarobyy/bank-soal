package repo

import (
	"context"

	"github.com/stretchr/testify/mock"
	"latih.in-be/internal/model"
)

type UserAnswerRepoMock struct {
	mock.Mock
}

func (m *UserAnswerRepoMock) Create(ctx context.Context, userAnswer *model.UserAnswer) error {
	args := m.Called(ctx, userAnswer)
	return args.Error(0)
}

func (m *UserAnswerRepoMock) GetById(ctx context.Context, id int) (*model.UserAnswer, error) {
	args := m.Called(ctx, id)
	return GetReturn[*model.UserAnswer](args)
}

func (m *UserAnswerRepoMock) GetMany(ctx context.Context, limit int, offset int) ([]model.UserAnswer, int64, error) {
	args := m.Called(ctx, limit, offset)

	var zero []model.UserAnswer
	if args.Get(0) == nil {
		return zero, 0, args.Error(2)
	}

	return args.Get(0).([]model.UserAnswer), args.Get(1).(int64), args.Error(2)
}

func (m *UserAnswerRepoMock) Update(ctx context.Context, id int, userAnswer *model.UserAnswer) (*model.UserAnswer, error) {
	args := m.Called(ctx, id, userAnswer)
	return GetReturn[*model.UserAnswer](args)
}

func (m *UserAnswerRepoMock) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *UserAnswerRepoMock) GetByExamSessionId(ctx context.Context, examSessionId int, limit int, offset int) ([]model.UserAnswer, int64, error) {
	args := m.Called(ctx, examSessionId, limit, offset)

	var zero []model.UserAnswer
	if args.Get(0) == nil {
		return zero, 0, args.Error(2)
	}

	return args.Get(0).([]model.UserAnswer), args.Get(1).(int64), args.Error(2)
}

func (m *UserAnswerRepoMock) GetByQuestionId(ctx context.Context, questionId int, limit int, offset int) ([]model.UserAnswer, int64, error) {
	args := m.Called(ctx, questionId, limit, offset)

	var zero []model.UserAnswer
	if args.Get(0) == nil {
		return zero, 0, args.Error(2)
	}

	return args.Get(0).([]model.UserAnswer), args.Get(1).(int64), args.Error(2)
}

func (m *UserAnswerRepoMock) GetUserAnswer(ctx context.Context, userId int, examSessionId int, limit int, offset int) ([]model.UserAnswer, int64, error) {
	args := m.Called(ctx, userId, examSessionId, limit, offset)

	var zero []model.UserAnswer
	if args.Get(0) == nil {
		return zero, 0, args.Error(2)
	}

	return args.Get(0).([]model.UserAnswer), args.Get(1).(int64), args.Error(2)
}

func (m *UserAnswerRepoMock) GetAllUserAnswers(ctx context.Context, userId int, examSessionId int) ([]model.UserAnswer, error) {
	args := m.Called(ctx, userId, examSessionId)

	var zero []model.UserAnswer
	if args.Get(0) == nil {
		return zero, args.Error(1)
	}

	return args.Get(0).([]model.UserAnswer), args.Error(1)
}

func (m *UserAnswerRepoMock) CheckUserAnswer(ctx context.Context, userId int, sessionId int, questionId int) (*model.UserAnswer, error) {
	args := m.Called(ctx, userId, sessionId, questionId)
	return GetReturn[*model.UserAnswer](args)
}
