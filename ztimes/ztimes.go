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

// GetMonDays ...
func GetMonDays(year, mon int) (days int) {

	return
}
