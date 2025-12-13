package repo

import (
	"context"

	"github.com/stretchr/testify/mock"
	"latih.in-be/internal/model"
)

type SubjectRepoMock struct {
	mock.Mock
}

func (m *SubjectRepoMock) Create(ctx context.Context, subject model.Subject) error {
	args := m.Called(ctx, subject)
	return args.Error(0)
}

func (m *SubjectRepoMock) GetById(ctx context.Context, id int) (*model.Subject, error) {
	args := m.Called(ctx, id)
	return GetReturn[*model.Subject](args)
}

func (m *SubjectRepoMock) GetByCode(ctx context.Context, code string) (*model.Subject, error) {
	args := m.Called(ctx, code)
	return GetReturn[*model.Subject](args)
}

// func (m *SubjectRepoMock) GetByTitle(ctx context.Context, title model.SubjectTitle) (*model.Subject, error) {
// 	args := m.Called(ctx, title)
// 	return GetReturn[*model.Subject](args)
// }

func (m *SubjectRepoMock) Update(ctx context.Context, subject model.Subject, id int) (*model.Subject, error) {
	args := m.Called(ctx, subject, id)
	return GetReturn[*model.Subject](args)
}

func (m *SubjectRepoMock) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *SubjectRepoMock) GetMany(ctx context.Context, limit int, offset int) ([]model.Subject, int64, error) {
	args := m.Called(ctx, limit, offset)

	var zero []model.Subject
	if args.Get(0) == nil {
		return zero, 0, args.Error(2)
	}

	return args.Get(0).([]model.Subject), args.Get(1).(int64), args.Error(2)
}

// func (m *SubjectRepoMock) GetByTitleContains(ctx context.Context, title string, limit int, offset int) ([]model.Subject, int64, error) {
// 	args := m.Called(ctx, title, limit, offset)

// 	var zero []model.Subject
// 	if args.Get(0) == nil {
// 		return zero, 0, args.Error(2)
// 	}

// 	return args.Get(0).([]model.Subject), args.Get(1).(int64), args.Error(2)
// }
