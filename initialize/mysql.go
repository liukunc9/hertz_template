package initialize

import (
	"fmt"
	"github.com/liukunc9/hertz_template/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

func initDB() {
	mysqlCfg := global.Config.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		mysqlCfg.Username,
		mysqlCfg.Password,
		mysqlCfg.Host,
		mysqlCfg.Port,
		mysqlCfg.Dbname)

	MysqlConfig := mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,   // string 类型字段默认长度
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}

	db, err := gorm.Open(mysql.New(MysqlConfig), getGormConfig())

	if err != nil {
		panic(err)
	}
	global.DB = db
}

func getGormConfig() *gorm.Config {
	mysqlCfg := global.Config.Mysql
	config := new(gorm.Config)
	namingStrategy := schema.NamingStrategy{
		SingularTable: true,
	}
	if len(mysqlCfg.TablePrefix) > 0 {
		namingStrategy.TablePrefix = mysqlCfg.TablePrefix
	}

	_default := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold: 200 * time.Millisecond,
		LogLevel:      logger.Warn,
		Colorful:      true,
	})
	switch mysqlCfg.LogMode {
	case "silent", "Silent":
		config.Logger = _default.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = _default.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = _default.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = _default.LogMode(logger.Info)
	default:
		config.Logger = _default.LogMode(logger.Info)
	}

	return config
}
