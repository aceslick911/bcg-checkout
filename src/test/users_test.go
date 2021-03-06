package test

import (
	"fmt"
	"testing"

	"github.com/aceslick911/bcg-checkout/internal/pkg/config"
	"github.com/aceslick911/bcg-checkout/internal/pkg/db"
	models "github.com/aceslick911/bcg-checkout/internal/pkg/models/users"
	"github.com/aceslick911/bcg-checkout/internal/pkg/persistence"
)

var userTest models.User

func SetupUsers() {
	config.Setup("./config.yml")
	db.SetupDB()
	db.GetDB().Exec("DELETE FROM users")
}

func TestAddUser(t *testing.T) {
	SetupUsers()
	user := models.User{
		Firstname: "Angelo",
		Lastname:  "Perera",
		Username:  "aceslick911",
		Hash:      "hash",
		Role:      models.UserRole{RoleName: "user"},
	}
	s := persistence.GetUserRepository()
	if err := s.Add(&user); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	userTest = user
}

func TestGetAllUsers(t *testing.T) {
	s := persistence.GetUserRepository()
	if _, err := s.All(); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestGetUserById(t *testing.T) {
	db.SetupDB()
	s := persistence.GetUserRepository()
	if _, err := s.Get(fmt.Sprint(userTest.ID)); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}
