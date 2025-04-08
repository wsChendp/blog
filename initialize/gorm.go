package initialize

import (
	"server/global"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitGorm() *gorm.DB {
	mysqlConfig := global.Config.Mysql

	db, err := gorm.Open(mysql.Open(mysqlConfig.Dsn()), &gorm.Config{
		Logger: logger.Default.LogMode(mysqlConfig.LogLevel()),
	})
	if err != nil {
		global.Log.Error("数据库连接失败", zap.Error(err))
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(mysqlConfig.MaxIdleConns)
	sqlDB.SetMaxOpenConns(mysqlConfig.MaxOpenConns)

	return db
}
