package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"sign-in/app/user/service/internal/biz"
	"sign-in/app/user/service/internal/data/ent"

	pb "sign-in/api/user/v1"
)

type UserService struct {
	pb.UnimplementedUserServiceServer

	uc  *biz.UserUsecase
	log *log.Helper
}

var _ pb.UserServiceServer = (*UserService)(nil)

func NewUserService(uc *biz.UserUsecase, logger log.Logger) *UserService {
	return &UserService{uc: uc, log: log.NewHelper(logger)}
}

func (s *UserService) GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (reply *pb.GetUserByIdReply, err error) {
	users, err := s.uc.GetUserById(ctx, req.Id)
	if err != nil {
		switch err.(type) {
		case *ent.NotFoundError:
			err = pb.ErrorUserNotFound("error: get user failed, reason: user not found")
		default:
			err = pb.ErrorUnknownError("error: get user failed, reason: unknown error")
		}
	} else {
		reply = &pb.GetUserByIdReply{
			User: bulk2ProtoUser(users),
		}
	}
	return
}
func (s *UserService) GetUserByName(ctx context.Context, req *pb.GetUserByNameRequest) (reply *pb.GetUserByNameReply, err error) {
	user, err := s.uc.GetUserByName(ctx, req.Name)
	if err != nil {
		switch err.(type) {
		case *ent.NotFoundError:
			err = pb.ErrorUserNotFound("error: get user failed, reason: user not found")
		default:
			err = pb.ErrorUnknownError("error: get user failed, reason: unknown error")
		}
	} else {
		reply = &pb.GetUserByNameReply{
			User: toProtoUser(user),
		}
	}
	return
}
func (s *UserService) SearchUserByName(ctx context.Context, req *pb.SearchUserByNameRequest) (reply *pb.SearchUserByNameReply, err error) {
	users, err := s.uc.SearchUserByName(ctx, req.Name)
	if err != nil {
		err = pb.ErrorUnknownError("error: search user failed, reason: unknown error")
	} else {
		reply = &pb.SearchUserByNameReply{
			User: bulk2ProtoUser(users),
		}
	}
	return
}
func (s *UserService) SaveUser(ctx context.Context, req *pb.SaveUserRequest) (reply *pb.SaveUserReply, err error) {
	user, err := s.uc.SaveUser(ctx, toBizUser(req.User))
	if err != nil {
		switch err.(type) {
		case *ent.ConstraintError:
			err = pb.ErrorUserExist("error: save user failed, reason: user name existed")
		default:
			err = pb.ErrorUnknownError("error: save user failed, reason: unknown error")
		}
	} else {
		reply = &pb.SaveUserReply{
			User: toProtoUser(user),
		}
	}
	return
}
func (s *UserService) RemoveUser(ctx context.Context, req *pb.RemoveUserRequest) (reply *pb.RemoveUserReply, err error) {
	err = s.uc.RemoveUser(ctx, req.Id)
	if err != nil {
		err = pb.ErrorUnknownError("error: remove user failed, reason: unknown error")
	}
	reply = &pb.RemoveUserReply{
		Code: 200,
	}
	return
}

func bulk2ProtoUser(source []*biz.User) []*pb.User {
	user := make([]*pb.User, 0, len(source))
	for _, v := range source {
		user = append(user, toProtoUser(v))
	}
	return user
}

func toProtoUser(source *biz.User) *pb.User {
	return &pb.User{
		Id:       source.Id,
		Name:     source.Name,
		Password: source.Password,
		Phone:    source.Phone,
	}
}

func toBizUser(source *pb.User) *biz.User {
	return &biz.User{
		Id:       source.Id,
		Name:     source.Name,
		Password: source.Password,
		Phone:    source.Phone,
	}
}