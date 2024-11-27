package db

type ChampionshipPrize struct {
	Id				int64
	ChampionshipId	int64
	Prize           string
	Sort			int16
	CreatedAt		string

	Championship 	Championship
}
