package dots

import "github.com/google/uuid"

type User struct {
	UserID      uuid.UUID `json:"user_id"`
	UserName    string    `json:"user_name"`
	UserLastName string    `json:"user_last_name"`
	UserEmail   string    `json:"user_email"`
	UserPassword string    `json:"user_password"`
	UserPhone   string    `json:"user_phone"`
	UserAddress string    `json:"user_address"`
	UserState   string    `json:"user_state"`
	UserMunicipality string `json:"user_municipality"`
}

