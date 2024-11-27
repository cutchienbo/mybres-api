package db

type Court struct {
	Id			int64
	Name		string
	PhoneNumber	string
	Address		string
	Description	string
	LatLong 	string
	CreatedAt 	string
	
	Championships []Championship
}
