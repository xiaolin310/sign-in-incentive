package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type BalanceData struct {
	Balance float64
}

type WalletRepo interface {
	GetBalance(ctx context.Context) (*BalanceData, error)
}

type WalletUseCase struct {
	repo WalletRepo
	log  *log.Helper
}

func NewWalletUseCase(repo WalletRepo, logger log.Logger) *WalletUseCase {
	return &WalletUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *WalletUseCase) GetBalance(ctx context.Context) (*BalanceData, error) {
	return uc.repo.GetBalance(ctx)
}
