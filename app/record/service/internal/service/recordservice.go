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

	uc  *biz.RecordUsecase
	log *log.Helper
}

var _ v1.RecordServiceServer = (*RecordService)(nil)

func NewRecordService(uc *biz.RecordUsecase, logger log.Logger) *RecordService {
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
	}
	return toProtoSignInInfoReply(resp), err
}
func (s *RecordService) SignIn(ctx context.Context, req *v1.SignInRequest) (*v1.SignInInfoReply, error) {
	resp, err := s.uc.SignIn(ctx, req.User)
	if err != nil {
		switch err.(type) {
		case *ent.ConstraintError:
			err = v1.ErrorRecordExisted("")
		default:
			err = v1.ErrorUnknownError("")
		}
	}
	return toProtoSignInInfoReply(resp), err
}

func toProtoSignInInfoReply(source *biz.SignInResp) *v1.SignInInfoReply {
	return &v1.SignInInfoReply {
		RewardList:    source.RewardList,
		SignIndexList: source.SignInIndexList,
		IsSignToday:   source.IsSignInToday,
	}
}
