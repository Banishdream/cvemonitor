package services

import (
	"cve-monitor/dao"
	"cve-monitor/define"
	"cve-monitor/models"
	"cve-monitor/params"
	"cve-monitor/tool"
	"fmt"
	"github.com/prometheus/common/log"
	url2 "net/url"
	"strings"
	"time"
)

/*
UserRepoMonitor 用户仓库监控
users: 传入用户名
获取今天且不是fork的仓库写入数据库
*/
func UserRepoMonitor(users []string) {
	for _, user := range users {
		// 1. 构造url
		url := "https://api.github.com/users/" + user + "/repos"
		method := define.HttpMethodGET

		// 2. 解析返回的数据
		fmt.Printf("开始对 %s 进行抓取数据。。\n", url)
		var userRepoParam params.UserRepoParams
		if err := tool.ParseBody(url, method, &userRepoParam); err != nil {
			fmt.Println("解析body失败,", err)
			time.Sleep(define.RequestSleepTime)
			continue
		}

		// 3. 清晰数据, 将created_at不是今天的数据过滤, 插入今天的数据
		today := time.Now().Format("2006-01-02")
		for _, userRepo := range userRepoParam {
			// 将日期是今天且不是fork的仓库过滤出来
			if (userRepo.CreatedAt[:10] == today && userRepo.Fork == false) || define.WriteAll {
				um := models.UserMonitor{
					RepoName: userRepo.FullName,
				}
				// 写入数据库
				n := dao.CVEDao.InsertData(um)
				if n == 0 {
					fmt.Printf("插入 %s 数据失败: \n", userRepo.FullName)
					continue
				}
				fmt.Printf("插入 %s 数据成功: \n", userRepo.FullName)
			}
		}
		time.Sleep(define.RequestSleepTime)
	}
}

/*
CveMonitor CVE页面监控
*/
func CveMonitor() {
	year := time.Now().Format("2006")
	today := time.Now().Format("2006-01-02")

	// 1. 构造url
	url := "https://api.github.com/search/repositories?q=CVE-" + year + "&sort=updated"
	method := define.HttpMethodGET

	// 2. 解析返回的数据
	fmt.Printf("开始对 %s 进行抓取数据。。\n", url)
	var cveRepoParam params.CveRepoParams
	if err := tool.ParseBody(url, method, &cveRepoParam); err != nil {
		log.Error(err)
		return
	}

	// 3. 清晰数据, 插入今天的数据
	toolConf := tool.GetToolsConf()
	for _, item := range cveRepoParam.Items {
		//fmt.Printf("name: %s\n html_url: %s\n created_at: %s\n", item.Name, item.HtmlUrl, item.CreatedAt)
		// 3.1 将created_at不是今天的数据过滤
		if item.CreatedAt[:10] == today || define.WriteAll {

			// 3.2 过滤用户黑名单
			urls := strings.Split(item.HtmlUrl, "/")
			user := urls[len(urls)-2]
			for _, blackUser := range toolConf.BlackUser {
				if user == blackUser {
					fmt.Printf("用户: %s 在黑名单, 已过滤\n", user)
					continue
				}
			}

			// 3.3 查询CVE官网是否存在
			cveUrl := "https://cve.mitre.org/cgi-bin/cvename.cgi?name=" + item.Name
			if !tool.IsExistCVE(cveUrl) {
				log.Warnf("CVE官网中不存在 %s", item.Name)
				continue
			}

			// 3.4 查询数据中是否存在
			// 由于设置了 cve_name 唯一主键, 重复数据插入会直接 warning级别告警. 无需业务实现, 略
			cm := models.CveMonitor{
				CveName:  item.Name,
				PushedAt: item.CreatedAt,
				Url:      item.HtmlUrl,
			}

			n := dao.CVEDao.InsertData(cm)
			if n == 0 {
				fmt.Printf("插入 %s 数据失败: \n", item.Name)
				continue
			}
			fmt.Printf("插入 %s 数据, url: %s 成功: \n", item.Name, item.HtmlUrl)

		}
		time.Sleep(define.RequestSleepTime)
	}
}

