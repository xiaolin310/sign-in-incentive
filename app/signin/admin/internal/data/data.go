package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	recordV1 "sign-in/api/record/v1"
	"sign-in/app/signin/admin/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewRecordServiceClient, NewUserRecordRepo)

// Data .
type Data struct {
	log *log.Helper
	rc  recordV1.RecordServiceClient
}

// NewData .
func NewData(rc recordV1.RecordServiceClient, logger log.Logger) (*Data, error) {
	log := log.NewHelper(log.With(logger, "module", "signin-admin/data"))
	d := &Data{
		log: log,
		rc:  rc,
	}
	return d, nil
}

func NewRecordServiceClient(c *conf.Data) recordV1.RecordServiceClient {
	conn, err := grpc.Dial(
		c.RecordClient.Addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err.Error())
	}
	return recordV1.NewRecordServiceClient(conn)
}
