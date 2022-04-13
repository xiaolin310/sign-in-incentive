package data

import (
	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	recordv1 "sign-in/api/record/v1"
	userv1 "sign-in/api/user/v1"
	walletClientV1 "sign-in/api/virtualwallet/v1"
	"sign-in/app/signin/interface/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewUserRepo,
	NewUserClient,
	NewRecordClient,
	NewRecordRepo,
	NewWalletClient,
	NewWalletRepo)

// Data .
type Data struct {
	log *log.Helper
	uc 	userv1.UserServiceClient
	rc  recordv1.RecordServiceClient
	wc  walletClientV1.VirtualWalletClient
}

// NewData .
func NewData(
	c *conf.Data,
	logger log.Logger,
	uc userv1.UserServiceClient,
	rc recordv1.RecordServiceClient,
	wc walletClientV1.VirtualWalletClient,
) (*Data, error) {
	return &Data{
		log: log.NewHelper(log.With(logger, "module", "data")),
		uc:  uc,
		rc:  rc,
		wc:  wc,
	}, nil
}

func NewUserClient(config *conf.Data) userv1.UserServiceClient {
	conn, err := grpc.Dial(
		config.Userclient.Addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err.Error())
	}
	return userv1.NewUserServiceClient(conn)
}

func NewRecordClient(config *conf.Data) recordv1.RecordServiceClient {
	conn, err := grpc.Dial(
		config.Recordclient.Addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err.Error())
	}
	return recordv1.NewRecordServiceClient(conn)
}

func NewWalletClient(config *conf.Data) walletClientV1.VirtualWalletClient {
	conn, err := grpc.Dial(
		config.Walletclient.Addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err.Error())
	}
	return walletClientV1.NewVirtualWalletClient(conn)
}
//func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
//	cleanup := func() {
//		log.NewHelper(logger).Info("closing the data resources")
//	}
//	return &Data{}, cleanup, nil
//}
