package service

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/xuri/excelize/v2"
	"latih.in-be/internal/model"
	"latih.in-be/internal/repository"
)

type XlsPathService interface {
	SaveXlsPath(ctx context.Context, filepath string) error
	GetById(ctx context.Context, id int) (*model.XlsPath, error)
	Delete(ctx context.Context, id int) error
	GetMany(ctx context.Context, limit int, offset int) ([]model.XlsPath, error)
	ExportUsersToExcel(users []model.BulkUserOutput, storageDir string) (filename string, filepath string, err error)
}

type xlspathService struct {
	repo repository.XlsPathRepository
}

func NewXlsPathService(repo repository.XlsPathRepository) XlsPathService {
	return &xlspathService{
		repo: repo,
	}
}

func (s *xlspathService) SaveXlsPath(ctx context.Context, filepath string) error {
	xls := model.XlsPath{
		FilePath: filepath,
	}

	err := s.repo.Create(ctx, xls)
	if err != nil {
		return fmt.Errorf("failed to create xlspath: %w", err)
	}

	return nil
}

func (s *xlspathService) GetById(ctx context.Context, id int) (*model.XlsPath, error) {
	data, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("data with id %d not found: %w", id, err)
	}
	return data, nil
}

func (s *xlspathService) GetMany(ctx context.Context, limit int, offset int) ([]model.XlsPath, error) {
	data, err := s.repo.GetMany(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get all data: %w", err)
	}
	return data, nil
}

func (s *xlspathService) Delete(ctx context.Context, id int) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete data: %w", err)
	}
	return nil
}

func (s *xlspathService) ExportUsersToExcel(users []model.BulkUserOutput, storageDir string) (filename string, filepath string, err error) {
	f := excelize.NewFile()

	if err := f.SetCellValue("Sheet1", "A1", "NIM"); err != nil {
		return "", "", fmt.Errorf("failed to set header A1: %w", err)
	}
	if err := f.SetCellValue("Sheet1", "B1", "Password"); err != nil {
		return "", "", fmt.Errorf("failed to set header B1: %w", err)
	}

	for i := range users {
		row := strconv.Itoa(i + 2)
		if err := f.SetCellValue("Sheet1", "A"+row, users[i].Nim); err != nil {
			return "", "", fmt.Errorf("failed to set NIM at row %s: %w", row, err)
		}
		if err := f.SetCellValue("Sheet1", "B"+row, users[i].Password); err != nil {
			return "", "", fmt.Errorf("failed to set Password at row %s: %w", row, err)
		}
	}

	if err := os.MkdirAll(storageDir, os.ModePerm); err != nil {
		return "", "", fmt.Errorf("failed to create storage directory: %w", err)
	}

	timestamp := time.Now().Format("20060102_150405")
	filename = fmt.Sprintf("bulk_users_%s.xlsx", timestamp)
	filepath = fmt.Sprintf("%s/%s", storageDir, filename)

	if err := f.SaveAs(filepath); err != nil {
		return "", "", fmt.Errorf("failed to save excel file: %w", err)
	}

	return filename, filepath, nil
}
