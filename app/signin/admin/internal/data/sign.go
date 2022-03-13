package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	recordV1 "sign-in/api/record/v1"
	"sign-in/app/signin/admin/internal/biz"
)

type userRecordRepo struct {
	data *Data
	log  *log.Helper
}

var _ biz.UserRecordRepo = (*userRecordRepo)(nil)

func NewUserRecordRepo(data *Data, logger log.Logger) biz.UserRecordRepo {
	return &userRecordRepo{data: data, log: log.NewHelper(logger)}
}

func (r userRecordRepo) GetUserSignInRecord(ctx context.Context, userId int64,
	days []string) ([]*biz.SignInRecord, error) {
	reply, err := r.data.rc.GetUserSignInRecord(ctx, &recordV1.GetUserSignInRecordRequest{
		User: userId,
		Day: days,
	})
	if err != nil {
		return nil, err
	}
	// if reply == nil ||
	return bulk2BizSignInRecord(reply.UserRecord), nil
}

func bulk2BizSignInRecord(source []*recordV1.UserRecord) []*biz.SignInRecord {
	record := make([]*biz.SignInRecord, 0 ,len(source))
	for _ ,v := range source {
		record = append(record, &biz.SignInRecord{
			Day:    v.Day,
			Index: int(v.Index),
			Reward: v.Reward,
		})
	}
	return record
}

