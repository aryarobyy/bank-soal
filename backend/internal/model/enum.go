package model

type Difficulty string
type Role string
type Status string
type SubjectTitle string

const (
	DifficultyEasy   Difficulty = "easy"
	DifficultyMedium Difficulty = "medium"
	DifficultyHard   Difficulty = "hard"
)

const (
	StatusPassed    Status = "passed"
	StatusNotPassed Status = "not_passed"
)

const (
	RoleAdmin      Role = "admin"
	RoleUser       Role = "user"
	RoleSuperAdmin Role = "super_admin"
	RoleLecturer   Role = "lecturer"
)

const (
	SubjectKalkulus SubjectTitle = "Kalkulus"
	SubjectMatDis   SubjectTitle = "Matematika Diskrit"
	SubjectAutomata SubjectTitle = "Teori Bahasa dan Automata"
	SubjectData     SubjectTitle = "Basis Data Lanjut"
	SubjectMetNum   SubjectTitle = "Metode Numerik"
)
