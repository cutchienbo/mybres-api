package db

type Team struct {
	UserId		int64
	PartnerId	int64
	CreatedAt	string
	User 		User
	Partner		User
}
