package database

import (
	"enroll/utils"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbConnection struct {
	Connection *gorm.DB
}

var (
	Database DbConnection
)

func (dbConnection *DbConnection) InitConnection() error {
	var dsnPostgres string
	dsnPostgres = fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		utils.ConfigurationEnvs.DatabaseHost,
		utils.ConfigurationEnvs.DatabaseUser,
		utils.ConfigurationEnvs.DatabasePass,
		utils.ConfigurationEnvs.DatabaseName,
		utils.ConfigurationEnvs.DatabasePort,
		utils.ConfigurationEnvs.DatabaseSSLMode,
	)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)
	db, err := gorm.Open(postgres.Open(dsnPostgres), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return err
	}
	dbConnection.Connection = db
	return nil
}
