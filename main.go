package main

import (
	"fmt"

	"github.com/zpbird/zp-go-mod/zdirfiles"
)

// main aaa...
func main() {
	// y, _ := strconv.Atoi(zinput.Input("年份", zinput.RegYear))
	// m, _ := strconv.Atoi(zinput.Input("月份", zinput.RegMon))
	// monDays := ztimes.GetMonDays(y, m)
	// fmt.Println(monDays)

	b, e := zdirfiles.CopyFile("e:/新建文本文档.txt", "e:/zzz/abc.txt", true)
	fmt.Println(b, e)
	err := zdirfiles.MakeDir("e:/zzz")
	fmt.Println(err)
}
