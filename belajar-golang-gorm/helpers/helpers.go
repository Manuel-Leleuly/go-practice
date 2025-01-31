package helpers

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetGORMDatabaseUrl(params map[string]string) string {
	const dbUrl = "root:root@tcp(127.0.0.1:3306)/belajar_golang_gorm"
	queryParams := ""

	for k, v := range params {
		selectedParam := k + "=" + v
		if queryParams == "" {
			queryParams = "?" + selectedParam
		} else {
			queryParams += "&" + selectedParam
		}
	}

	return dbUrl + queryParams
}

func OpenConnection() *gorm.DB {
	dialect := mysql.Open(GetGORMDatabaseUrl(map[string]string{
		"charset":   "utf8mb4",
		"parseTime": "True",
		"loc":       "Local",
	}))

	db, err := gorm.Open(dialect, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)

	return db
}
