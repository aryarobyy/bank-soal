package repository

import (
	"context"

	"gorm.io/gorm"
	"latih.in-be/internal/model"
)

type ExamQuestionRepository interface {
	AddQuestionToExam(ctx context.Context, examId int, questionIds []int) error
	UpdateQuestionsInExam(ctx context.Context, examId int, questionIds []int) error
	RemoveQuestionsFromExam(ctx context.Context, examId int, questionIds []int) error
}

type examQuestionRepository struct {
	db *gorm.DB
}

func NewExamQuestionRepository(db *gorm.DB) ExamQuestionRepository {
	return &examQuestionRepository{db: db}
}

func (r *examQuestionRepository) AddQuestionToExam(ctx context.Context, examId int, questionIds []int) error {
	var exam model.Exam
	if err := r.db.WithContext(ctx).
		First(&exam, examId).
		Error; err != nil {
		return err
	}

	var questions []model.Question
	if err := r.db.WithContext(ctx).
		Where("id IN ?", questionIds).
		Find(&questions).
		Error; err != nil {
		return err
	}

	return r.db.WithContext(ctx).
		Model(&exam).
		Association("Questions").
		Append(&questions)
}

func (r *examQuestionRepository) UpdateQuestionsInExam(ctx context.Context, examId int, questionIds []int) error {
	var exam model.Exam
	if err := r.db.WithContext(ctx).
		First(&exam, examId).
		Error; err != nil {
		return err
	}

	var questions []model.Question
	if err := r.db.WithContext(ctx).
		Where("id IN ?", questionIds).
		Find(&questions).
		Error; err != nil {
		return err
	}

	return r.db.WithContext(ctx).
		Model(&exam).
		Association("Questions").
		Replace(&questions)
}

func (r *examQuestionRepository) RemoveQuestionsFromExam(ctx context.Context, examId int, questionIds []int) error {
	var exam model.Exam
	if err := r.db.WithContext(ctx).
		First(&exam, examId).
		Error; err != nil {
		return err
	}

	var questions []model.Question
	if err := r.db.WithContext(ctx).
		Where("id IN ?", questionIds).
		Find(&questions).
		Error; err != nil {
		return err
	}

	return r.db.WithContext(ctx).
		Model(&exam).
		Association("Questions").
		Delete(&questions)
}
