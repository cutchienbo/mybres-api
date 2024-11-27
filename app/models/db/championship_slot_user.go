package db

type ChampionshipSlotUser struct {
	ChampionshipSlotId	int64
	UserId				int64
	CreatedAt		string

	ChampionshipSlot 	ChampionshipSlot
}
