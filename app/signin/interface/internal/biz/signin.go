package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type SignInData struct {
	RewardList      []float64
	SignInIndexList []bool
	IsSignToday     bool
}

type RecordRepo interface {
	GetSignInInfo(context.Context) (*SignInData, error)
	SignIn(context.Context) (*SignInData, error)
}

type SignInUsecase struct {
	record RecordRepo
	log    *log.Helper
}

func NewSignInUsecase(repo RecordRepo, logger log.Logger) *SignInUsecase {
	return &SignInUsecase{record: repo, log: log.NewHelper(logger)}
}

func (uc *SignInUsecase) GetSignInInfo(ctx context.Context) (*SignInData, error) {
	return uc.record.GetSignInInfo(ctx)
}

func (uc *SignInUsecase) SignIn(ctx context.Context) (*SignInData, error) {
	return uc.record.SignIn(ctx)
}