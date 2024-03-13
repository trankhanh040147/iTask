package mysqlprovider

import (
	"fmt"
	"gorm.io/gorm/logger"
	"iTask/config"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQL(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.Mysql.User, cfg.Mysql.Password,
		cfg.Mysql.Host, cfg.Mysql.Port, cfg.Mysql.DBName)

	log := logger.Default.LogMode(logger.Info)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: log,
	})
	if err != nil {
		return nil, err
	}

	// Connection Pool Settings
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(cfg.Mysql.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.Mysql.MaxIdleConns)
	if cfg.Mysql.MaxConnLifetime == "hour" {
		sqlDB.SetConnMaxLifetime(time.Hour)
	}

	return db, nil
}
