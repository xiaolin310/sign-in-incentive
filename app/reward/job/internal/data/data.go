package data

import (
	"github.com/Shopify/sarama"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	walletClientV1 "sign-in/api/virtualwallet/v1"
	"sign-in/app/reward/job/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewWalletClient, NewKafkaConsumer, NewRewardRepo)

// Data .
type Data struct {
	log *log.Helper
	kc  sarama.Consumer
	wc  walletClientV1.VirtualWalletClient
}

// NewData .
func NewData(consumer sarama.Consumer, wc walletClientV1.VirtualWalletClient, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "reward-job/data"))

	d := &Data{
		log: log,
		kc:  consumer,
		wc:  wc,
	}
	return d, func() {
		if err := d.kc.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}

func NewWalletClient(c *conf.Data) walletClientV1.VirtualWalletClient {
	conn, err := grpc.Dial(
		c.WalletClient.Addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err.Error())
	}
	return walletClientV1.NewVirtualWalletClient(conn)

}

func NewKafkaConsumer(c *conf.Data) sarama.Consumer {
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumer(c.Kafka.Addr, config)
	if err != nil {
		panic(err.Error())
	}
	return consumer
}
