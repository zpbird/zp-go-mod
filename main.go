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

	b, e := zdirfiles.GetDirFileList("D:\\project\\monitor\\template", "file")
	fmt.Println(b, e)

}
