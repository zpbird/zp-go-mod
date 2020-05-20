package main

import (
	"fmt"
	"strconv"

	"github.com/zpbird/zp-go-mod/zinput"
	"github.com/zpbird/zp-go-mod/ztimes"
)

// main aaa...
func main() {

	zinput.Clear()
	y, _ := strconv.Atoi(zinput.Input("年份", zinput.RegYear))
	m, _ := strconv.Atoi(zinput.Input("月份", zinput.RegMon))
	monDays := ztimes.GetMonDays(y, m)
	fmt.Println(monDays)
}
