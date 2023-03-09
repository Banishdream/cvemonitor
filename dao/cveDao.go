package dao

import (
	"cve-monitor/define"
	"cve-monitor/models"
	"cve-monitor/tool"
	"fmt"
	"github.com/prometheus/common/log"
)

type cveDao struct {
	*tool.Orm
}

var CVEDao = &cveDao{tool.DbEngine}

func (cveD *cveDao) InsertData(data interface{}) int64 {
	switch v := data.(type) {
	case models.CveMonitor:
		n, err := cveD.Insert(&v)
		if err != nil && define.Debug {
			log.Warnln(err)
		}
		return n
	case models.UserMonitor:
		n, err := cveD.Insert(&v)
		if err != nil && define.Debug {
			log.Warnln(err)
		}
		return n
	case models.KeywordMonitor:
		n, err := cveD.Insert(&v)
		if err != nil && define.Debug {
			log.Warnln(err)
		}
		return n
	case models.RedTeamToolsMonitor:
		n, err := cveD.Insert(&v)
		if err != nil && define.Debug {
			log.Warnln(err)
		}
		return n
	default:
		fmt.Println("undefined type！")
	}
	return 0
}

func (cveD *cveDao) UpdateData(data models.RedTeamToolsMonitor, id int64) int64 {
	n, err := cveD.ID(id).Update(&data)
	if err != nil && define.Debug {
		log.Warnf("update mysql err: %s", err)
		return 0
	}
	return n
}

func (cveD *cveDao) SelectData(toolName string) (int64, string) {
	var tp models.RedTeamToolsMonitor
	if _, err := cveD.Where("tool_name = ?", toolName).Get(&tp); err != nil {
		if define.Debug {
			log.Warnln(err.Error())
		}
		return 0, ""
	}
	return tp.Id, tp.PushedAt
}
