package global

import (
	"fmt"
	"github.com/spf13/viper"
	"go-easy-admin/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var (
	GORM *gorm.DB
	err  error
)

// 初始化数据库

func InitMysql() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("mysql.DbUser"),
		viper.GetString("mysql.DbPwd"),
		viper.GetString("mysql.DbHost"),
		viper.GetInt("mysql.DbPort"),
		viper.GetString("mysql.DbName"))
	GORM, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic("数据库连接失败" + err.Error())
	}
	if viper.GetInt("mysql.ActiveDebug") == 1 {
		GORM = GORM.Debug()
	}
	// 开启连接池
	db, _ := GORM.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	TPLogger.Info("数据库初始化成功!!!")

}

// 初始化数据库表

func InitMysqlTables() {
	err = GORM.AutoMigrate(
		models.OperationLog{},
		models.User{},
		models.Menu{},
		models.Role{},
		models.Dept{},
		models.APIPath{},
	)
	if err != nil {
		TPLogger.Error("生成数据表失败", err)
		os.Exit(0)
	}
	TPLogger.Info("生成数据表成功!!!")
}
