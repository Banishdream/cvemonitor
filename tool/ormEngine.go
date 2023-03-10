package tool

import (
	"cve-monitor/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type Orm struct {
	*xorm.Engine
}

var DbEngine = &Orm{}

/*
OrmEngine
@describe 数据库初始化操作
*/
func OrmEngine(cfg *AppConfig) error {
	db := cfg.Database
	// caocong:cc1764..@tcp(localhost:13306)/cvemonitor?charset=utf8mb4
	conn := db.Username + ":" + db.Password + "@tcp(" + db.Host + ":" + db.Port + ")/" + db.DbName + "?charset=" + db.Charset
	engine, err := xorm.NewEngine(db.Driver, conn)
	if err != nil {
		return err
	}

	engine.ShowSQL(db.ShowSql)

	err = engine.Sync2(
		new(models.CveMonitor),
		new(models.KeywordMonitor),
		new(models.UserMonitor),
		new(models.RedTeamToolsMonitor),
	)
	if err != nil {
		return err
	}

	DbEngine.Engine = engine
	return nil
}
