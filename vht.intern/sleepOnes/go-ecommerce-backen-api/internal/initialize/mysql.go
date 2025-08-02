package initialize

import (
	"fmt"
	"go-ecommerce-backend-api/global"
	"go-ecommerce-backend-api/internal/po"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CheckPanicError(err error, message string) {
	if err != nil {
		global.Logger.Error(message, zap.Error(err))
		panic(err)
	}
}

func InitMySQL() {
	m := global.Config.MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.Username,
		m.Password,
		m.Host,
		m.Port,
		m.DbName,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	CheckPanicError(err, "InitialMysql initialization failed!")
	global.Logger.Info("InitialMysql successfully!", zap.Bool("success", true))

	global.Mdb = db

	SetPool()
	migrateTables()

}

func SetPool() {
	m := global.Config.MySQL
	sqlDB, err := global.Mdb.DB()
	CheckPanicError(err, "Set Pool failed!")

	sqlDB.SetMaxIdleConns(m.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime))
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)

}

func migrateTables() {
	err := global.Mdb.AutoMigrate(
		&po.User{},
		&po.Role{},
	)

	CheckPanicError(err, "Migration Tables failed!")

	global.Logger.Info("Migration Tables successfully!", zap.Bool("success", true))

}
