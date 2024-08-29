package dbutil

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/lastares/claymore/protobuf/conf"
)

func New(dbConfig *conf.Data_Database, gormConfig gorm.Config) (*gorm.DB, error) {
	mysqlConfig := InitConfig(
		WithDriver(dbConfig.Driver),
		DSN(dbConfig.Source),
		DefaultStringSize(256), // 为字符串(string)字段设置大小。默认情况下，对于没有大小、没有主键、没有定义索引且没有默认值的字段，将使用db类型“longext”

		DisableDatetimePrecision(true),  // 禁用日期时间精度支持。但是这在MySQL 5.6之前不支持
		DontSupportRenameColumn(true),   // 重命名列时使用change,但是在MySQL 8、MariaDB之前不支持重命名
		DontSupportRenameIndex(true),    // 重命名索引时删除并创建索引。但是在MySQL 5.7、MariaDB之前不支持重命名索引
		SkipInitializeWithVersion(true), // 是否根据当前 MySQL 版本自动配置
	)
	db, err := gorm.Open(mysql.New(mysqlConfig), &gormConfig)
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return db, err
	}
	// 连接池配置
	sqlDB.SetConnMaxLifetime(dbConfig.ConnectionLifeTime.AsDuration()) // 每一个连接的生命周期
	sqlDB.SetMaxIdleConns(int(dbConfig.MaxIdleConnections))            // 是设置空闲时的最大连接数
	sqlDB.SetMaxOpenConns(int(dbConfig.MaxOpenConnections))            // 设置与数据库的最大打开连接数
	return db, err
}

// GormConfig 配置GORM的日志和事务设置。
// 参数:
//
//	app - 提供应用配置，用于判断当前环境以设置日志级别。
//
// 返回值:
//
//	返回一个gorm.Config对象，用于配置GORM的行为。
func GormConfig(app *conf.App) gorm.Config {
	// 根据应用环境设置日志模式，默认为警告级别
	logMode := logger.Warn
	// 如果是开发环境，切换到信息级别日志，记录所有SQL执行
	if app.IsDevelopment() {
		logMode = logger.Info
	}
	// 创建一个新的logger实例，配置GORM日志行为
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer，将日志输出到标准输出
		logger.Config{
			SlowThreshold:             1 * time.Second, // 慢查询阈值设置为1秒
			LogLevel:                  logMode,         // 设置日志级别
			IgnoreRecordNotFoundError: true,            // 忽略记录未找到错误
		},
	)
	// 返回配置的GORM配置对象
	return gorm.Config{
		Logger:                 newLogger, // 使用新创建的logger实例
		SkipDefaultTransaction: true,      // 跳过默认事务，通常用于提高性能
		PrepareStmt:            true,      // 启用预处理语句，可以提高性能和安全性
	}
}
