package api

type CreateUserRequest struct {
	Name     string `json:"name"`
	Level    int    `json:"level"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AddPrinterToClass struct {
	ClassId   int `json:"name"`
	PrinterId int `json:"active"`
}

type CreatePrinterRequest struct {
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

type CreateClassesRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Active      bool   `json:"active"`
}
