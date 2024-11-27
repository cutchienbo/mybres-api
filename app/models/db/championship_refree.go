package db

type ChampionshipRefree struct {
	ChampionshipId  int64
	UserId 			int64
	CreatedAt		string

	Championship 	Championship
	User 			User
}
