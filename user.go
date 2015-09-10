package timetables

func NewUser(harvestID int, firstName, lastName, email string) User {
	return User{
		HarvestID: harvestID,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}
}

type User struct {
	ID        string
	HarvestID int
	FirstName string
	LastName  string
	Email     string
}
