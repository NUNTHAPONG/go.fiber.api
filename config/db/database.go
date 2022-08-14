package db

import (
	"fmt"

	"web-api/config/env"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var Context *gorm.DB

func ConnectDb() {
	e := env.Db()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=Asia/Bangkok",
		e.HOST,
		e.USERNAME,
		e.PASSWORD,
		e.DATABASE,
		e.PORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: e.SCHEMA + ".",
		},
		Logger: logger.Default.LogMode(logger.Info),
		DryRun: true,
	})
	if err != nil {
		panic("Failed to create a connection to database")
	}
	Context = db
	fmt.Println("Database is Connected.")
}

func DisconnectDb() {
	conn, err := Context.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	conn.Close()
	fmt.Println("Database is Closed.")
}

func Create(model interface{}, req interface{}) error {
	tx := Context.Model(model).Omit("Xmin", "State").Create(req)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func Read(model interface{}, req interface{}, condition interface{}) error {
	tx := Context.Model(model).Omit("State").Order("1").Where(condition).Find(req)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func Update(model interface{}, req interface{}, condition interface{}) error {
	tx := Context.Model(model).Omit("CrBy", "CrSys", "CrDate", "Xmin", "State").Where(condition).Save(req)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func Delete(model interface{}, req interface{}, condition interface{}) error {
	tx := Context.Model(model).Where(condition).Delete(req)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}