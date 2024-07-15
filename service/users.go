package service

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"time"
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
		u.Logger.Error(fmt.Sprintf("error with editing user profile from db: %s", err))
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
		u.Logger.Error(fmt.Sprintf("error with deleting users from db: %s", err))
		return nil, err
	}
	return resp, nil
}

func (u *UserService) UpdatePassword(ctx context.Context, in *pb.RequestUpdatePassword) (
	*pb.ResponseUpdatePassword, error) {
	err := u.UserRepo.UpdatePassword(in.Email, in.NewPassword)
	if err != nil {
		u.Logger.Error(fmt.Sprintf("error with updating users from db: %s", err))
		return nil, err
	}
	return &pb.ResponseUpdatePassword{
		Message: "Password updated successfully",
	}, nil
}

func (u *UserService) GetUserStatistic(ctx context.Context, in *pb.RequestGetUserStatistic) (
	*pb.ResponseGetUserStatistic, error) {

	numberOfVisitedCountries, err := u.UserRepo.FindNumberOfVisitedCountries(in.Id)
	if err != nil {
		u.Logger.Error(fmt.Sprintf("error with getting user's statistics from db: %s", err))
		return nil, err
	}

	resp := pb.ResponseGetUserStatistic{
		UserId:           in.Id,
		CountriesVisited: int64(numberOfVisitedCountries),
	}

	// finding number of stories

	// finding number of comments

	// finding number of likes

	// finding user's last active

	return &resp, nil
}

func (u *UserService) Follow(ctx context.Context, in *pb.RequestFollow) (
	*pb.ResponseFollow, error) {
	if in.FollowerId == in.FollowingId {
		u.Logger.Error("error with follower and folloving id are the same")
		return nil, fmt.Errorf("error with follower and folloving id are the same")
	}

	err := u.UserRepo.Follow(in)
	if err != nil {
		u.Logger.Error(fmt.Sprintf("error with following a user from db: %s", err))
		return nil, err
	}
	return &pb.ResponseFollow{
		FollowerId:  in.FollowerId,
		FollowingId: in.FollowingId,
		FollowedAt:  time.Now().String(),
	}, nil
}

func (u *UserService) GetFollowers(ctx context.Context, in *pb.RequestGetFollowers) (
	*pb.ResponseGetFollowers, error) {
	resp, err := u.UserRepo.GetFollowers(in)
	if err != nil {
		u.Logger.Error(fmt.Sprintf("error with getting a user's followers from db: %s", err))
		return nil, err
	}
	return resp, nil
}
