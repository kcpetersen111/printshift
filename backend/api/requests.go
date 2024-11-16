package api

type ListPrintersRequest struct {
	ClassId int `json:"classId"`
}

type UpdatePrinterRequest struct {
	PrinterId   string `json:"printerId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsActive    bool   `json:"isActive"`
}

type Printer struct {
	PrinterId   string `json:"printerId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsActive    bool   `json:"isActive"`
}

type CreateUserRequest struct {
	Name        string `json:"name"`
	AccessLevel int    `json:"accessLevel"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

type UpdateUserRequest struct {
	Name        string `json:"name"`
	AccessLevel int    `json:"accessLevel,omitempty"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

type UserLoginResponse struct {
	Name        string `json:"name"`
	AccessLevel int    `json:"accessLevel"`
	Email       string `json:"email"`
}

type Professor struct {
	ProfessorId string `json:"professorId"`
	Name        string `json:"name"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AddPrinterToClass struct {
	ClassId   int `json:"classId"`
	PrinterId int `json:"printerId"`
}

type CreatePrinterRequest struct {
	Name     string `json:"name"`
	IsActive bool   `json:"isActive"`
}

type UpdatePrinter struct {
	PrinterId string               `json:"printerId"`
	Printer   CreatePrinterRequest `json:"printer"`
}

type CreateClassesRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	IsActive    bool   `json:"isActive"`
	ProfessorId int    `json:"professorId"`
}

type UpdateClass struct {
	ClassId     int    `json:"classId"`
	ProfessorId int    `json:"professorId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsActive    bool   `json:"isActive"`
}

type ListClassesRequest struct {
	UserId int `json:"userId"`
}

type AddUserToClass struct {
	UserId  int `json:"userId"`
	ClassId int `json:"classId"`
}

type RemoveUserFromClass struct {
	UserId  int `json:"userId"`
	ClassId int `json:"classId"`
}

type AddPrinterTime struct {
	PrinterId  int `json:"printerId"`
	TimeslotId int `json:"timeslotId"`
}

type BookSpecificPrinter struct {
	UserId    int `json:"userId"`
	StartTime int `json:"startTime"`
	EndTime   int `json:"endTime"`
	PrinterId int `json:"printerId"`
}

type CreateAvailableClassTime struct {
	StartTime int `json:"startTime"`
	EndTime   int `json:"endTime"`
	ClassId   int `json:"classId"`
}

type CreateAvailablePrinterTime struct {
	StartTime int `json:"startTime"`
	EndTime   int `json:"endTime"`
	PrinterId int `json:"printerId"`
}

type Class struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsActive    string `json:"isActive"`
}
