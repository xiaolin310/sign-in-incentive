package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"math"
	"sign-in/app/virtualwallet/service/internal/data/ent"
)

type VirtualWallet struct {
	Id      uuid.UUID
	UserId  int64
	Balance float64
}

type VirtualWalletRepo interface {
	GetWallet(ctx context.Context, userId int64) (*VirtualWallet, error)
	CreateWallet(ctx context.Context, wallet *VirtualWallet) (*VirtualWallet, error)
	UpdateBalance(ctx context.Context, wallet *VirtualWallet) error
}

type VirtualWalletUsecase struct {
	repo VirtualWalletRepo
	log  *log.Helper
}


func NewVirtualWalletUsecase(repo VirtualWalletRepo, logger log.Logger) *VirtualWalletUsecase {
	return &VirtualWalletUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *VirtualWalletUsecase) GetBalance(ctx context.Context, userId int64) (float64, error) {
	wallet, err := uc.getVirtualWallet(ctx, userId)
	if err != nil {
		return -1, err
	}
	return wallet.Balance, nil
}

// Debit 扣款
func (uc *VirtualWalletUsecase) Debit(ctx context.Context, userId int64, amount float64) error {
	wallet, err := uc.getVirtualWallet(ctx, userId)
	if err != nil {
		return err
	}

	if compareTo(wallet.Balance, amount) < 0 {
		return errors.New(406, "balance is less than amount", "BALANCE_NOT_SUFFICIENT")
	}
	newBalance := wallet.Balance - amount
	err = uc.repo.UpdateBalance(ctx, &VirtualWallet{
		UserId:  userId,
		Balance: newBalance,
	})
	if err != nil {
		return err
	}
	return nil
}

// Credit 入账
func (uc *VirtualWalletUsecase) Credit(ctx context.Context, userId int64, amount float64) error {
	wallet, err := uc.getVirtualWallet(ctx, userId)
	if err != nil {
		return err
	}

	if compareTo(amount, 0) < 0 {
		return errors.New(406, "invalid amount, amount < 0", "INVALID_AMOUNT")
	}
	newBalance := wallet.Balance + amount
	err = uc.repo.UpdateBalance(ctx, &VirtualWallet{
		UserId:  userId,
		Balance: newBalance,
	})
	if err != nil {
		return err
	}
	return nil
}

func (uc *VirtualWalletUsecase) getVirtualWallet(ctx context.Context, userId int64) (*VirtualWallet, error) {
	wallet, err := uc.repo.GetWallet(ctx, userId)
	if err != nil {
		switch err.(type) {
		// 未找到钱包，需要创建
		case *ent.NotFoundError:
			createWallet, createErr := uc.repo.CreateWallet(ctx, &VirtualWallet{UserId: userId})
			if createErr != nil {
				return nil, createErr
			} else {
				return createWallet, nil
			}
		default:
			return nil, err
		}
	}
	return wallet, nil
}

// compareTo, 0为相等 1为a大于b -1为a小于b
func compareTo(a ,b float64) int8 {
	accuracy := 0.000001
	if math.Abs(a-b) < accuracy {
		return 0
	}
	if math.Max(a, b) == a {
		return 1
	} else {
		return -1
	}
}