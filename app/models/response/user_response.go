package response

import (
	// "fmt"
	"sample/app/models/db"
)

type UserDetailsResponse struct {
	// UserId   string `json:"userId"`
	// UserName string `json:"userName"`
	// Des      string `json:"des"`
	Users []db.User
}
