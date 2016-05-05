package models

type UserInRole struct {
	RoleID string `json:"role_id"`
	Role   Role   `json:"role"`
	UserID string `json:"user_id"`
	User   User   `json:"user"`
}
