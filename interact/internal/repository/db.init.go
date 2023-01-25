package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"strings"
	"time"
)

var DB *gorm.DB

func InitDB() {
	host := viper.GetString("mysql.host")
	port := viper.GetString("mysql.port")
	database := viper.GetString("mysql.database")
	username := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	charset := viper.GetString("mysql.charset")
	dsn := strings.Join(
		[]string{username, ":", password, "@tcp(", host, ":", port, ")/", database, "?charset=", charset, "&parseTime=true"}, "")
	err := Database(dsn)
	if err != nil {
		panic(err)
	}
}

// gorm的定义
func Database(dsn string) error {
	var ormLogger logger.Interface
	// 不同环境，输出日志不同
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,  // 禁用datetime的经度(mysql5.6前不支持)
		DontSupportRenameIndex:    true,  // 重命名索引的时候采用删除并新建的方式(mysql5.7之前不支持重命名索引)
		DontSupportRenameColumn:   true,  // 用 change 重命名列 (mysql8之前的数据库不支持重命名列)
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 单数蛇形命名策略
		},
	})
	if err != nil {
		return err
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)                  // 设置空闲连接池
	sqlDB.SetMaxOpenConns(100)                 // 设置最大连接数
	sqlDB.SetConnMaxLifetime(time.Second * 30) // 设置连接最长存在时间
	DB = db
	migration()
	return err
}
