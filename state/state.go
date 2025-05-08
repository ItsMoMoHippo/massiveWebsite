package state

type queryFilter struct {
	Club string
}

var Current queryFilter

func SetClub(club string) {
	Current.Club = club
}

func GetClub() string {
	return Current.Club
}