/*
KeywordMonitor 关键字的监控
*/
func KeywordMonitor() {
	toolConf := tool.GetToolsConf()
	today := time.Now().Format("2006-01-02")

	for _, keyword := range toolConf.KeywordList {

		// 1. 构造url
		url := "https://api.github.com/search/repositories?q=" + url2.QueryEscape(keyword) + "&sort=updated"
		method := define.HttpMethodGET

		// 2. 解析返回的数据
		var keywordParam params.KeywordParams
		if err := tool.ParseBody(url, method, &keywordParam); err != nil {
			log.Error(err)
			return
		}

		// 3. 清晰数据, 插入今天的数据
		for _, item := range keywordParam.Items {
			// 3.1 将created_at不是今天的数据过滤
			if item.CreatedAt[:10] == today || define.WriteAll {

				// 3.2 过滤用户黑名单
				urls := strings.Split(item.HtmlUrl, "/")
				user := urls[len(urls)-2]
				for _, blackUser := range toolConf.BlackUser {
					if user == blackUser {
						fmt.Printf("用户: %s 在黑名单, 已过滤\n", user)
						continue
					}
				}

				// 3.3 查询数据中是否存在
				// 由于设置了 cve_name 唯一主键, 重复数据插入会直接 warning级别告警. 无需业务实现, 略
				cm := models.KeywordMonitor{
					KeywordName: item.Name,
					PushedAt:    item.CreatedAt,
					Url:         item.HtmlUrl,
				}

				n := dao.CVEDao.InsertData(cm)
				if n == 0 {
					fmt.Printf("插入 %s 数据失败: \n", item.Name)
					continue
				}
				fmt.Printf("插入 %s 数据, url: %s 成功: \n", item.Name, item.HtmlUrl)

			}
		}
		time.Sleep(define.RequestSleepTime)
	}
}

func ToolMonitor() {
	toolConf := tool.GetToolsConf()
	// 1.拿到工具地址url
	for _, toolUrl := range toolConf.ToolsList {
		// 2. 解析返回的数据
		url := toolUrl
		method := define.HttpMethodGET

		var toolParam params.ToolParams
		if err := tool.ParseBody(url, method, &toolParam); err != nil {
			log.Error(err)
			return
		}

		// 3.获取最新的 tagName
		tagUrl := url + "/releases"
		var toolTagParam params.ToolTagParams
		if err := tool.ParseBody(tagUrl, method, &toolTagParam); err != nil {
			log.Error(err)
			return
		}

		// 4. 清洗数据, 构建数据
		// 4.1 如果有tagName, 则取最新的 否则 no tag
		tagName := "no tag"

		//fmt.Printf("toolTagParam: %v, type: %T\n", toolTagParam, toolTagParam)
		// 4.2 判断是否是空的 结构体
		if len(toolTagParam) != 0 {
			tagName = toolTagParam[0].TagName
		}

		tp := models.RedTeamToolsMonitor{
			ToolName: toolParam.Name,
			PushedAt: toolParam.PushedAt,
			TagName:  tagName,
		}

		// 5. 插入数据
		// 如果是第一次, 则插入数据, 否则比对 pushed_at 时间, 有更新则更新数据
		id, pushedAt := dao.CVEDao.SelectData(tp.ToolName)
		var n int64
		if id == 0 {
			n = dao.CVEDao.InsertData(tp)
		} else if pushedAt != tp.PushedAt {
			n = dao.CVEDao.UpdateData(tp, id)
		}
		if n == 0 {
			fmt.Printf("未更新 %s 数据\n", tp.TagName)
			continue
		} else {
			fmt.Printf("已更新 %s 数据, url: %s:\n", tp.TagName, toolUrl)
		}
		time.Sleep(define.RequestSleepTime)
	}
}
