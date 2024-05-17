package db

import (
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"strings"
	"time"
	"wxapp/config"
)

var _db *gorm.DB

func LoadMysql() {
	Mconf := config.GetConfig().Mysql
	path := strings.Join([]string{Mconf.UserName, ":", Mconf.Password, "@tcp(", Mconf.Host, ":", Mconf.Port, ")/", Mconf.DbName, "?charset=", Mconf.Charset, "&parseTime=", Mconf.ParseTime, "&loc=", Mconf.Loc}, "")
	dialector := mysql.New(mysql.Config{
		DSN:                           path,
		DefaultStringSize:             256,
		DisableDatetimePrecision:      true,
		DontSupportRenameColumnUnique: true,
		DontSupportRenameIndex:        true,
		SkipInitializeWithVersion:     false,
	})

	db, err := gorm.Open(dialector, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "wxapp_",
			SingularTable: true,
		},
		Logger:      GetSqlLogger(),
		PrepareStmt: true,
	})
	if err != nil {
		panic(fmt.Errorf("mysql connect error: %s", err))
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(20)  // 设置连接池，空闲
	sqlDB.SetMaxIdleConns(100) // 打开
	sqlDB.SetConnMaxLifetime(time.Second * 30)

	_db = db

	err = migrate()
	if err != nil {
		panic(fmt.Errorf("mysql migrate error: %s", err))

	}
	fmt.Println("mysql connect success...")
}

func NewDBClient(ctx context.Context) *gorm.DB {
	db := _db
	return db.WithContext(ctx)
}

func GetSqlLogger() logger.Interface {
	env := strings.ToLower(config.GetConfig().System.Mode)
	if env == "debug" {
		return logger.Default.LogMode(logger.Info)
	}
	return logger.Default
}

func migrate() error {
	_db.Set("gorm:table_options", "charset=utf8mb4")
	_db.Set("gorm:table_options", "ENGINE=InnoDB")
	return _db.AutoMigrate()
}
