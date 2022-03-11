package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"sign-in/app/virtualwallet/service/internal/conf"
	"sign-in/app/virtualwallet/service/internal/data/ent"
	"sign-in/app/virtualwallet/service/internal/data/ent/migrate"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewWalletRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	db *ent.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, error) {
	client, err := ent.Open(
		c.Database.Driver,
		c.Database.Source,
	)
	if err != nil {
		panic(err.Error())
	}
	if err := client.Schema.Create(context.TODO(), migrate.WithDropIndex(true), migrate.WithDropColumn(true)); err != nil {
		panic(err.Error())
	}
	return &Data{
		db: client.Debug(),
	}, nil
}

//// NewData .
//func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
//	cleanup := func() {
//		log.NewHelper(logger).Info("closing the data resources")
//	}
//	return &Data{}, cleanup, nil
//}
