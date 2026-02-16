package repo

import (
	"context"

	"github.com/stretchr/testify/mock"
	"latih.in-be/internal/model"
)

type ExamRepoMock struct {
	mock.Mock
}

func (m *ExamRepoMock) Create(ctx context.Context, req model.CreateExam) error {
	args := m.Called(ctx, req)
	return args.Error(0)
}

func (m *ExamRepoMock) GetById(ctx context.Context, id int) (*model.Exam, error) {
	args := m.Called(ctx, id)
	return GetReturn[*model.Exam](args)
}

func (m *ExamRepoMock) Update(ctx context.Context, e model.Exam, id int) (*model.Exam, error) {
	args := m.Called(ctx, e, id)
	return GetReturn[*model.Exam](args)
}

func (m *ExamRepoMock) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *ExamRepoMock) GetMany(ctx context.Context, limit int, offset int) ([]model.Exam, int64, error) {
	args := m.Called(ctx, limit, offset)

	var zero []model.Exam
	if args.Get(0) == nil {
		return zero, 0, args.Error(2)
	}

	return args.Get(0).([]model.Exam), args.Get(1).(int64), args.Error(2)
}

func (m *ExamRepoMock) StartSession(ctx context.Context, id int) (*model.Exam, error) {
	args := m.Called(ctx, id)
	return GetReturn[*model.Exam](args)
}

func (m *ExamRepoMock) UpdateScore(ctx context.Context, examId int, score int) error {
	args := m.Called(ctx, examId, score)
	return args.Error(0)
}

func (m *ExamRepoMock) AddQuestions(ctx context.Context, examId int, questionIds []int) error {
	args := m.Called(ctx, examId, questionIds)
	return args.Error(0)
}

func (m *ExamRepoMock) ReplaceQuestions(ctx context.Context, examId int, questionIds []int) error {
	args := m.Called(ctx, examId, questionIds)
	return args.Error(0)
}

func (m *ExamRepoMock) RemoveQuestions(ctx context.Context, examId int, questionIds []int) error {
	args := m.Called(ctx, examId, questionIds)
	return args.Error(0)
}

func (m *ExamRepoMock) CheckQuestion(ctx context.Context, examId int, questionId int) (bool, error) {
	args := m.Called(ctx, examId, questionId)

	if args.Get(0) == nil {
		return false, args.Error(1)
	}

	return args.Get(0).(bool), args.Error(1)
}
