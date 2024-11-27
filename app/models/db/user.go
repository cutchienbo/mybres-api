package db

type User struct {
	Id      	int64
	Name    	string
	Email 		string
	Sex 		bool
	Birthday 	string
	Avatar 		string
	Description string
	Password 	string
	Score 		int64
	IsRefree	bool
	Status 		int16
	CreatedAt 	string
}
