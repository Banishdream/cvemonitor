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
UserRepoMonitor
@describe	用户仓库监控,获取今天且不是fork的仓库写入数据库
@param	users []string "传入用户名"
*/
func UserRepoMonitor(users []string) {
	for _, user := range users {
		// TODO 1 构造url
		url := "https://api.github.com/users/" + user + "/repos"
		method := define.HttpMethodGET

		// TODO 2 解析返回的数据
		fmt.Printf("开始对 %s 进行抓取数据。。\n", url)
		var userRepoParam params.UserRepoParams
		if err := tool.ParseBody(url, method, &userRepoParam); err != nil {
			fmt.Println("解析body失败,", err)
			time.Sleep(define.RequestSleepTime)
			continue
		}

		// TODO 3 清晰数据, 将created_at不是今天的数据过滤, 插入今天的数据
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
				msg := fmt.Sprintf("有新的数据更新: %s\nURL:%s", userRepo.FullName, url)
				fmt.Println(msg)
				tool.SendMsg(msg)
			}
		}
		time.Sleep(define.RequestSleepTime)
	}
}

/*
CveMonitor
@describe	CVE页面监控,抓取数据写入数据库
*/
func CveMonitor() {
	year := time.Now().Format("2006")
	today := time.Now().Format("2006-01-02")

	// TODO 1 构造url
	url := "https://api.github.com/search/repositories?q=CVE-" + year + "&sort=updated"
	method := define.HttpMethodGET

	// TODO 2 解析返回的数据
	fmt.Printf("开始对 %s 进行抓取数据。。\n", url)
	var cveRepoParam params.CveRepoParams
	if err := tool.ParseBody(url, method, &cveRepoParam); err != nil {
		log.Error(err)
		return
	}

	// TODO 3 清洗数据, 插入今天的数据
	toolConf := tool.GetToolsConf()
	for _, item := range cveRepoParam.Items {

		//fmt.Printf("name: %s\n html_url: %s\n created_at: %s\n", item.Name, item.HtmlUrl, item.CreatedAt)
		// TODO 3.1 将created_at不是今天的数据过滤
		if item.CreatedAt[:10] == today || define.WriteAll {
			// TODO 3.2过滤用户黑名单
			urls := strings.Split(item.HtmlUrl, "/")
			user := urls[len(urls)-2]
			for _, blackUser := range toolConf.BlackUser {
				if user == blackUser {
					fmt.Printf("用户: %s 在黑名单, 已过滤\n", user)
					continue
				}
			}

			// TODO 3.3查询CVE官网是否存在,如果不存在跳过
			cveUrl := "https://cve.mitre.org/cgi-bin/cvename.cgi?name=" + item.Name
			if !tool.IsExistCVE(cveUrl) {
				log.Warnf("CVE官网中不存在 %s", item.Name)
				continue
			}

			// TODO 3.4 查询数据中是否存在,不存在写入数据
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
			msg := fmt.Sprintf("有新的数据更新: %s\nURL:%s", item.Name, item.HtmlUrl)
			fmt.Println(msg)
			tool.SendMsg(msg)
		}
		time.Sleep(define.RequestSleepTime)
	}
}

/*
KeywordMonitor
@describe	关键字的监控
*/
func KeywordMonitor() {
	toolConf := tool.GetToolsConf()
	today := time.Now().Format("2006-01-02")

	for _, keyword := range toolConf.KeywordList {

		// TODO 1 构造url
		url := "https://api.github.com/search/repositories?q=" + url2.QueryEscape(keyword) + "&sort=updated"
		method := define.HttpMethodGET

		// TODO 2 解析返回的数据
		var keywordParam params.KeywordParams
		if err := tool.ParseBody(url, method, &keywordParam); err != nil {
			log.Error(err)
			return
		}

		// TODO 3 清晰数据, 插入今天的数据
		for _, item := range keywordParam.Items {
			// 3.1 将created_at不是今天的数据过滤
			if item.CreatedAt[:10] == today || define.WriteAll {

				// TODO 3.2 过滤用户黑名单
				urls := strings.Split(item.HtmlUrl, "/")
				user := urls[len(urls)-2]
				for _, blackUser := range toolConf.BlackUser {
					if user == blackUser {
						fmt.Printf("用户: %s 在黑名单, 已过滤\n", user)
						continue
					}
				}

				// TODO 3.3 查询数据中是否存在,不存在则写入数据库
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
				msg := fmt.Sprintf("有新的数据更新: %s\nURL:%s", item.Name, item.HtmlUrl)
				fmt.Println(msg)
				tool.SendMsg(msg)
			}
		}
		time.Sleep(define.RequestSleepTime)
	}
}

/*
ToolMonitor
@describe	工具监控
*/
func ToolMonitor() {
	toolConf := tool.GetToolsConf()
	// TODO 1 拿到工具地址url
	for _, toolUrl := range toolConf.ToolsList {
		// TODO 2 解析返回的数据
		url := toolUrl
		method := define.HttpMethodGET

		var toolParam params.ToolParams
		if err := tool.ParseBody(url, method, &toolParam); err != nil {
			log.Error(err)
			return
		}

		// TODO 3 获取最新的 tagName, 如果没有则赋值 "no tag"
		tagUrl := url + "/releases"
		var toolTagParam params.ToolTagParams
		if err := tool.ParseBody(tagUrl, method, &toolTagParam); err != nil {
			log.Error(err)
			return
		}
		tagName := "no tag"

		//fmt.Printf("toolTagParam: %v, type: %T\n", toolTagParam, toolTagParam)
		// TODO 4 判断结构体不为空
		if len(toolTagParam) != 0 {
			tagName = toolTagParam[0].TagName
		}
		// TODO 4.1 构造 RedTeamToolsMonitor
		tp := models.RedTeamToolsMonitor{
			ToolName: toolParam.Name,
			PushedAt: toolParam.PushedAt,
			TagName:  tagName,
		}

		// TODO 5 插入数据
		// 如果是第一次, 则插入数据, 否则比对 pushed_at 时间, 有更新则更新数据
		id, pushedAt := dao.CVEDao.SelectData(tp.ToolName)
		var n int64
		if id == 0 { // 数据库中查询不到数据直接插入
			n = dao.CVEDao.InsertData(tp)
		} else if pushedAt != tp.PushedAt { // 查询到数据则更新
			n = dao.CVEDao.UpdateData(tp, id)
		}
		if n == 0 {
			fmt.Printf("未更新 %s 数据\n", tp.TagName)
			continue
		} else {
			msg := fmt.Sprintf("有新的数据更新: %s\nURL:%s", tp.TagName, toolUrl)
			fmt.Println(msg)
			tool.SendMsg(msg)
		}
		time.Sleep(define.RequestSleepTime)
	}
}
