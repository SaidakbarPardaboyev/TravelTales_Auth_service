package postgres

import (
	"database/sql"
	"fmt"
	"log/slog"
	"time"
	pb "travel/genproto/users"
	"travel/pkg/logger"
)

type UserRepo struct {
	Logger *slog.Logger
	DB     *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	logger := logger.NewLogger()
	return &UserRepo{
		Logger: logger,
		DB:     db,
	}
}

func (u *UserRepo) GetProfile(userId string) (*pb.ResponseGetProfile, error) {
	resp := pb.ResponseGetProfile{Id: userId}

	query := `
		select
			username, email, full_name, bio, countries_visited,
			created_at, updated_at
		from
			users
		where 
			id = $1 and
			deleted_at is null
	`
	bio := sql.NullString{}
	err := u.DB.QueryRow(query, userId).Scan(&resp.Username, &resp.Email,
		&resp.FullName, &bio, &resp.CountriesVisited,
		&resp.CreatedAt, &resp.UpdatedAt)

	resp.Bio = bio.String
	return &resp, err
}

func (u *UserRepo) ValidateUser(userId string) bool {
	query := `
		select
			exists (
				select 1
				from users
				where id = $1 and
				deleted_at is null
			)
	`
	var exists bool
	u.DB.QueryRow(query, userId).Scan(&exists)
	return exists
}

func (u *UserRepo) EditProfile(user *pb.RequestEditProfile) (
	*pb.ResponseEditProfile, error) {

	query := `
		update
			users
		set
			full_name = $1,
			bio = $2,
			countries_visited = $3,
			updated_at = $4
		where
			id = $5 and
			deleted_at is null
		returning id, username, email, full_name, bio, countries_visited,
		updated_at
	`
	resp := pb.ResponseEditProfile{}
	err := u.DB.QueryRow(query, user.FullName, user.Bio, user.CountriesVisited,
		time.Now(), user.Id).Scan(&resp.Id, &resp.Username, &resp.Email, &resp.FullName,
		&resp.Bio, &resp.CountriesVisited, &resp.UpdatedAt)

	return &resp, err
}

func (u *UserRepo) GetUsers(filter *pb.RequestGetUsers) (
	*pb.ResponseGetUsers, error) {

	query := `
		select
			id, username, full_name, countries_visited
		from 
			users
		where
			deleted_at is null
		limit $1
		offset $2
	`

	rows, err := u.DB.Query(query, filter.Limit, filter.Page*filter.Limit)
	if err != nil {
		return nil, err
	}

	resp := pb.ResponseGetUsers{}
	for rows.Next() {
		user := pb.User{}
		err := rows.Scan(&user.Id, &user.Username, &user.FullName,
			&user.CountriesVisited)
		if err != nil {
			return nil, err
		}

		resp.Users = append(resp.Users, &user)
	}

	resp.Limit = filter.Limit
	resp.Page = filter.Page

	total, err := u.CountUsers()
	if err != nil {
		return nil, err
	}
	resp.Total = int64(total)

	return &resp, nil
}

func (u *UserRepo) CountUsers() (int, error) {
	query := `
		select
			count(*)
		from 
			users
		where
			deleted_at is null
	`

	res := 0
	err := u.DB.QueryRow(query).Scan(&res)
	return res, err
}

func (u *UserRepo) DeleteUser(userId string) (*pb.ResponseDeleteUser, error) {

	query := `
		update 
			users
		set
			deleted_at = $1
		where
			id = $2 and
			deleted_at is null
	`

	res, err := u.DB.Exec(query, time.Now(), userId)
	if err != nil {
		return nil, err
	}
	if num, _ := res.RowsAffected(); num <= 0 {
		return nil, fmt.Errorf("the user was already deleted")
	}
	return &pb.ResponseDeleteUser{Message: "User is deleted successfully"}, nil
}

func (u *UserRepo) UpdatePassword(email, password string) error {

	query := `
		update
			users
		set
			password = $1
		where
			email = $2 and 
			deleted_at is null
	`

	res, err := u.DB.Exec(query, password, email)
	if err != nil {
		return err
	}
	if num, _ := res.RowsAffected(); num <= 0 {
		return fmt.Errorf("the user was already deleted")
	}
	return nil
}

func (u *UserRepo) FindNumberOfVisitedCountries(userId string) (
	int, error) {

	query := `
		select
			countries_visited
		from
			users
		where
			id = $1 and 
			deleted_at is null
	`
	numberOfVisitedCountries := 0
	err := u.DB.QueryRow(query, userId).Scan(&numberOfVisitedCountries)
	if err != nil || err == sql.ErrNoRows {
		return 0, err
	}
	return numberOfVisitedCountries, nil
}

func (u *UserRepo) Follow(req *pb.RequestFollow) error {

	query := `
		insert into following(
			follower_id, following_id, followed_at
		) values (
			$1, $2, $3
		)
	`

	_, err := u.DB.Exec(query, req.FollowerId, req.FollowingId, time.Now())
	return err
}

func (u *UserRepo) GetFollowers(filter *pb.RequestGetFollowers) (
	*pb.ResponseGetFollowers, error) {

	query := `
		select
			follower_id
		from
			following
		where
			following_id = $1
		limit $2
		offset $3
	`

	rows, err := u.DB.Query(query, filter.UserId, filter.Limit,
		filter.Page*filter.Limit)
	if err != nil {
		return nil, err
	}

	resp := pb.ResponseGetFollowers{}
	for rows.Next() {
		var id string
		err := rows.Scan(&id)
		if err != nil {
			return nil, err
		}

		follower, err := u.GetFollowerInfo(id)
		if err != nil {
			if err == sql.ErrNoRows {
				continue
			}
			return nil, err
		}

		resp.Followers = append(resp.Followers, follower)
	}

	return &resp, nil
}

func (u *UserRepo) GetFollowerInfo(userid string) (
	*pb.Follower, error) {

	query := `
		select
			id, username, full_name
		from 
			users
		where
			id = $1 and
			deleted_at is null
	`

	res := pb.Follower{}
	err := u.DB.QueryRow(query, userid).Scan(&res.Id, &res.Username,
		&res.FullName)

	return &res, err
}
