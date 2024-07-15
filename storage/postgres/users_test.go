package postgres

import (
	"database/sql"
	"log"
	"testing"
	pb "travel/genproto/users"
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

func TestValidateUser(t *testing.T) {
	id := "9446b610-2ee7-46b4-98a1-ff905b016d2b"

	res := NewRepo().ValidateUser(id)
	if !res {
		t.Error("id not found")
	}
}

func TestEditProfile(t *testing.T) {
	req := pb.RequestEditProfile{
		Id:               "9446b610-2ee7-46b4-98a1-ff905b016d2b",
		FullName:         "Muhammadjon Ko'palov",
		Bio:              "",
		CountriesVisited: 2,
	}
	_, err := NewRepo().EditProfile(&req)
	if err != nil {
		t.Error(err)
	}
}

func TestGetUsers(t *testing.T) {
	filter := pb.RequestGetUsers{
		Page:  0,
		Limit: 5,
	}
	_, err := NewRepo().GetUsers(&filter)
	if err != nil {
		t.Error(err)
	}
}

func TestCountUsers(t *testing.T) {
	_, err := NewRepo().CountUsers()
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteUser(t *testing.T) {
	id := "ff9ae172-18f9-4f81-98ff-5db600ce05a7"
	_, err := NewRepo().DeleteUser(id)
	if err != nil {
		t.Error(err)
	}
}

func TestUpdatePassword(t *testing.T) {
	email := "p@gmail.com"
	password := "87654321"
	err := NewRepo().UpdatePassword(email, password)
	if err != nil {
		t.Error(err)
	}
}

func TestFindNumberOfVisitedCountries(t *testing.T) {
	id := "9446b610-2ee7-46b4-98a1-ff905b016d2b"
	_, err := NewRepo().FindNumberOfVisitedCountries(id)
	if err != nil {
		t.Error(err)
	}
}

func TestFollow(t *testing.T) {
	req := pb.RequestFollow{
		FollowerId:  "ff9ae172-18f9-4f81-98ff-5db600ce05a7",
		FollowingId: "9446b610-2ee7-46b4-98a1-ff905b016d2b",
	}
	err := NewRepo().Follow(&req)
	if err != nil || err == sql.ErrNoRows {
		t.Error(err)
	}
}

func TestGetFollowers(t *testing.T) {
	filter := pb.RequestGetFollowers{
		UserId: "9446b610-2ee7-46b4-98a1-ff905b016d2b",
		Limit:  10,
		Page:   0,
	}
	_, err := NewRepo().GetFollowers(&filter)
	if err != nil {
		t.Error(err)
	}
}

func TestGetFollowerInfo(t *testing.T) {
	id := "9446b610-2ee7-46b4-98a1-ff905b016d2b"
	_, err := NewRepo().GetProfile(id)
	if err != nil {
		t.Error(err)
	}
}
