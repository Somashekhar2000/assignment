package database

import (
	"assignment/graph/model"
	"context"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenDBConnection() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=1234 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Connection() (*gorm.DB, error) {
	db, err := OpenDBConnection()
	if err != nil {
		return nil, err
	}

	sdb, err := db.DB()
	if err != nil {
		return nil, err
	}

	ctx, cancle := context.WithTimeout(context.Background(), time.Second*5)
	defer cancle()

	err = sdb.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	err = db.Migrator().AutoMigrate(&model.User{})
	if err != nil {
		return nil, err
	}

	return db, nil

}
