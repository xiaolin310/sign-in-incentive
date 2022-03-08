package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"sign-in/app/record/service/internal/data/ent"
	"time"
)

type Record struct {
	Id          int64
	UserId      int64
	SignInIndex int
	Reward      float64
	SignInDay   string
}

type SignInResp struct {
	RewardList        []float64
	SignInIndexList   []bool
	IsSignInToday     bool
}

type RecordRepo interface {
	// GetSignInInfo 整个7天周期内，index之前都是已签到状态，包含index，index的值从0到6
	GetSignInInfo(ctx context.Context, index int, todayState bool) (*SignInResp, error)
	SignIn(ctx context.Context, record *Record) (*SignInResp, error)
	Reward(ctx context.Context, userId int64) error
	GetTodaySignInRecord(ctx context.Context, userId int64, today string) (*Record, error)
	GetYesterdaySignInDex(ctx context.Context, userId int64, yesDate string) (int, error)
}

type RecordUsecase struct {
	repo RecordRepo
	log  *log.Helper
}

func NewRecordUsecase(repo RecordRepo, logger log.Logger) *RecordUsecase {
	return &RecordUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *RecordUsecase) GetSignInInfo(ctx context.Context, userId int64) (*SignInResp, error) {
	ntime := time.Now()
	today := getDate2Str(ntime)
	todaySignState := false

	index, err := uc.getTodayIndex(ctx, userId, ntime)
	if err != nil {
		return nil, err
	}
	_, err = uc.repo.GetTodaySignInRecord(ctx, userId, today)
	if err != nil {
		switch err.(type) {
		case *ent.NotFoundError:
			todaySignState = false
		default:
			return nil, err
		}
	} else {
		todaySignState = true
	}
	if !todaySignState {
		index--
	}
	// index == -1需要单独处理
	return uc.repo.GetSignInInfo(ctx, index, todaySignState)
}

func (uc *RecordUsecase) SignIn(ctx context.Context, userId int64) (*SignInResp, error) {
	ntime := time.Now()
	today := getDate2Str(ntime)

	index, err := uc.getTodayIndex(ctx, userId, ntime)
	if err != nil {
		return nil, err
	}

	reply, err := uc.repo.SignIn(ctx, &Record{
		UserId:      userId,
		SignInIndex: index,
		SignInDay:   today,
	})
	if err != nil {
		return nil, err
	}
	// heat job for increasing the balance of the virtual wallet
	err = uc.repo.Reward(ctx, userId)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (uc *RecordUsecase) getTodayIndex(ctx context.Context, userId int64, now time.Time) (int, error) {
	yesterday := getDate2Str(now.AddDate(0, 0, -1))
	// 获取昨天签到记录
	index, err := uc.repo.GetYesterdaySignInDex(ctx, userId, yesterday)
	if err != nil {
		switch err.(type) {
		case *ent.NotFoundError:
			index = 0
		default:
			return -1, err
		}
	} else {
		index = (index+1) % 7
	}
	return index, nil

}

func getDate2Str(curTime time.Time) string {
	dateStr := curTime.Format("20060102")
	return dateStr
}