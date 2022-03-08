package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	recordv1 "sign-in/api/record/v1"
	"sign-in/app/signin/interface/internal/biz"
)

type recordRepo struct {
	data *Data
	log  *log.Helper
}

var _ biz.RecordRepo = (*recordRepo)(nil)

func NewRecordRepo(data *Data, logger log.Logger) biz.RecordRepo {
	return &recordRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *recordRepo) GetSignInInfo(ctx context.Context) (*biz.SignInData, error) {
	reply, err := r.data.rc.GetSignInInfo(ctx, &recordv1.GetSignInInfoRequest{
		User: int64(ctx.Value("user").(int)),
	})
	if err != nil {
		return nil, err
	}
	return toBizSignInData(reply), nil
}

func (r *recordRepo) SignIn(ctx context.Context) (*biz.SignInData, error) {
	reply, err := r.data.rc.SignIn(ctx, &recordv1.SignInRequest{
		User: int64(ctx.Value("user").(int)),
	})
	if err != nil {
		return nil, err
	}
	return toBizSignInData(reply), nil
}

func toBizSignInData(source *recordv1.SignInInfoReply) *biz.SignInData {
	return &biz.SignInData{
		RewardList:      source.RewardList,
		SignInIndexList: source.SignIndexList,
		IsSignToday:     source.IsSignToday,
	}
}

