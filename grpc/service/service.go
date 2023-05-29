package service

import (
	"context"
	"fmt"
	"user_service/config"
	"user_service/genproto/order_service"
	"user_service/genproto/user_service"
	"user_service/grpc/client"
	"user_service/pkg/logger"
	"user_service/storage"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type userService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	user_service.UnimplementedUserServiceServer
}

func NewUserService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, svcs client.ServiceManagerI) *userService {
	return &userService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: svcs,
	}
}

func (b *userService) Create(ctx context.Context, req *user_service.CreateUserRequest) (resp *user_service.User, err error) {
	b.log.Info("---CreateUser--->", logger.Any("req", req))

	b.services.OrderService().Create(ctx, &order_service.CreateOrderRequest{
		UserID:"test",
		ProductID: "test",
		TotalSum: 2000,
	})


	pKey, err := b.strg.User().Create(ctx, req)

	if err != nil {
		b.log.Error("!!!CreateUser--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return b.strg.User().GetById(ctx, pKey)
}

func (u *userService) GetById(ctx context.Context, pKey *user_service.UserPKey) (resp *user_service.User, err error) {
	u.log.Info("---GetUserById--->", logger.Any("pKey", pKey))

	resp, err = u.strg.User().GetById(ctx, pKey)
	if err != nil {
		u.log.Error("!!!GetUserById--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return u.strg.User().GetById(ctx, pKey)
}

func (u *userService) GetAll(ctx context.Context, req *user_service.GetAllUsersRequest) (resp *user_service.GetAllUsersResponse, err error) {
	u.log.Info("---GetAllUsers--->", logger.Any("req", req))

	resp, err = u.strg.User().GetAll(ctx, req)
	if err != nil {
		u.log.Error("!!!GetAllUsers--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (u *userService) Delete(ctx context.Context, pKey *user_service.UserPKey) (resp *emptypb.Empty ,err error) {
	u.log.Info("---DeleteUser--->", logger.Any("pKey", pKey))

	resp = &emptypb.Empty{}
	err = u.strg.User().Delete(ctx, pKey)
	if err != nil {
		u.log.Error("!!!DeleteOrder--->", logger.Error(err))
		return nil,status.Error(codes.InvalidArgument, err.Error())
	}
	return resp, nil
}

// func (u *userService) Update(ctx context.Context, req *user_service.UpdateUserRequest) (resp *emptypb.Empty, err error) {
// 	u.log.Info("---UpdateUser--->", logger.Any("req", req))

// 	resp = &emptypb.Empty{}
// 	err = u.strg.User().Update(ctx, req)
// 	if err != nil {
// 		u.log.Error("!!!UpdateUser--->", logger.Error(err))
// 		return nil, status.Error(codes.InvalidArgument, err.Error())
// 	}

// 	return resp, nil
// }

func (u *userService) GetByPhone(ctx context.Context, req *user_service.GetUserByPhoneRequest) (resp *user_service.GetUserByPhoneResponse, err error) {
	u.log.Info("---GetUserByPhone--->", logger.Any("req", req))


	resp, err = u.strg.User().GetByPhone(ctx, req)
	fmt.Println("resp", resp)
	if err != nil {
		u.log.Error("!!!GetUserByPhone--->", logger.Error(err))
		return resp, status.Error(codes.InvalidArgument, err.Error())
	}

	fmt.Println(u.strg.User().GetByPhone(ctx, req))
	return u.strg.User().GetByPhone(ctx, req)
}
