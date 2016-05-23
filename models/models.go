package models

type Flash struct {
	Type    string
	Message string
}

type Token struct {
	TransactionalInformationDTO
	Token string `json:"token"`
}

type LoginDTO struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type TransactionalInformationDTO struct {
	ReturnStatus     bool                   `json:"return_status"`
	ReturnMessage    []string               `json:"return_message"`
	ValidationErrors map[string]interface{} `json:"validation_errors"`
	TotalPages       int                    `json:"total_pages"`
	TotalRows        int                    `json:"total_rows"`
	PageSize         int                    `json:"page_size"`
	IsAuthenticated  bool                   `json:"is_authenticated"`
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
	TransactionalInformationDTO
	MenuItems []ApplicationMenuDTO `json:"menu_items"`
}
