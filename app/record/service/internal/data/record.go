package data

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/go-kratos/kratos/v2/log"
	"sign-in/app/record/service/internal/biz"
	"sign-in/app/record/service/internal/data/ent"
	"sign-in/app/record/service/internal/data/ent/record"
	"strconv"
	"strings"
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
	// kafka topic
	topic = "reward"
	// ErrWorkTimeNotFoundInCache record not found in cache
	ErrWorkTimeNotFoundInCache = errors.New("error: user record not existed in cache")

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

func (r *recordRepo) Reward(ctx context.Context, userId int64, index int) error {
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.ByteEncoder(buildMessage(userId, rewardList[index]))

	partition, offset, err := r.data.kafka.SendMessage(msg)
	if err != nil {
		r.log.Errorf("Send Message to Kafka failed, err: %v", err)
		return err
	}
	r.log.Infof("Producer message sent to Kafka, partition:%v offset:%v\n", partition, offset)
	return nil
}

func (r *recordRepo) GetSignInRecord(ctx context.Context, userId int64, dates []string) ([]*biz.Record, error) {
	result, err := r.getUserRecordFromCache(ctx, userId)
	if err != nil {
		result, err = r.getUserRecordFromDB(ctx, userId, dates)
		if err != nil {
			return nil, err
		}
		return result, r.writeUserRecord2Cache(ctx, userId, result)
	}
	return result, nil
}

func (r *recordRepo) getUserRecordFromDB(ctx context.Context, userId int64,
	dates []string) ([]*biz.Record, error) {
	result, err := r.data.db.Record.Query().
		Where(
			record.UserIDEQ(userId),
			record.SignInDayIn(dates...),
		).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return bulk2BizRecord(result), nil

}

func (r *recordRepo) getUserRecordFromCache(ctx context.Context, userId int64) ([]*biz.Record, error) {
	data, err := r.data.rc.Get(ctx, strconv.Itoa(int(userId))).Result()
	if err != nil {
		return nil, ErrWorkTimeNotFoundInCache
	}
	b := []byte(data)
	var recordList []*biz.Record
	err = json.Unmarshal(b, &recordList)
	if err != nil {
		return nil, err
	}
	return recordList, nil

}

func (r *recordRepo) writeUserRecord2Cache(ctx context.Context, userId int64,
	record []*biz.Record) error {
	data, err := json.Marshal(record)
	if err != nil {
		return err
	}
	return r.data.rc.Set(ctx, strconv.Itoa(int(userId)), data, 0).Err()
}


func bulk2BizRecord(source []*ent.Record) []*biz.Record {
	res := make([]*biz.Record, 0, len(source))
	for _, v := range source {
		res = append(res, toBizRecord(v))
	}
	return res

}

func toBizRecord(source *ent.Record) *biz.Record {
	return &biz.Record{
		Id:          int64(source.ID),
		UserId:      source.UserID,
		SignInIndex: source.SignInIndex,
		Reward:      source.Reward,
		SignInDay:   source.SignInDay,
	}
}

func buildMessage(userId int64, amount float64) string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("%d-%.2f", userId, amount))
	return builder.String()

}