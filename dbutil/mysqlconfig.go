package dbutil

import (
	"gorm.io/driver/mysql"
)

type Config func(*mysql.Config)

// WithDSN data source name
func WithDSN(dsn string) Config {
	return func(c *mysql.Config) {
		c.DSN = dsn
	}
}

func WithDriver(driver string) Config {
	return func(c *mysql.Config) {
		c.DriverName = driver
	}
}

// DefaultStringSize string 类型字段的默认长度
func DefaultStringSize(size uint) Config {
	return func(c *mysql.Config) {
		c.DefaultStringSize = size
	}
}

// DisableDatetimePrecision 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
func DisableDatetimePrecision(disabled bool) Config {
	return func(c *mysql.Config) {
		c.DisableDatetimePrecision = disabled
	}
}

// DontSupportRenameIndex 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
func DontSupportRenameIndex(dontSupport bool) Config {
	return func(c *mysql.Config) {
		c.DontSupportRenameIndex = dontSupport
	}
}

// DontSupportRenameColumn 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
func DontSupportRenameColumn(dontSupport bool) Config {
	return func(c *mysql.Config) {
		c.DontSupportRenameColumn = dontSupport
	}
}

// SkipInitializeWithVersion 根据当前 MySQL 版本自动配置
func SkipInitializeWithVersion(skipped bool) Config {
	return func(c *mysql.Config) {
		c.SkipInitializeWithVersion = skipped
	}
}

func InitConfig(cfgs ...Config) mysql.Config {
	config := mysql.Config{}
	for _, cfg := range cfgs {
		cfg(&config)
	}
	return config
}
