package tool

import (
	"cve-monitor/define"
	"cve-monitor/models"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"time"
)

type Orm struct {
	*xorm.Engine
	drive   string
	conn    string
	showSql bool
}

var DbEngine = &Orm{}

/*
OrmEngine
@describe 数据库初始化操作
*/
func OrmEngine(cfg *AppConfig) error {

	if cfg.Sqlite3DB.Enable == true {
		DbEngine.drive = cfg.Sqlite3DB.Drive
	} else if cfg.MysqlDB.Enable == true {
		DbEngine.drive = cfg.MysqlDB.Driver
	} else {
		return errors.New("没有可选择的数据库")
	}

	switch DbEngine.drive {
	case "mysql":
		db := cfg.MysqlDB
		// caocong:cc1764..@tcp(localhost:13306)/cvemonitor?charset=utf8mb4
		DbEngine.conn = db.Username + ":" + db.Password + "@tcp(" + db.Host + ":" + db.Port + ")/" + db.DbName + "?charset=" + db.Charset

	case "sqlite", "sqlite3":
		db := cfg.Sqlite3DB
		_, err := os.Stat(db.DBFilePath)
		if err != nil {
			if os.IsNotExist(err) {
				_, err = os.Create(db.DBFilePath)
				if err != nil {
					return err
				}
			} else {
				return err
			}
		}
		DbEngine.conn = db.DBFilePath

	default:
		return errors.New("没有可选择的 sqlite 或 mysql 数据库")
	}

	engine, err := xorm.NewEngine(DbEngine.drive, DbEngine.conn)
	if err != nil {
		return err
	}

	engine.ShowSQL(DbEngine.showSql)

	err = engine.Sync2(
		new(models.CveMonitor),
		new(models.KeywordMonitor),
		new(models.UserMonitor),
		new(models.RedTeamToolsMonitor),
	)

	if err != nil {
		return err
	}
	engine.SetConnMaxLifetime(define.MysqlWaitTimeout * time.Second)
	DbEngine.Engine = engine

	return nil
}
