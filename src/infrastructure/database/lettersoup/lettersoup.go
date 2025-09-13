package database

import (
	contractsProviders "lettersoup/src/application/contracts/providers"
	initdb "lettersoup/src/infrastructure/database/lettersoup/init_db"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type lettersoupDB struct{}

var DB *gorm.DB

func (ggsbd lettersoupDB) SetUp(host string, port string, user string, password string, dbname string, ssl *bool, logger contractsProviders.ILoggerProvider) {
	var sslmode string
	if ssl != nil && *ssl {
		logger.Info("SSL is enabled")
		sslmode = "require"
	} else {
		logger.Info("SSL is disabled")
		sslmode = "disable"
	}
	dsn := "host=" + host + " port=" + port + " user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=" + sslmode
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Panic("Error connecting to database", err)
	}
	DB = db
	logger.Info("Database connection established")
	initdb.InitMigrate(db, logger)
}

var Lettersoupdb *lettersoupDB

func init() {
	Lettersoupdb = &lettersoupDB{}
}
