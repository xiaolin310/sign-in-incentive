package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"sign-in/app/signin/admin/internal/biz"

	v1 "sign-in/api/signin/admin/v1"
)

type SignInAdminService struct {
	v1.UnimplementedSignInAdminServiceServer

	uc *biz.UserRecordUseCase
	log *log.Helper
}

var _ v1.SignInAdminServiceServer = (*SignInAdminService)(nil)

func NewSignInAdminServiceService(uc *biz.UserRecordUseCase, logger log.Logger) *SignInAdminService {
	return &SignInAdminService{uc : uc, log: log.NewHelper(logger)}
}

func (s *SignInAdminService) GetSignInRecord(ctx context.Context, req *v1.GetSignInRecordRequest) (reply *v1.GetSignInRecordReply, err error) {
	record, err := s.uc.GetUserRecord(ctx, req.User, req.Day)
	switch  {
	case err != nil:
		err = v1.ErrorUnknownError(err.Error())
	case len(record) == 0:
		err = v1.ErrorUserRecordNotFound("error: user sign in record not existed")
	}
	if err == nil {
		reply = &v1.GetSignInRecordReply{
			Data: bulk2ProtoUserSignInRecord(record),
		}
	}
	return
}

func bulk2ProtoUserSignInRecord(source []*biz.UserRecord) []*v1.UserSignInRecord {
	userRecord := make([]*v1.UserSignInRecord, 0, len(source))
	for _, v := range source {
		userRecord = append(userRecord, &v1.UserSignInRecord{
			User: v.UserId,
			SignRecord: bulk2ProtoSignInRecord(v.SignRecord),
		})
	}
	return userRecord
}

func bulk2ProtoSignInRecord(source []*biz.SignInRecord) []*v1.SignInRecord {
	record := make([]*v1.SignInRecord, 0, len(source))
	for _, v := range source {
		record = append(record, &v1.SignInRecord{
			Day:       v.Day,
			SignIndex: int32(v.Index),
			Reward:    v.Reward,
		})
	}
	return record
}