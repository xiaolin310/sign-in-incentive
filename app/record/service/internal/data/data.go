package data

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"sign-in/app/record/service/internal/conf"
	"sign-in/app/record/service/internal/data/ent"
	"sign-in/app/record/service/internal/data/ent/migrate"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewRecordRepo, NewKafkaProducer, NewEntClient)

// Data .
type Data struct {
	log   *log.Helper
	db    *ent.Client
	kafka sarama.SyncProducer
	rc    *redis.Client
}

// NewData .
func NewData(c *conf.Data, entClient *ent.Client, producer sarama.SyncProducer, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "record-service/data"))

	rc := redis.NewClient(&redis.Options{
		Addr: c.Redis.Addr,
		DB:   0,
	})

	d := &Data{
		db:    entClient,
		log:   log,
		kafka: producer,
		rc:    rc,
	}

	return d, func() {
		if err := d.kafka.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}


func NewEntClient(conf *conf.Data, logger log.Logger) *ent.Client {
	log := log.NewHelper(log.With(logger, "module", "record-service/data/ent"))

	client, err := ent.Open(
		conf.Database.Driver,
		conf.Database.Source,
	)
	if err != nil {
		log.Fatalf("failed opening connection to db: %v", err)
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background(), migrate.WithDropIndex(true), migrate.WithDropColumn(true)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}

func NewKafkaProducer(c *conf.Data) sarama.SyncProducer {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	client, err := sarama.NewSyncProducer(c.Kafka.Addr, config)
	if err != nil {
		panic(err.Error())
	}
	//defer client.Close()
	return  client
}


// NewData .
//func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
//	cleanup := func() {
//		log.NewHelper(logger).Info("closing the data resources")
//	}
//	return &Data{}, cleanup, nil
//}
