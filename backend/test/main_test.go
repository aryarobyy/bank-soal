package test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"latih.in-be/internal/model"
	"latih.in-be/internal/service"
	"latih.in-be/test/repo"
)

var userRepo = repo.UserRepoMock{Mock: mock.Mock{}}
var questRepo = repo.QuestionRepoMock{Mock: mock.Mock{}}
var optRepo = repo.OptionRepoMock{Mock: mock.Mock{}}
var sessionRepo = repo.ExamSessionRepoMock{Mock: mock.Mock{}}
var examRepo = repo.ExamRepoMock{Mock: mock.Mock{}}
var answerRepo = repo.UserAnswerRepoMock{Mock: mock.Mock{}}
var subjectRepo = repo.SubjectRepoMock{Mock: mock.Mock{}}

var userSrv = service.NewUserService(&userRepo)
var questSrv = service.NewQuestionService(&questRepo, &userRepo, &optRepo, &subjectRepo)
var sessionSrv = service.NewExamSessionService(&sessionRepo, &examRepo, &answerRepo, &questRepo)

func TestRegister(t *testing.T) {
	ctx := context.Background()

	rawPassword := "rahasia123"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)

	nim := "G1A023007"
	nip := "198901182015042004"
	username := "ilhamgoat"

	mockUser := model.User{
		Name:     "Ilham Septina",
		Faculty:  "Teknik",
		Major:    "Informatika",
		Role:     model.RoleUser,
		Nim:      &nim,
		Password: string(hashedPassword),
	}

	mockLecturer := model.User{
		Name:     "Kurniawan setiadi",
		Faculty:  "Teknik",
		Major:    "Informatika",
		Role:     model.RoleLecturer,
		Nip:      &nip,
		Password: string(hashedPassword),
	}

	mockAdmin := model.User{
		Name:     "Kurniawan setiadi",
		Faculty:  "Teknik",
		Major:    "Informatika",
		Role:     model.RoleAdmin,
		Username: &username,
		Password: string(hashedPassword),
	}

	errNotFound := errors.New("record not found")

	tests := []struct {
		name          string
		mockBehavior  func()
		registerCred  model.RegisterCredential
		expectedUser  *model.User
		expectedError bool
		errorContains string
		requesterRole model.Role
	}{
		{
			name: "Success user register (admin)",
			registerCred: model.RegisterCredential{
				Name:     "ilham kurniawan",
				Password: rawPassword,
				Nim:      "Y1G025003",
				Role:     model.RoleUser,
				Major:    "Informatika",
				Faculty:  "Teknik",
				Email:    "mhs@test.com",
			},
			mockBehavior: func() {
				userRepo.Mock.On("GetByEmail", mock.Anything, "mhs@test.com").
					Return(nil, errNotFound).Once()

				userRepo.Mock.On("GetByUsn", mock.Anything, "").
					Return(nil, errNotFound).Once()

				userRepo.Mock.On("GetByNim", mock.Anything, "Y1G025003").
					Return(nil, errNotFound).Once()

				userRepo.Mock.On("Register", mock.Anything, mock.MatchedBy(func(user model.User) bool {
					return user.Name == "ilham kurniawan" && user.Role == model.RoleUser
				})).
					Return(&mockUser, nil).Once()
			},
			expectedUser:  &mockUser,
			errorContains: "Register successess",
			expectedError: false,
			requesterRole: model.RoleAdmin,
		},
		{
			name: "Success admin register (Super admin)",
			registerCred: model.RegisterCredential{
				Name:     "ilham kurniawan",
				Password: rawPassword,
				Username: "Ilhamisthegoat",
				Role:     model.RoleAdmin,
				Major:    "Informatika",
				Faculty:  "Teknik",
				Email:    "admin@test.com",
			},
			mockBehavior: func() {
				userRepo.Mock.On("GetByEmail", mock.Anything, "admin@test.com").
					Return((*model.User)(nil), errNotFound).Once()

				userRepo.Mock.On("GetByUsn", mock.Anything, "Ilhamisthegoat").
					Return((*model.User)(nil), errNotFound).Once()

				userRepo.Mock.On("Register", mock.Anything, mock.MatchedBy(func(user model.User) bool {
					return user.Name == "ilham kurniawan" && user.Role == model.RoleAdmin
				})).
					Return(&mockAdmin, nil).Once()
			},
			expectedUser:  &mockAdmin,
			errorContains: "Register success",
			expectedError: false,
			requesterRole: model.RoleSuperAdmin,
		},
		{
			name: "Success lecturer register (admin)",
			registerCred: model.RegisterCredential{
				Name:     "ilham kurniawan",
				Password: rawPassword,
				Nip:      "198205122008011002",
				Role:     model.RoleLecturer,
				Major:    "Informatika",
				Faculty:  "Teknik",
				Email:    "dosen@test.com",
			},
			mockBehavior: func() {
				userRepo.Mock.On("GetByEmail", mock.Anything, "dosen@test.com").
					Return((*model.User)(nil), errNotFound).Once()

				userRepo.Mock.On("GetByUsn", mock.Anything, "").
					Return(nil, errNotFound).Once()

				userRepo.Mock.On("GetByNip", mock.Anything, "198205122008011002").
					Return((*model.User)(nil), errNotFound).Once()

				userRepo.Mock.On("Register", mock.Anything, mock.MatchedBy(func(user model.User) bool {
					return user.Name == "ilham kurniawan" && user.Role == model.RoleLecturer
				})).
					Return(&mockLecturer, nil).Once()
			},
			expectedUser:  &mockLecturer,
			errorContains: "Register success",
			expectedError: false,
			requesterRole: model.RoleAdmin,
		},
		{
			name: "Failed lecturer register (Missing NIP)",
			registerCred: model.RegisterCredential{
				Name:     "ilham kurniawan",
				Password: rawPassword,
				Nip:      "",
				Role:     model.RoleLecturer,
				Major:    "Informatika",
				Faculty:  "Teknik",
				Email:    "fail@test.com",
			},
			mockBehavior: func() {
				userRepo.Mock.On("GetByEmail", mock.Anything, "fail@test.com").
					Return((*model.User)(nil), errNotFound).Once()

				userRepo.Mock.On("GetByUsn", mock.Anything, "").
					Return(nil, errNotFound).Once()
			},
			expectedUser:  nil,
			errorContains: "lecturer must have NIP",
			expectedError: true,
			requesterRole: model.RoleAdmin,
		},
		{
			name: "Failed admin create another admin (admin)",
			registerCred: model.RegisterCredential{
				Name:     "Calon Admin",
				Password: rawPassword,
				Username: "wannabeadmin",
				Role:     model.RoleAdmin,
				Email:    "hack@test.com",
				Major:    "Informatika",
				Faculty:  "Teknik",
			},
			mockBehavior: func() {
				userRepo.Mock.On("GetByEmail", mock.Anything, "hack@test.com").
					Return((*model.User)(nil), errNotFound).Once()

				userRepo.Mock.On("GetByUsn", mock.Anything, "wannabeadmin").
					Return((*model.User)(nil), errNotFound).Once()
			},
			expectedUser:  nil,
			errorContains: "you cant access this role",
			expectedError: true,
			requesterRole: model.RoleAdmin,
		},
		{
			name: "Failed super admin create user",
			registerCred: model.RegisterCredential{
				Name:     "Mahasiswa Biasa",
				Password: rawPassword,
				Nim:      "Y1G025999",
				Role:     model.RoleUser,
				Email:    "wrong@test.com",
				Major:    "Informatika",
				Faculty:  "Teknik",
			},
			mockBehavior: func() {
				userRepo.Mock.On("GetByEmail", mock.Anything, "wrong@test.com").
					Return((*model.User)(nil), errNotFound).Once()

				userRepo.Mock.On("GetByUsn", mock.Anything, "").
					Return(nil, errNotFound).Once()

				userRepo.Mock.On("GetByNim", mock.Anything, "Y1G025999").
					Return(nil, errNotFound).Once()

				userRepo.Mock.On("Register", mock.Anything, mock.MatchedBy(func(user model.User) bool {
					return user.Name == "Mahasiswa Biasa" && user.Role == model.RoleUser
				})).
					Return(&mockUser, nil).Once()
			},
			expectedUser:  &mockUser,
			errorContains: "super admin cannot create non admin",
			expectedError: true,
			requesterRole: model.RoleSuperAdmin,
		},
		{
			name: "Failed super admin create lecturer",
			registerCred: model.RegisterCredential{
				Name:     "Dosen Baru",
				Password: rawPassword,
				Nip:      "198205122008011005",
				Role:     model.RoleLecturer,
				Email:    "wrong_lec@test.com",
				Major:    "Informatika",
				Faculty:  "Teknik",
			},
			mockBehavior: func() {
				userRepo.Mock.On("GetByEmail", mock.Anything, "wrong_lec@test.com").
					Return((*model.User)(nil), errNotFound).Once()

				userRepo.Mock.On("GetByUsn", mock.Anything, "").
					Return(nil, errNotFound).Once()

				userRepo.Mock.On("GetByNip", mock.Anything, "198205122008011005").
					Return(nil, errNotFound).Once()

				userRepo.Mock.On("Register", mock.Anything, mock.MatchedBy(func(user model.User) bool {
					return user.Name == "Dosen Baru" && user.Role == model.RoleLecturer
				})).
					Return(&mockLecturer, nil).Once()
			},
			expectedUser:  &mockLecturer,
			errorContains: "super admin cannot create non admin",
			expectedError: true,
			requesterRole: model.RoleSuperAdmin,
		},
		{
			name: "Failed lecturer register user",
			registerCred: model.RegisterCredential{
				Name:     "Mahasiswa Bimbingan",
				Password: rawPassword,
				Nim:      "Y1G025005",
				Role:     model.RoleUser,
				Email:    "dosen_iseng@test.com",
				Major:    "Informatika",
				Faculty:  "Teknik",
			},
			mockBehavior: func() {
				userRepo.Mock.On("GetByEmail", mock.Anything, "dosen_iseng@test.com").
					Return((*model.User)(nil), errNotFound).Once()

				userRepo.Mock.On("GetByUsn", mock.Anything, "").
					Return(nil, errNotFound).Once()

				userRepo.Mock.On("GetByNim", mock.Anything, "Y1G025005").
					Return(nil, errNotFound).Once()

				userRepo.Mock.On("Register", mock.Anything, mock.MatchedBy(func(user model.User) bool {
					return user.Name == "Mahasiswa Bimbingan" && user.Role == model.RoleUser
				})).
					Return(&mockUser, nil).Once()
			},
			expectedUser:  &mockUser,
			errorContains: "",
			expectedError: false,
			requesterRole: model.RoleLecturer,
		},
		{
			name: "Failed to add user (wrong nim format)",
			registerCred: model.RegisterCredential{
				Name:     "Teman Mahasiswa",
				Password: rawPassword,
				Nim:      "HUSS7SN8976",
				Role:     model.RoleUser,
				Email:    "intruder@test.com",
				Major:    "Informatika",
				Faculty:  "Teknik",
			},
			mockBehavior: func() {
				userRepo.Mock.On("GetByEmail", mock.Anything, "intruder@test.com").
					Return((*model.User)(nil), errNotFound).Once()

				userRepo.Mock.On("GetByUsn", mock.Anything, "").
					Return(nil, errNotFound).Once()

				userRepo.Mock.On("GetByNim", mock.Anything, "HUSS7SN8976").
					Return(nil, errNotFound).Once()

				userRepo.Mock.On("Register", mock.Anything, mock.MatchedBy(func(user model.User) bool {
					return user.Name == "Teman Mahasiswa" && user.Role == model.RoleUser
				})).
					Return(&mockUser, nil).Once()
			},
			expectedUser:  &mockUser,
			errorContains: "invalid nim format",
			expectedError: true,
			requesterRole: model.RoleUser,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockBehavior != nil {
				tt.mockBehavior()
			}

			err := userSrv.Register(ctx, tt.registerCred, tt.requesterRole)

			if tt.expectedError {
				if assert.Error(t, err) {
					if tt.errorContains != "" {
						assert.Contains(t, err.Error(), tt.errorContains)
					}
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
func TestQuestion(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name          string
		mockBehavior  func()
		createCred    model.Question
		expectedError bool
		errorContains string
	}{
		{
			name: "Create question success",
			createCred: model.Question{
				SubjectId:    2,
				CreatorId:    1,
				QuestionText: "Assalamu alaikum",
				Difficulty:   model.DifficultyHard,
				Options: []model.Option{
					{OptionLabel: "A", OptionText: "sdsadksasdms", IsCorrect: false},
					{OptionLabel: "B", OptionText: "Naik 2 sadsadksadksajlipat", IsCorrect: true},
					{OptionLabel: "C", OptionText: "Naik lebih dari 2 kali lipat", IsCorrect: false},
					{OptionLabel: "D", OptionText: "Turun setengahnya", IsCorrect: false},
				},
			},
			mockBehavior: func() {
				questRepo.Mock.On("CreateWithOptions", mock.Anything, mock.MatchedBy(
					func(q model.Question) bool {
						return q.SubjectId == 2 &&
							q.CreatorId == 1 &&
							q.QuestionText == "Assalamu alaikum" &&
							q.Difficulty == model.DifficultyHard &&
							len(q.Options) == 4
					},
				)).Return(nil).Once()
			},
			expectedError: false,
		},
		{
			name: "Create question failed, duplicate label",
			createCred: model.Question{
				SubjectId:    2,
				CreatorId:    1,
				QuestionText: "Another question",
				Difficulty:   model.DifficultyEasy,
				Options: []model.Option{
					{OptionLabel: "A", OptionText: "option A", IsCorrect: true},
					{OptionLabel: "B", OptionText: "option B", IsCorrect: false},
					{OptionLabel: "B", OptionText: "option C", IsCorrect: true},
					{OptionLabel: "D", OptionText: "option D", IsCorrect: false},
				},
			},
			mockBehavior: func() {
			},
			expectedError: true,
			errorContains: "duplicate option label",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			err := questSrv.CreateWithOptions(ctx, tt.createCred)

			if tt.expectedError {
				assert.Error(t, err)
				if tt.errorContains != "" {
					assert.Contains(t, err.Error(), tt.errorContains)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestStartExam(t *testing.T) {
	ctx := context.Background()
	now := time.Now()
	examStartTime := now.Add(-1 * time.Hour)
	examEndTime := now.Add(1 * time.Hour)

	mockSession := model.ExamSession{
		Id:        1,
		UserId:    2,
		ExamId:    1,
		StartedAt: now,
		Status:    model.SessionInProgress,
	}

	mockExam := model.Exam{
		Id:         1,
		Title:      "Test Exam",
		Difficulty: model.DifficultyEasy,
		LongTime:   60,
		CreatorId:  1,
		StartedAt:  &examStartTime,
		FinishedAt: &examEndTime,
		Score:      100,
	}

	tests := []struct {
		name          string
		mockBehavior  func()
		startCred     model.ExamSession
		userId        int
		examId        int
		expectedError bool
		errorContains string
	}{
		{
			name: "Success user valid",
			startCred: model.ExamSession{
				Status: model.SessionInProgress,
			},
			userId: 2,
			examId: 1,
			mockBehavior: func() {
				sessionRepo.Mock.On("CheckUserSession", ctx, 2, 1).Return(nil, gorm.ErrRecordNotFound).Once()
				examRepo.Mock.On("GetById", ctx, 1).Return(&mockExam, nil).Once()
				sessionRepo.Mock.On("Create", ctx, mock.MatchedBy(func(s model.ExamSession) bool {
					return s.UserId == 2 && s.ExamId == 1 && s.Status == model.SessionInProgress
				})).Return(&mockSession, nil).Once()
			},
			expectedError: false,
		},
		{
			name: "Success check user session",
			startCred: model.ExamSession{
				Status: model.SessionInProgress,
			},
			userId: 2,
			examId: 1,
			mockBehavior: func() {
				sessionRepo.Mock.On("CheckUserSession", ctx, 2, 1).Return(&mockSession, nil).Once()
			},
			expectedError: false,
		},
		{
			name: "Success Check Question",
			startCred: model.ExamSession{
				Status: model.SessionInProgress,
			},
			userId: 2,
			examId: 1,
			mockBehavior: func() {
				sessionRepo.Mock.On("CheckUserSession", ctx, 2, 1).Return(nil, gorm.ErrRecordNotFound).Once()
				examRepo.Mock.On("GetById", ctx, 1).Return(&mockExam, nil).Once()
				sessionRepo.Mock.On("Create", ctx, mock.MatchedBy(func(s model.ExamSession) bool {
					return s.UserId == 2 && s.ExamId == 1 && s.Status == model.SessionInProgress
				})).Return(&mockSession, nil).Once()
			},
			expectedError: false,
		},
		{
			name: "Exam doesn't exist",
			startCred: model.ExamSession{
				Status: model.SessionInProgress,
			},
			userId: 2,
			examId: 999,
			mockBehavior: func() {
				sessionRepo.Mock.On("CheckUserSession", ctx, 2, 999).Return(nil, gorm.ErrRecordNotFound).Once()
				examRepo.Mock.On("GetById", ctx, 999).Return((*model.Exam)(nil), errors.New("exam not found")).Once()
			},
			expectedError: true,
			errorContains: "exam didnt exist",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			session, err := sessionSrv.Create(ctx, tt.startCred, tt.userId, tt.examId)

			if tt.expectedError {
				assert.Error(t, err)
				if tt.errorContains != "" {
					assert.Contains(t, err.Error(), tt.errorContains)
				}
			} else {
				assert.NoError(t, err)
				if tt.name == "Check user session" {
					assert.Equal(t, &mockSession, session)
				} else {
					assert.NotNil(t, session)
					assert.Equal(t, tt.userId, session.UserId)
					assert.Equal(t, tt.examId, session.ExamId)
					assert.Equal(t, model.SessionInProgress, session.Status)
				}
			}
		})
	}
}

func TestFinishSession(t *testing.T) {
	ctx := context.Background()

	now := time.Now()
	sessionStartTime := now.Add(-30 * time.Minute)
	mockSession1 := model.ExamSession{
		Id:        1,
		UserId:    2,
		ExamId:    1,
		StartedAt: sessionStartTime,
		Status:    model.SessionInProgress,
	}

	mockSession2 := model.ExamSession{
		Id:        2,
		UserId:    2,
		ExamId:    2,
		StartedAt: sessionStartTime,
		Status:    model.SessionInProgress,
	}

	mockSession3 := model.ExamSession{
		Id:        3,
		UserId:    2,
		ExamId:    3,
		StartedAt: sessionStartTime,
		Status:    model.SessionInProgress,
	}

	mockUserAnswer1 := []model.UserAnswer{
		{
			Id:            1,
			UserId:        2,
			QuestionId:    1,
			ExamSessionId: 1,
			Answer:        "Option A",
			IsCorrect:     true,
		},
	}

	mockUserAnswer2 := []model.UserAnswer{
		{
			Id:            2,
			UserId:        2,
			QuestionId:    2,
			ExamSessionId: 3,
			Answer:        "Option B",
			IsCorrect:     true,
		},
	}

	mockQuestion1 := model.Question{
		Id:           1,
		QuestionText: "Test Question 1",
		Difficulty:   model.DifficultyEasy,
		Score:        5,
		CreatorId:    1,
		SubjectId:    1,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	mockQuestion2 := model.Question{
		Id:           2,
		QuestionText: "Test Question 2",
		Difficulty:   model.DifficultyMedium,
		Score:        10,
		CreatorId:    1,
		SubjectId:    1,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	tests := []struct {
		name          string
		mockBehavior  func()
		userId        int
		sessionId     int
		expectedError bool
		errorContains string
	}{
		{
			name:      "Finished exam",
			userId:    2,
			sessionId: 1,
			mockBehavior: func() {
				sessionRepo.Mock.On("GetById", ctx, 1).Return(&mockSession1, nil).Once()
				answerRepo.Mock.On("GetAllUserAnswers", ctx, 2, 1).Return(mockUserAnswer1, nil).Once()

				questRepo.Mock.On("GetByExamId", ctx, 1).Return([]model.Question{mockQuestion1}, nil).Once()

				questRepo.Mock.On("GetById", ctx, 1).Return(&mockQuestion1, nil).Twice()

				sessionRepo.Mock.On("FinishExam", ctx, 1, mock.MatchedBy(func(f model.FinishExam) bool {
					return f.Status == model.SessionFinished
				})).Return(&model.ExamSession{
					Id:         1,
					UserId:     2,
					ExamId:     1,
					FinishedAt: &now,
					Status:     model.SessionFinished,
					Score:      5,
					MaxScore:   5,
					Percentage: 100.0,
					IsPassed:   true,
				}, nil).Once()
			},
			expectedError: false,
		},
		{
			name:      "Check session",
			userId:    2,
			sessionId: 10,
			mockBehavior: func() {
				sessionRepo.Mock.On("GetById", ctx, 10).Return((*model.ExamSession)(nil), errors.New("session not found")).Once()
			},
			expectedError: true,
			errorContains: "failed to find session",
		},
		{
			name:      "Check exam time",
			userId:    2,
			sessionId: 2,
			mockBehavior: func() {
				sessionRepo.Mock.On("GetById", ctx, 2).Return(&mockSession2, nil).Once()

				answerRepo.Mock.On("GetAllUserAnswers", ctx, 2, 2).Return([]model.UserAnswer(nil), errors.New("failed to get answers")).Once()
			},
			expectedError: true,
			errorContains: "failed to calculate score",
		},
		{
			name:      "Submit answers",
			userId:    2,
			sessionId: 3,
			mockBehavior: func() {
				sessionRepo.Mock.On("GetById", ctx, 3).Return(&mockSession3, nil).Once()
				answerRepo.Mock.On("GetAllUserAnswers", ctx, 2, 3).Return(mockUserAnswer2, nil).Once()
				questRepo.Mock.On("GetByExamId", ctx, 3).Return([]model.Question{mockQuestion2}, nil).Once()
				questRepo.Mock.On("GetById", ctx, 2).Return(&mockQuestion2, nil).Twice()
				sessionRepo.Mock.On("FinishExam", ctx, 3, mock.MatchedBy(func(f model.FinishExam) bool {
					return f.Status == model.SessionFinished
				})).Return(&model.ExamSession{
					Id:         3,
					UserId:     2,
					ExamId:     3,
					FinishedAt: &now,
					Status:     model.SessionFinished,
					Score:      10,
					MaxScore:   10,
					Percentage: 100.0,
					IsPassed:   true,
				}, nil).Once()
			},
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			session, err := sessionSrv.FinishExam(ctx, tt.userId, tt.sessionId)

			if tt.expectedError {
				assert.Error(t, err)
				if tt.errorContains != "" {
					assert.Contains(t, err.Error(), tt.errorContains)
				}
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, session)
				assert.Equal(t, model.SessionFinished, session.Status)
			}
		})
	}
}
