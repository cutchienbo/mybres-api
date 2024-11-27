package db

type ChampionshipSlot struct {
	Id				int64
	ChampionshipId	int64
	Score			int64
	CreatedAt		string

	Championship 	Championship
}
