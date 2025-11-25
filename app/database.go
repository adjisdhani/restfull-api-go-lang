package app

import (
	"belajar_golang_restful_api/helper"
	"database/sql"
	"fmt"
	"time"
)

func NewDB(configuration helper.Config) *sql.DB {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configuration.DBUser, configuration.DBPassword, configuration.DBHost, configuration.DBPort, configuration.DBName)
	sqlDB, err := sql.Open("mysql", dataSourceName)
	helper.PanicIfError(err)

	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxLifetime(10 * time.Minute)
	sqlDB.SetConnMaxIdleTime(10 * time.Minute)

	return sqlDB
}
