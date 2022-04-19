package config

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	//mysql driver
	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (cfg Config) MysqlConnect() (*sql.DB, error) {
	dbConnString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.db.User,
		cfg.db.Password,
		cfg.db.Host,
		cfg.db.Port,
		cfg.db.Name,
	)

	db, err := sql.Open("mysql", dbConnString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func (cfg Config) MongoConnect() (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	clientOptions := options.Client()

	connString := fmt.Sprintf("mongodb://%s:%s", cfg.db.Host, cfg.db.Port)
	clientOptions.ApplyURI(connString)

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	if err := client.Connect(ctx); err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client.Database(cfg.db.Name), nil
}
