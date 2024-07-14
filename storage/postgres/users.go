package postgres

import (
	"database/sql"
	"log/slog"
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
			id = $1
	`
	bio := sql.NullString{}
	err := u.DB.QueryRow(query, userId).Scan(&resp.Username, &resp.Email,
		&resp.FullName, &bio, &resp.CountriesVisited,
		&resp.CreatedAt, &resp.UpdatedAt)

	resp.Bio = bio.String
	return &resp, err
}
