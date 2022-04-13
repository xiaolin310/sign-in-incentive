package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	recordv1 "sign-in/api/record/v1"
	userv1 "sign-in/api/user/v1"
	"sign-in/app/signin/interface/internal/biz"

	v1 "sign-in/api/signin/interface/v1"
)

type SingInService struct {
	v1.UnimplementedSingInInterfaceServer

	au  *biz.AuthUsecase
	su  *biz.SignInUsecase
	wu  *biz.WalletUseCase
	log *log.Helper
}

var _ v1.SingInInterfaceServer = (*SingInService)(nil)

func NewSingInService(
	au *biz.AuthUsecase,
	su *biz.SignInUsecase,
	wu *biz.WalletUseCase,
	logger log.Logger,
) *SingInService {
	return &SingInService{
		au:  au,
		su:  su,
		wu:  wu,
		log: log.NewHelper(logger),
	}
}

func (s *SingInService) Register(ctx context.Context, req *v1.RegisterRequest) (reply *v1.RegisterReply, err error) {
	user, err := s.au.Register(ctx, &biz.User{
		Name:     req.Name,
		Password: req.Password,
		Phone:    req.Phone,
	})
	if err != nil {
		switch {
		case userv1.IsUserExist(err):
			err = v1.ErrorUserExistedError(err.Error())
		default:
			err = v1.ErrorUnknownError(err.Error())
		}
	} else {
		reply = &v1.RegisterReply {
			Id:    user.Id,
			Name:  user.Name,
			Phone: user.Phone,
		}
	}
	return
}

func (s *SingInService) Login(ctx context.Context, req *v1.LoginRequest) (reply *v1.LoginReply, err error) {
	token, err := s.au.Login(ctx, req.Username, req.Password)
	if err != nil {
		switch {
		case userv1.IsUserNotFound(err):
			err = v1.ErrorUserNotFound(err.Error())
		case err == biz.ErrIncorrectPassword:
			err = v1.ErrorPasswordIncorrect(err.Error())
		default:
			err = v1.ErrorUnknownError(err.Error())
		}
	} else {
		reply = &v1.LoginReply{
			Token: token,
		}
	}
	return
}

func (s *SingInService) GetSignInInfo(ctx context.Context, req *v1.GetSignInInfoRequest) (*v1.SignInInfoReply, error) {
	signindata, err := s.su.GetSignInInfo(ctx)
	if err != nil {
		switch {
		case recordv1.IsRecordNotFound(err):
			err = nil
		default:
			err = v1.ErrorUnknownError(err.Error())
		}
	}
	return toProtoSignInData(signindata), err
}


func (s *SingInService) SignIn(ctx context.Context, req *v1.SignInRequest) (*v1.SignInInfoReply, error) {
	singindata, err := s.su.SignIn(ctx)
	if err != nil {
		switch {
		case recordv1.IsRecordExisted(err):
			err = v1.ErrorRecordExisted(err.Error())
		default:
			err = v1.ErrorUnknownError(err.Error())
		}
	}
	return toProtoSignInData(singindata), err
}

func (s *SingInService) GetBalance(ctx context.Context, req *v1.GetBalanceRequest) (*v1.GetBalanceReply, error) {
	res, err := s.wu.GetBalance(ctx)
	if err != nil {
		err = v1.ErrorUnknownError(err.Error())
	}
	return &v1.GetBalanceReply{
		Balance: res.Balance,
	}, err
}

func toProtoSignInData(source *biz.SignInData) *v1.SignInInfoReply {
	return &v1.SignInInfoReply{
		RewardList:    source.RewardList,
		SignIndexList: source.SignInIndexList,
		IsSignToday:   source.IsSignToday,
	}
}