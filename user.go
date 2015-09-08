package timetables

func CreateUser(harvestID int, firstName, lastName, email string) User {
	return User{
		ID:        "1",
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
