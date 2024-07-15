package service

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"travel/pkg/logger"
	"travel/storage/postgres"

	pb "travel/genproto/users"
)

type UserService struct {
	pb.UnimplementedUsersServer
	Logger   *slog.Logger
	UserRepo *postgres.UserRepo
}

func NewUserService(db *sql.DB) *UserService {
	userRepo := postgres.NewUserRepo(db)
	Logger := logger.NewLogger()
	return &UserService{
		Logger:   Logger,
		UserRepo: userRepo,
	}
}

func (u *UserService) GetProfile(ctx context.Context, in *pb.RequestGetProfile) (
	*pb.ResponseGetProfile, error) {
	resp, err := u.UserRepo.GetProfile(in.Id)
	if err != nil {
		u.Logger.Error(fmt.Sprintf("error with Getting user profile from db: %s", err))
		return nil, err
	}
	return resp, nil
}

func (u *UserService) ValidateUser(ctx context.Context, in *pb.RequestGetProfile) (
	*pb.Status, error) {
	resp := u.UserRepo.ValidateUser(in.Id)
	return &pb.Status{Success: resp}, nil
}

func (u *UserService) EditProfile(ctx context.Context, in *pb.RequestEditProfile) (
	*pb.ResponseEditProfile, error) {
	resp, err := u.UserRepo.EditProfile(in)
	if err != nil {
		u.Logger.Error(fmt.Sprintf("error with Getting user profile from db: %s", err))
		return nil, err
	}
	return resp, nil
}

func (u *UserService) GetUsers(ctx context.Context, in *pb.RequestGetUsers) (*pb.ResponseGetUsers, error) {
	resp, err := u.UserRepo.GetUsers(in)
	if err != nil {
		u.Logger.Error(fmt.Sprintf("error with Getting users from db: %s", err))
		return nil, err
	}
	return resp, nil
}

func (u *UserService) DeleteUser(ctx context.Context, in *pb.RequestDeleteUser) (
	*pb.ResponseDeleteUser, error) {
	resp, err := u.UserRepo.DeleteUser(in.Id)
	if err != nil {
		u.Logger.Error(fmt.Sprintf("error with Getting users from db: %s", err))
		return nil, err
	}
	return resp, nil
}	

// func (u *UserService) UpdatePassword(ctx context.Context, in *pb.RequestUpdatePassword) (*pb.ResponseUpdatePassword, error) {
// }
// func (u *UserService) GetUserStatistic(ctx context.Context, in *pb.RequestGetUserStatistic) (*pb.ResponseGetUserStatistic, error) {
// }
// func (u *UserService) Follow(ctx context.Context, in *pb.RequestFollow) (*pb.ResponseFollow, error) {
// }
// func (u *UserService) GetFollowers(ctx context.Context, in *pb.RequestGetFollowers) (*pb.ResponseGetFollowers, error) {
// }
