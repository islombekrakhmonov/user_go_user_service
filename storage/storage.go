package storage

import (
	"context"
	"user_service/genproto/user_service"
)

type StorageI interface	{
	CloseDB()
	User() UserRepoI
}

type UserRepoI interface {
	Create(ctx context.Context, req *user_service.CreateUserRequest) (pKey *user_service.UserPKey, err error)
	Get(ctx context.Context, pKey *user_service.UserPKey) (resp *user_service.User, err error)
	GetAll(ctx context.Context, req *user_service.GetAllUsersRequest) (resp *user_service.GetAllUsersResponse, err error)
	Delete(ctx context.Context, pKey *user_service.UserPKey) (err error)
	Update(ctx context.Context, pKey *user_service.UserPKey) (err error)
}