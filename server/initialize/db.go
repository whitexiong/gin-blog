package initialize

import (
	"blog/common"
	"blog/dao"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

func InitDb() {
	switch common.CONFIG.System.DBType {
	case "mysql":
		InitMysql()
	default:
		InitMysql()
	}

}

func InitMysql() {
	sql := common.CONFIG.Mysql
	if db, err := gorm.Open("mysql", sql.Username+":"+sql.Password+"@("+sql.Path+")/"+sql.Dbname+"?"+sql.Config); err != nil {
		common.LOG.Error("DefaultDB 数据库启动异常", zap.Any("err", err))
		os.Exit(1)
	} else {
		common.DB = db
		db.DB().SetMaxIdleConns(sql.MaxIdleConns)
		db.DB().SetMaxOpenConns(sql.MaxOpenConns)
		db.LogMode(sql.LogMode)
		db.SingularTable(true)
		DBTableMigrate()
	}

}

func DBTableMigrate() {
	common.DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").AutoMigrate(
		dao.SysRole{},  // 角色表
		dao.SysMenu{},  // 菜单表
		dao.SysUser{},  // 用户
		dao.SysDept{},  // 部门
		dao.SysOpLog{}, // 日志
	)
	common.LOG.Info("register table success")
}
