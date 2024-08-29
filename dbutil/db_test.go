package dbutil

import (
	"testing"

	"google.golang.org/protobuf/types/known/durationpb"

	"github.com/lastares/claymore/generalutil"
	"github.com/lastares/claymore/protobuf/conf"
)

func TestNewDB(t *testing.T) {
	dsn := "root:root@tcp(127.0.0.1:3306)/my_test?charset=utf8mb4&parseTime=True&loc=Local"
	databaseConf := &conf.Data_Database{
		Driver:             "mysql",
		Source:             dsn,
		MaxOpenConnections: 100,
		MaxIdleConnections: 100,
		ConnectionLifeTime: &durationpb.Duration{
			Seconds: 3600,
		},
	}
	app := &conf.App{
		Name: "test",
		Env:  "dev",
	}
	gormConfig := GormConfig(app)
	gorm, err := New(databaseConf, gormConfig)
	if err != nil {
		t.Errorf("NewDB error: %v", err)
	}
	type User struct {
		ID   int32
		Name string
	}
	var u *User
	if err = gorm.Table("user").Select("*").First(&u).Error; err != nil {
		t.Errorf("NewDB error: %v", err)
	}
	generalutil.PrettyPrintStruct(u)
}
