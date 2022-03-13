package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "sign-in/api/virtualwallet/v1"
	"sign-in/app/virtualwallet/service/internal/biz"
)

type VirtualWalletService struct {
	v1.UnimplementedVirtualWalletServer

	uc  *biz.VirtualWalletUsecase
	log *log.Helper

}

var _ v1.VirtualWalletServer = (*VirtualWalletService)(nil)

func NewVirtualWalletService(uc *biz.VirtualWalletUsecase, logger log.Logger) *VirtualWalletService {
	return &VirtualWalletService{uc: uc, log: log.NewHelper(logger)}
}

func (s *VirtualWalletService) GetBalance(ctx context.Context, req *v1.GetBalanceRequest) (*v1.GetBalanceReply, error) {
	balance, err := s.uc.GetBalance(ctx, req.User)
	if err != nil {
		err = v1.ErrorUnknownError("error: get balance failed, reason: unknown error")
	}
	return &v1.GetBalanceReply{Balance: balance}, err
}

func (s *VirtualWalletService) Debit(ctx context.Context, req *v1.DebitRequest) (*v1.DebitReply, error) {
	err := s.uc.Debit(ctx, req.User, req.Amount)
	if err != nil {
		// TODO: how to handle the new error, BALANCE_NOT_SUFFICIENT
		return nil, err
	}
	return &v1.DebitReply{Code: 200}, nil
}

func (s *VirtualWalletService) Credit(ctx context.Context, req *v1.CreditRequest) (*v1.CreditReply, error) {
	err := s.uc.Credit(ctx, req.User, req.Amount)
	if err != nil {
		// TODO: how to handle the new error, INVALID_AMOUNT
		return nil, err

	}
	return &v1.CreditReply{Code: 200}, nil
}

