package test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
	"latih.in-be/internal/model"
	"latih.in-be/internal/service"
	"latih.in-be/test/repo"
)

var userRepo = repo.UserRepoMock{Mock: mock.Mock{}}
var userSrv = service.NewUserService(&userRepo)

var questRepo = repo.QuestionRepoMock{Mock: mock.Mock{}}
var optRepo = repo.OptionRepoMock{Mock: mock.Mock{}}
var questService = service.NewQuestionService(&questRepo, &userRepo, &optRepo)

func TestRegister(t *testing.T) {
	ctx := context.Background()

	rawPassword := "rahasia123"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)

	mockUser := model.User{
		Name:     "Ilham Septina",
		Faculty:  "Teknik",
		Major:    "Informatika",
		Role:     model.RoleUser,
		Password: string(hashedPassword),
	}

	mockLecturer := model.User{
		Name:     "Kurniawan setiadi",
		Faculty:  "Teknik",
		Major:    "Informatika",
		Role:     model.RoleLecturer,
		Password: string(hashedPassword),
	}

	mockAdmin := model.User{
		Name:     "Kurniawan setiadi",
		Faculty:  "Teknik",
		Major:    "Informatika",
		Role:     model.RoleAdmin,
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
			name: "Success user register",
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
			name: "Success admin register",
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
			name: "Success lecturer register",
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
				Nip:      "", // NIP is empty
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
			name: "Failed admin create another admin",
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
			name: "Success super admin create user",
			registerCred: model.RegisterCredential{
				Name:     "Mahasiswa Biasa",
				Password: rawPassword,
				Nim:      "Y1G025999", // Required for user role
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
			errorContains: "",
			expectedError: false,
			requesterRole: model.RoleSuperAdmin,
		},
		{
			name: "Success super admin create lecturer",
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
			errorContains: "",
			expectedError: false,
			requesterRole: model.RoleSuperAdmin,
		},
		{
			name: "Success user register another user",
			registerCred: model.RegisterCredential{
				Name:     "Teman Mahasiswa",
				Password: rawPassword,
				Nim:      "Y1G025004",
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

				userRepo.Mock.On("GetByNim", mock.Anything, "Y1G025004").
					Return(nil, errNotFound).Once()

				userRepo.Mock.On("Register", mock.Anything, mock.MatchedBy(func(user model.User) bool {
					return user.Name == "Teman Mahasiswa" && user.Role == model.RoleUser
				})).
					Return(&mockUser, nil).Once()
			},
			expectedUser:  &mockUser,
			errorContains: "",
			expectedError: false,
			requesterRole: model.RoleUser,
		},
		{
			name: "Success lecturer register user",
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			err := userSrv.Register(ctx, tt.registerCred, tt.requesterRole)

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

func TestCreateQ(t *testing.T) {
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
			name: "Create question failure",
			createCred: model.Question{
				SubjectId:    2,
				CreatorId:    1,
				QuestionText: "Another question",
				Difficulty:   model.DifficultyEasy,
				Options: []model.Option{
					{OptionLabel: "A", OptionText: "option A", IsCorrect: true},
					{OptionLabel: "B", OptionText: "option B", IsCorrect: false},
				},
			},
			mockBehavior: func() {
				questRepo.Mock.On("CreateWithOptions", mock.Anything, mock.MatchedBy(
					func(q model.Question) bool {
						return q.SubjectId == 2 &&
							q.CreatorId == 1 &&
							q.QuestionText == "Another question" &&
							q.Difficulty == model.DifficultyEasy &&
							len(q.Options) == 2
					},
				)).Return(errors.New("database error")).Once()
			},
			expectedError: true,
			errorContains: "database error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			err := questService.CreateWithOptions(ctx, tt.createCred)

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
