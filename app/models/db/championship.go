package db

type Championship struct {
	Id			int64
	CourtId 	int64
	Name		string
	Thumbnail	string
	Format 		int16
	MinScore	int64
	MaxScore	int64
	DateStart	string
	DateCancel 	string
	TimeStart	string
	TimeEnd		string
	CourtCount	int16
	Description string
	Level		string
	Address		string
	Expense		int64
	IsCensor	int16
	PayMethods	string
	Status 		int16
	CreatedAt	string

	Court Court
}
