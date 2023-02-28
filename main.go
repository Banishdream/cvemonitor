package main

import (
	"cve-monitor/pkg"
	"fmt"
)

func main() {
	fmt.Println("CVE、github指定工具和仓库，监控中...")
	pkg.Create_database()
	pkg.SettingS()
	return
}
