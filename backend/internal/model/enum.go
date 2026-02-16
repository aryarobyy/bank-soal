package model

type (
	Difficulty    string
	Role          string
	Status        string
	SubjectTitle  string
	SessionStatus string
)

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

const (
	SessionInProgress SessionStatus = "in_progress"
	SessionFinished   SessionStatus = "finished"
	SessionExpired    SessionStatus = "expired"
)
