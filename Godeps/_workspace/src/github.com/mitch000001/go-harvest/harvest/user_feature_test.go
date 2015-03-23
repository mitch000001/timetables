// +build feature

package harvest_test

import (
	"testing"
	"time"

	"github.com/mitch000001/timetables/Godeps/_workspace/src/github.com/mitch000001/go-harvest/harvest"
)

func TestFindAllUsersUpdatedSince(t *testing.T) {
	client := createClient(t)
	updatedSince := time.Now().Add(-2 * time.Second)
	t.Logf("UpdatedSince: %+#v\n", updatedSince)
	users, err := client.Users.AllUpdatedSince(updatedSince)
	if err != nil {
		t.Fatalf("Got error %T with message: %s\n", err, err.Error())
	}
	for _, u := range users {
		t.Logf("User: '%+#v'\n", u)
	}
	if len(users) != 1 {
		t.Fatalf("Expected 1 user, got %d", len(users))
	}
}

func TestFindAllUsers(t *testing.T) {
	client := createClient(t)
	users, err := client.Users.All()
	if err != nil {
		t.Fatalf("Got error %T with message: %s\n", err, err.Error())
	}
	if len(users) != 1 {
		t.Fatalf("Expected 1 user, got %d", len(users))
	}
	if users[0] == nil {
		t.Fatal("Expected user not to be nil")
	}
	for _, u := range users {
		t.Logf("User: %+#v\n", u)
	}
}

func TestFindUser(t *testing.T) {
	client := createClient(t)

	// Find existing user
	users, err := client.Users.All()
	if err != nil {
		t.Fatalf("Got error %T with message: %s\n", err, err.Error())
	}
	first := users[0]
	user, err := client.Users.Find(first.ID)
	if err != nil {
		t.Fatalf("Got error %T with message: %s\n", err, err.Error())
	}
	if first.ID != user.ID {
		t.Fatalf("Expect to find user with id %d, got user %#v\n", first.Id, user)
	}

	// No user with that id
	user, err = client.Users.Find(-1)
	if err != nil {
		expectedErrorMessage := "No user found with id -1"
		if err.Error() != expectedErrorMessage {
			t.Fatalf("Expected ResponseError with message '%s', got error %T with message: %s\n", expectedErrorMessage, err, err.Error())
		}
	}
	if user != nil {
		t.Fatalf("Expected user to be nil, got '%+#v'", user)
	}
}

func TestCreateAndDeleteUser(t *testing.T) {
	client := createClient(t)
	user := harvest.User{
		FirstName: "Foo",
		LastName:  "Bar",
		Email:     "foo@example.com"}
	createdUser, err := client.Users.Create(&user)
	if err != nil {
		t.Fatalf("Got error %T with message: %s\n", err, err.Error())
	}
	if createdUser == nil {
		t.Fatal("Expected user not to be nil")
	}
	t.Logf("Got returned user: %+#v\n", createdUser)
	deleted, err := client.Users.Delete(&user)
	if err != nil {
		t.Fatalf("Got error %T with message: %s\n", err, err.Error())
	}
	if !deleted {
		t.Fatalf("Could not delete user created for test")
	}
}

func TestUpdateUser(t *testing.T) {
	client := createClient(t)
	user := &harvest.User{
		FirstName:  "Foo",
		LastName:   "Bar",
		Email:      "foo@example.com",
		Department: "Old Department",
	}
	user, err := client.Users.Create(user)
	if err != nil {
		panic(err)
	}
	defer client.Users.Delete(user)
	user.Department = "New Department"
	updatedUser, err := client.Users.Update(user)
	if err != nil {
		t.Fatalf("Got error %T with message: %s\n", err, err.Error())
	}
	if updatedUser.Department != user.Department {
		t.Fatalf("Expected updated department to equal '%s', got '%s'", user.Department, updatedUser.Department)
	}

	// Wrong updates
	user.Email = "hdhi6556"
	updatedUser, err = client.Users.Update(user)
	if err == nil {
		t.Fatal("Expected ResponseError, got nil")
	}
	if updatedUser != nil {
		t.Fatalf("Expected user to be nil, got '%+#v'", updatedUser)
	}

}
