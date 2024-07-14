package postgres

import (
	"log"
	"testing"
)

func NewRepo() *UserRepo {
	db, err := ConnectDB()
	if err != nil {
		log.Panic(err)
	}
	return NewUserRepo(db)
}
func TestGetProfile(t *testing.T) {
	id := "9446b610-2ee7-46b4-98a1-ff905b016d2b"

	_, err := NewRepo().GetProfile(id)
	if err != nil {
		t.Error(err)
	}
}
