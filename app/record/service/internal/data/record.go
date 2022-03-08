package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"sign-in/app/record/service/internal/biz"
	"sign-in/app/record/service/internal/data/ent"
	"sign-in/app/record/service/internal/data/ent/record"
	"time"
)

type recordRepo struct {
	data *Data
	log  *log.Helper
}

var _ biz.RecordRepo = (*recordRepo)(nil)

// var ErrRecordNotComplete = errors.New("error: user clockin records does not complete")

var (
	// TODO: 暂时固定7天奖励列表。后期可以考虑对接运营配置平台或者第三方平台
	rewardList = []float64{100.0, 200.0, 300.0, 400.0, 500.0, 600.0, 700.0}
	// 值为0或1，长度为7，用true代表当前周期内已经签到的那天，
	// 如[true,false,false,false,false,false,false]代表当前周期第一天已经签到
	//signInIndexList []bool
	// 代表当天是否已经签到，已签到值为true
	//isSignInToday bool
)

func NewRecordRepo(data *Data, logger log.Logger) biz.RecordRepo {
	return &recordRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *recordRepo) GetSignInInfo(ctx context.Context, index int, todayState bool) (*biz.SignInResp, error) {
	signInIndexList := make([]bool, 7)

	for i := 0; i <= index; i++ {
		signInIndexList[i] = true
	}

	return &biz.SignInResp{
		RewardList: rewardList,
		SignInIndexList: signInIndexList,
		IsSignInToday: todayState,
	}, nil

}

func (r *recordRepo) SignIn(ctx context.Context, record *biz.Record) (*biz.SignInResp, error) {
	signInIndexList := make([]bool, 7)
	for i := 0; i <= record.SignInIndex; i++ {
		signInIndexList[i] = true
	}

	_, err := r.data.db.Record.Create().
		SetUserID(record.UserId).
		SetSignInIndex(record.SignInIndex).
		SetReward(rewardList[record.SignInIndex]).
		SetSignInDay(record.SignInDay).
		SetCreatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &biz.SignInResp{
		RewardList: rewardList,
		SignInIndexList: signInIndexList,
		IsSignInToday: true,
	}, nil


}

func (r *recordRepo) Reward(ctx context.Context, userId int64) error {
	// TODO: 调用kafka
	return nil
}

func (r *recordRepo) GetTodaySignInRecord(ctx context.Context, userId int64, today string) (*biz.Record, error) {
	record, err := r.data.db.Record.Query().
		Where(
		    record.UserIDEQ(userId),
			record.SignInDayEQ(today),
			).
		First(ctx)
	if err != nil {
		return nil, err
	}
	return toBizRecord(record), nil
}

func (r *recordRepo) GetYesterdaySignInDex(ctx context.Context, userId int64, yesDate string) (int, error) {
	index, err := r.data.db.Record.Query().
		Where(
			record.UserIDEQ(userId),
			record.SignInDayEQ(yesDate),
			).
		Select(record.FieldSignInIndex).
		Int(ctx)

	if err != nil {
		return -1, err
	}
	return index, nil
}


func toBizRecord(source *ent.Record) *biz.Record {
	return &biz.Record{
		Id: int64(source.ID),
		UserId: source.UserID,
		SignInIndex: source.SignInIndex,
		Reward: source.Reward,
		SignInDay: source.SignInDay,
	}
}