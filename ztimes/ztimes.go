// Package ztimes import "zp-go-mod/times"
package ztimes

import (
	"fmt"
	"time"
)

const (
	tFormat1 = "2006-01-02 15:04:05"
)

var (
	loc, _       = time.LoadLocation("Local")
	tLocation, _ = time.LoadLocation("Asia/Shanghai")
	tZone        = time.FixedZone("CST", 8*3600)
	now          = time.Now()
)

// GetNow 获取当前时间...
func GetNow() {
	fmt.Println(now.Format(tFormat1))
}

// GetMonDays 获取指定年月的天数...
func GetMonDays(year int, month int) (days int) {

	if month != 2 {
		if month == 4 || month == 6 || month == 9 || month == 11 {
			days = 30

		} else {
			days = 31
		}
	} else {
		if ((year%4) == 0 && (year%100) != 0) || (year%400) == 0 {
			days = 29
		} else {
			days = 28
		}
	}
	return
}
