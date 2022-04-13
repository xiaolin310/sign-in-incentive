package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	walletClientV1 "sign-in/api/virtualwallet/v1"
	"sign-in/app/signin/interface/internal/biz"
)

type walletRepo struct {
	data *Data
	log  *log.Helper
}

var _ biz.WalletRepo = (*walletRepo)(nil)

func NewWalletRepo(data *Data, logger log.Logger) biz.WalletRepo {
	return &walletRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func(repo *walletRepo) GetBalance(ctx context.Context) (*biz.BalanceData, error) {
	val, err := repo.data.wc.GetBalance(ctx, &walletClientV1.GetBalanceRequest{
		User: int64(ctx.Value("user").(int)),
	})
	if err != nil {
		return nil, err
	}
	return &biz.BalanceData{
		Balance: val.Balance,
	}, nil
}
