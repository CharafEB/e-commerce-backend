package dots

import "github.com/google/uuid"

type Car struct {
	UserID     uuid.UUID `json:"user_id"`
	ProductsID uuid.UUID `json:"prod_id"`
}
