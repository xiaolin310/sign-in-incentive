	package biz

	import (
		"context"
		"github.com/go-kratos/kratos/v2/log"
		"golang.org/x/sync/errgroup"
	)

	type SignInRecord struct {
		Day    string
		Index  int
		Reward float64
	}

	type UserRecord struct {
		UserId     int64
		SignRecord []*SignInRecord
	}



	type UserRecordRepo interface {
		GetUserSignInRecord(ctx context.Context, userId int64, days []string) ([]*SignInRecord, error)
	}

	type UserRecordUseCase struct {
		repo UserRecordRepo
		log  *log.Helper
	}

	func NewUserRecordUseCase(repo UserRecordRepo, logger log.Logger) *UserRecordUseCase {
		return &UserRecordUseCase{repo: repo, log: log.NewHelper(logger)}
	}

	func (uc *UserRecordUseCase) GetUserRecord(ctx context.Context, userIds []int64,
		days []string) ([]*UserRecord, error) {
		record := make([]*UserRecord, len(userIds))
		g, ctx := errgroup.WithContext(ctx)
		for k, v := range userIds {
			index := k
			uid := v
			g.Go(func() error {
				reply, err := uc.repo.GetUserSignInRecord(ctx, uid, days)
				if err != nil {
					return err
				}
				record[index] = &UserRecord{
					UserId: uid,
					SignRecord: reply,
				}
				return nil
			})
		}
		if err := g.Wait(); err != nil {
			return nil, err
		}
		return record, nil
	}
