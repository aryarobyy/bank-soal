package repo

import (
	"context"

	"github.com/stretchr/testify/mock"
	"latih.in-be/internal/model"
)

type ExamSessionRepoMock struct {
	mock.Mock
}

func (m *ExamSessionRepoMock) Create(ctx context.Context, e model.ExamSession) (*model.ExamSession, error) {
	args := m.Called(ctx, e)
	return GetReturn[*model.ExamSession](args)
}

func (m *ExamSessionRepoMock) GetById(ctx context.Context, id int) (*model.ExamSession, error) {
	args := m.Called(ctx, id)
	return GetReturn[*model.ExamSession](args)
}

func (m *ExamSessionRepoMock) Update(ctx context.Context, id int, e model.UpdateExamSession) (*model.ExamSession, error) {
	args := m.Called(ctx, id, e)
	return GetReturn[*model.ExamSession](args)
}

func (m *ExamSessionRepoMock) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *ExamSessionRepoMock) GetMany(ctx context.Context, examId int, limit int, offset int) ([]model.ExamSession, int64, error) {
	args := m.Called(ctx, examId, limit, offset)

	var zero []model.ExamSession
	if args.Get(0) == nil {
		return zero, 0, args.Error(2)
	}

	return args.Get(0).([]model.ExamSession), args.Get(1).(int64), args.Error(2)
}

func (m *ExamSessionRepoMock) UpdateCurrNo(ctx context.Context, id int, no model.UpdateCurrNo) (*model.ExamSession, error) {
	args := m.Called(ctx, id, no)
	return GetReturn[*model.ExamSession](args)
}

func (m *ExamSessionRepoMock) FinishExam(ctx context.Context, id int, e model.FinishExam) (*model.ExamSession, error) {
	args := m.Called(ctx, id, e)
	return GetReturn[*model.ExamSession](args)
}

func (m *ExamSessionRepoMock) CheckUserSession(ctx context.Context, userId int, examId int) (*model.ExamSession, error) {
	args := m.Called(ctx, userId, examId)
	return GetReturn[*model.ExamSession](args)
}

func (m *ExamSessionRepoMock) GetScore(ctx context.Context, sessionId int, userId int) (*model.ExamSession, error) {
	args := m.Called(ctx, sessionId, userId)
	return GetReturn[*model.ExamSession](args)
}

func (m *ExamSessionRepoMock) GetUserSession(ctx context.Context, userId int, limit int, offset int) ([]model.ExamSession, int64, error) {
	args := m.Called(ctx, userId, limit, offset)

	var zero []model.ExamSession
	if args.Get(0) == nil {
		return zero, 0, args.Error(2)
	}

	return args.Get(0).([]model.ExamSession), args.Get(1).(int64), args.Error(2)
}
