package api

type CreateUserRequest struct {
	Name        string `json:"name"`
	AccessLevel int    `json:"accessLevel"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

type UserLoginResponse struct {
	Name        string `json:"name"`
	AccessLevel int    `json:"accessLevel"`
	Email       string `json:"email"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AddPrinterToClass struct {
	ClassId   int `json:"name"`
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
}

type UpdateClass struct {
	ClassId string               `json:"classId"`
	Class   CreateClassesRequest `json:"class"`
}

type AddUserToClass struct {
	UserId  string `json:"userId"`
	ClassId int    `json:"classId"`
}

type RemoveUserFromClass struct {
	UserId  string `json:"userId"`
	ClassId int    `json:"classId"`
}

type AddPrinterTime struct {
	PrinterId  string `json:"printerId"`
	TimeslotId int    `json:"timeslotId"`
}

type BookPrinter struct {
	UserId    string `json:"userId"`
	StartTime int    `json:"startTime"`
	EndTime   int    `json:"endTime"`
	PrinterId int    `json:"printerId"`
}

type CreateAvailableTime struct {
	StartTime int `json:"startTime"`
	EndTime   int `json:"endTime"`
	ClassId   int `json:"classId"`
}

type Class struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsActive    string `json:"isActive"`
}
