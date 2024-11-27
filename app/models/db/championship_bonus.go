package db

type ChampionshipBonus struct {
	Id				int64
	ChampionshipId	int64
	Rank			string
	Score			int64
	Sort			int16
	CreatedAt		string

	Championship 	Championship
}
