package repo

import (
	"context"

	"github.com/stretchr/testify/mock"
	"latih.in-be/internal/model"
)

type QuestionRepoMock struct {
	mock.Mock
}

func (m *QuestionRepoMock) Create(ctx context.Context, q *model.Question) error {
	args := m.Called(ctx, q)
	return args.Error(0)
}

func (m *QuestionRepoMock) GetById(ctx context.Context, id int) (*model.Question, error) {
	args := m.Called(ctx, id)
	return GetReturn[*model.Question](args)
}

func (m *QuestionRepoMock) GetMany(ctx context.Context, limit int, offset int) ([]model.Question, int64, error) {
	args := m.Called(ctx, limit, offset)

	var zero []model.Question
	if args.Get(0) == nil {
		return zero, 0, args.Error(2)
	}

	return args.Get(0).([]model.Question), args.Get(1).(int64), args.Error(2)
}

func (m *QuestionRepoMock) GetByExam(ctx context.Context, examId int, limit int, offset int) ([]model.Question, int64, error) {
	args := m.Called(ctx, examId, limit, offset)

	var zero []model.Question
	if args.Get(0) == nil {
		return zero, 0, args.Error(2)
	}

	return args.Get(0).([]model.Question), args.Get(1).(int64), args.Error(2)
}

func (m *QuestionRepoMock) Update(ctx context.Context, q model.UpdateQuestion, id int) (*model.Question, error) {
	args := m.Called(ctx, q, id)
	return GetReturn[*model.Question](args)
}

func (m *QuestionRepoMock) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *QuestionRepoMock) CreateWithOptions(ctx context.Context, question model.Question) error {
	args := m.Called(ctx, question)
	return args.Error(0)
}

func (m *QuestionRepoMock) CreateBatch(ctx context.Context, q []model.Question) error {
	args := m.Called(ctx, q)
	return args.Error(0)
}

func (m *QuestionRepoMock) GetByDifficult(ctx context.Context, diff string, limit int, offset int) ([]model.Question, int64, error) {
	args := m.Called(ctx, diff, limit, offset)

	var zero []model.Question
	if args.Get(0) == nil {
		return zero, 0, args.Error(2)
	}

	return args.Get(0).([]model.Question), args.Get(1).(int64), args.Error(2)
}

func (m *QuestionRepoMock) GetByCreatorId(ctx context.Context, creatorId int, limit int, offset int) ([]model.Question, int64, error) {
	args := m.Called(ctx, creatorId, limit, offset)

	var zero []model.Question
	if args.Get(0) == nil {
		return zero, 0, args.Error(2)
	}

	return args.Get(0).([]model.Question), args.Get(1).(int64), args.Error(2)
}

func (m *QuestionRepoMock) GetBySubject(ctx context.Context, subjectId int, limit int, offset int) ([]model.Question, int64, error) {
	args := m.Called(ctx, subjectId, limit, offset)

	var zero []model.Question
	if args.Get(0) == nil {
		return zero, 0, args.Error(2)
	}

	return args.Get(0).([]model.Question), args.Get(1).(int64), args.Error(2)
}

func (m *QuestionRepoMock) GetByExamId(ctx context.Context, examId int) ([]model.Question, error) {
	args := m.Called(ctx, examId)

	var zero []model.Question
	if args.Get(0) == nil {
		return zero, args.Error(1)
	}

	return args.Get(0).([]model.Question), args.Error(1)
}
