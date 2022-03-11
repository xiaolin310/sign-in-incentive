package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"sign-in/app/virtualwallet/service/internal/biz"
	"sign-in/app/virtualwallet/service/internal/data/ent"
	"sign-in/app/virtualwallet/service/internal/data/ent/virtualwallet"
	"time"
)

type walletRepo struct {
	data *Data
	log  *log.Helper
}

func NewWalletRepo(data *Data, logger log.Logger) biz.VirtualWalletRepo {
	return &walletRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *walletRepo) GetWallet(ctx context.Context, userId int64) (*biz.VirtualWallet, error) {
	wallet, err := r.data.db.VirtualWallet.Query().
		Where(
			virtualwallet.UserIDEQ(userId),
			).
		First(ctx)
	if err != nil {
		return nil, err
	}
	return toBizVirtualWallet(wallet), nil
}

func (r *walletRepo) CreateWallet(ctx context.Context, wallet *biz.VirtualWallet) (*biz.VirtualWallet, error) {
	walletEntity, err := r.data.db.VirtualWallet.Create().
		SetUserID(wallet.UserId).
		SetCreatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return toBizVirtualWallet(walletEntity), nil
}

func (r *walletRepo)UpdateBalance(ctx context.Context, wallet *biz.VirtualWallet) error {
	_, err := r.data.db.VirtualWallet.Update().
		Where(
			virtualwallet.UserIDEQ(wallet.UserId),
			).
		SetBalance(wallet.Balance).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
}


func toBizVirtualWallet(source *ent.VirtualWallet) *biz.VirtualWallet {
	return &biz.VirtualWallet {
		Id:      source.ID,
		UserId:  source.UserID,
		Balance: source.Balance,
	}
}