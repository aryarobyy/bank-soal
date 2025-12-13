package repo

import (
	"context"

	"github.com/stretchr/testify/mock"
	"latih.in-be/internal/model"
)

type XlsPathRepoMock struct {
	mock.Mock
}

func (m *XlsPathRepoMock) Create(ctx context.Context, o model.XlsPath) error {
	args := m.Called(ctx, o)
	return args.Error(0)
}

func (m *XlsPathRepoMock) GetById(ctx context.Context, id int) (*model.XlsPath, error) {
	args := m.Called(ctx, id)
	return GetReturn[*model.XlsPath](args)
}

func (m *XlsPathRepoMock) GetMany(ctx context.Context, limit int, offset int) ([]model.XlsPath, error) {
	args := m.Called(ctx, limit, offset)

	var zero []model.XlsPath
	if args.Get(0) == nil {
		return zero, args.Error(1)
	}

	return args.Get(0).([]model.XlsPath), args.Error(1)
}

func (m *XlsPathRepoMock) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
