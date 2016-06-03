package models

type Flash struct {
	Type    string
	Message string
}

type Response struct {
	ReturnStatus     bool
	ReturnMessage    []string
	ValidationErrors map[string]interface{}
	TotalPages       int
	TotalRows        int
	PageSize         int
	IsAuthenticated  bool
	Data             map[string]interface{}
}

type Token struct {
	TransactionalInformation
	Token string `json:"token"`
}

type LoginDTO struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type TransactionalInformation struct {
	ReturnStatus     bool
	ReturnMessage    []string
	ValidationErrors map[string]interface{}
	TotalPages       int
	TotalRows        int
	PageSize         int
	IsAuthenticated  bool
}

type ApplicationMenuDTO struct {
	MenuID                 string `json:"menu_id"`
	Description            string `json:"description"`
	Route                  string `json:"route"`
	Module                 string `json:"module"`
	MenuOrder              int    `json:"menu_order"`
	RequiresAuthentication bool   `json:"requires_authentication"`
}

//ApplicationModelDTO user for return from  controllers.InitializeApplication
type ApplicationModelDTO struct {
	TransactionalInformation
	MenuItems []ApplicationMenuDTO `json:"menu_items"`
}
