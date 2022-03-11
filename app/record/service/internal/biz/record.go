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
	Reward(ctx context.Context, userId int64, index int) error
	GetSignInRecord(ctx context.Context, userId int64, dates []string) ([]*Record, error)
}

type RecordUseCase struct {
	repo RecordRepo
	log  *log.Helper
}

func NewRecordUseCase(repo RecordRepo, logger log.Logger) *RecordUseCase {
	return &RecordUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *RecordUseCase) GetSignInInfo(ctx context.Context, userId int64) (*SignInResp, error) {
	ntime := time.Now()
	today := getDate2Str(ntime)
	todaySignState := false

	index, err := uc.getTodayIndex(ctx, userId, ntime)
	if err != nil {
		// 这里的error是获取记录的(查询库表）的unknown error
		return nil, err
	}
	_, err = uc.repo.GetSignInRecord(ctx, userId, []string{today})
	if err != nil {
		switch err.(type) {
		case *ent.NotFoundError:
			todaySignState = false
		default:
			// 这里的error是获取记录的(查询库表）的unknown error
			return nil, err
		}
	} else {
		todaySignState = true
	}
	if !todaySignState {
		index--
	}
	// index == -1时需要单独处理
	return uc.repo.GetSignInInfo(ctx, index, todaySignState)
}

func (uc *RecordUseCase) SignIn(ctx context.Context, userId int64) (*SignInResp, error) {
	ntime := time.Now()
	today := getDate2Str(ntime)

	index, err := uc.getTodayIndex(ctx, userId, ntime)
	if err != nil {
		// 这里的error是获取记录的(查询库表）的unknown error
		return nil, err
	}

	reply, err := uc.repo.SignIn(ctx, &Record{
		UserId:      userId,
		SignInIndex: index,
		SignInDay:   today,
	})
	if err != nil {
		// 写入record的error，记录已存在
		return nil, err
	}
	// heat job for increasing the balance of the virtual wallet
	err = uc.repo.Reward(ctx, userId, index)
	if err != nil {
		// 异步发奖的error，发消息给kafka失败
		return nil, err
	}
	return reply, nil
}

func (uc *RecordUseCase) getTodayIndex(ctx context.Context, userId int64, now time.Time) (int, error) {
	yesterday := getDate2Str(now.AddDate(0, 0, -1))
	index := 0
	// 获取昨天签到记录
	record, err := uc.repo.GetSignInRecord(ctx, userId, []string{yesterday})
	if err != nil {
		switch err.(type) {
		// 未找到昨天的记录，签到周期重置，SignIndex = 0
		case *ent.NotFoundError:
		default:
			return -1, err
		}
	} else {
		index = (record[0].SignInIndex + 1) % 7
	}
	return index, nil

}

func (uc *RecordUseCase) GetSignInRecord(ctx context.Context, userId int64,
	dates []string) ([]*Record, error) {
	return uc.repo.GetSignInRecord(ctx, userId,dates)
}

func getDate2Str(curTime time.Time) string {
	dateStr := curTime.Format("20060102")
	return dateStr
}