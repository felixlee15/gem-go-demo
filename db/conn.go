package db

import (
	"context"
	"database/sql"
	"sync"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/sirupsen/logrus"

	"gemdemo/config"
	"gemdemo/ent"
	"gemdemo/ent/migrate"
)

var (
	client   *ent.Client
	clientMU sync.Mutex
	driverMU sync.Mutex
)

func Init(ctx context.Context) {
	clientMU.Lock()
	defer clientMU.Unlock()
	client = newClient()
	if err := client.Schema.Create(ctx, migrate.WithDropIndex(true)); err != nil {
		logrus.Fatalf("migrate database schema: %s", err)
	}
	logrus.Info("Database connected successfully")
}

func newClient() *ent.Client {
	var err error
	dbConfig := config.AppConfigs.Database
	db, err := sql.Open("postgres", dbConfig.DSN())
	if err != nil {
		logrus.Fatalf("failed to connect to db: %v", err)
	}
	drv := entsql.OpenDB(dialect.Postgres, db)
	driverMU.Lock()
	defer driverMU.Unlock()
	return ent.NewClient(ent.Driver(drv))
}

func Close() {
	clientMU.Lock()
	defer clientMU.Unlock()
	if client != nil {
		_ = client.Close()
	}
}

// GetClient retrieves ent client from context; otherwise, creates new one.
func GetClient(ctx context.Context) *ent.Client {
	if ent.FromContext(ctx) != nil {
		return ent.FromContext(ctx)
	}

	if ent.TxFromContext(ctx) != nil {
		return ent.TxFromContext(ctx).Client()
	}

	clientMU.Lock()
	defer clientMU.Unlock()
	if client == nil {
		client = newClient()
	}
	return client
}
