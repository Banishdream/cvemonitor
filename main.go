package main

import (
	"cve-monitor/define"
	"cve-monitor/services"
	"cve-monitor/tool"
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {

	fmt.Println("解析配置参数...")
	appConf := tool.ParseAppConfig(define.FilePath + "/" + define.AppFileName)

	fmt.Println("获取本地红队工具链文件")
	toolConf := tool.ParseToolConf(define.FilePath + "/" + define.ToolFileName)
	fmt.Println(toolConf)

	fmt.Println("连接数据库...")
	if err := tool.OrmEngine(appConf); err != nil {
		log.Fatal("连接数据库失败: ", err)
	}

	fmt.Println("CVE、github指定工具和仓库，监控中...")

	/*
		监控有三种方法
		1.用户仓库的监控
		2.CVE页面的监控
		3.关键字监控
		4.红队工具监控
	*/

	var wg sync.WaitGroup

	// 1.用户仓库的监控
	wg.Add(1)
	go func() {
		for {
			services.UserRepoMonitor(toolConf.UserList)
			time.Sleep(define.FuncExecInterval)
		}
		wg.Done()
	}()

	// 2.CVE页面的监控
	wg.Add(2)
	go func() {
		for {
			services.CveMonitor()
			time.Sleep(define.FuncExecInterval)
		}
		wg.Done()
	}()

	// 3.关键字监控
	wg.Add(3)
	go func() {
		for {
			services.KeywordMonitor()
			time.Sleep(define.FuncExecInterval)
		}
		wg.Done()
	}()

	// 4.红队工具监控
	wg.Add(4)
	go func() {
		for {
			services.ToolMonitor()
			time.Sleep(define.FuncExecInterval)
		}
		wg.Done()
	}()

	wg.Wait()

	return
}
