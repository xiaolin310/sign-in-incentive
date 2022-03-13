package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"sign-in/app/record/service/internal/biz"
	"sign-in/app/record/service/internal/data/ent"

	v1 "sign-in/api/record/v1"
)

type RecordService struct {
	v1.UnimplementedRecordServiceServer

	uc  *biz.RecordUseCase
	log *log.Helper
}

var _ v1.RecordServiceServer = (*RecordService)(nil)

func NewRecordService(uc *biz.RecordUseCase, logger log.Logger) *RecordService {
	return &RecordService{uc: uc, log: log.NewHelper(logger)}}

func (s *RecordService) GetSignInInfo(ctx context.Context, req *v1.GetSignInInfoRequest) (*v1.SignInInfoReply, error) {
	resp, err := s.uc.GetSignInInfo(ctx, req.User)
	if err != nil {
		switch err.(type) {
		case *ent.NotFoundError:
			err = v1.ErrorRecordNotFound("error: get sign info failed, reason: record not found")
		default:
			err = v1.ErrorUnknownError("error: get sign info failed, reason: unknown error")
		}
		return nil, err
	}
	return toProtoSignInInfoReply(resp), err
}
func (s *RecordService) SignIn(ctx context.Context, req *v1.SignInRequest) (*v1.SignInInfoReply, error) {
	resp, err := s.uc.SignIn(ctx, req.User)
	if err != nil {
		switch err.(type) {
		case *ent.ConstraintError:
			err = v1.ErrorRecordExisted("error: sign in record existed")
		default:
			err = v1.ErrorUnknownError("error: unknown error")
		}
		return nil, err
	}
	return toProtoSignInInfoReply(resp), err
}

func (s *RecordService) GetUserSignInRecord(ctx context.Context,
	req *v1.GetUserSignInRecordRequest) (reply *v1.GetUserSignInRecordReply, err error) {
	res, err := s.uc.GetSignInRecord(ctx, req.User, req.Day)
	switch  {
	case err != nil:
		err = v1.ErrorUnknownError(err.Error())
	case len(res) == 0:
		err = v1.ErrorRecordNotFound("error: user sign in record not existed")
	}
	if err == nil {
		reply = &v1.GetUserSignInRecordReply{
			UserRecord: bulk2ProtoRecord(res),
		}
	}
	return
}

func toProtoSignInInfoReply(source *biz.SignInResp) *v1.SignInInfoReply {
	return &v1.SignInInfoReply {
		RewardList:    source.RewardList,
		SignIndexList: source.SignInIndexList,
		IsSignToday:   source.IsSignInToday,
	}
}

func bulk2ProtoRecord(source []*biz.Record) []*v1.UserRecord {
	result := make([]*v1.UserRecord, 0, len(source))
	for _, v := range source {
		v.SignInIndex++
		result = append(result, &v1.UserRecord{
			Day:    v.SignInDay,
			Index:  int32(v.SignInIndex),
			Reward: v.Reward,
		})
	}
	return result
}
