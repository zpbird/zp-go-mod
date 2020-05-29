// Package zinput import ...
package zinput

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
)

const (
	RegAll     = ""
	RegAllRule = "任意字符："

	// numbers
	RegNum     = "^[0-9]*$"
	RegNumRule = "[0-9]组成的纯数字："

	// time
	RegYear       = "^[0-9]{1}[0-9]{3}$"
	RegYearRul    = "[0000-9999]之间年份格式的数字："
	RegMon        = "^([1-9]{1}|1{1}[0-2]{1})$"
	RegMonRule    = "[1-12]之间月份格式的数字："
	RegDay        = "^[1-9]{1}$|^[1-2]{1}[0-9]{1}$|^3[01]{1}$"
	RegDayRule    = "[1-31]之间的数字："
	RegHours      = "^[0-9]{1}$|^1[0-9]{1}$|^2[0-3]{1}$"
	RegHoursRule  = "[0-23]之间的数字："
	RegMinSec     = "^[0-9]{1}$|^[1-5]{1}[0-9]{1}$"
	RegMinSecRule = "[0-59]之间的数字："

	// phoneNumber
	RegPhone = ""
	// email
	RegEmail = ""
	// y or n (yes or no)
	RegYn     = "^[yn]$"
	RegYnRule = "[y或n]字符："
)

// --------------------------------------------------------------------------------
// 工具函数------------------------------------------------------------------------
// --------------------------------------------------------------------------------

// --------------------------------------------------------------------------------
// 基本函数------------------------------------------------------------------------
// --------------------------------------------------------------------------------

// Clear 清屏...
func Clear() {
	// fmt.Printf("\x1bc") //清屏
	cmd := exec.Command("cmd.exe", "/c", "cls") //windows清屏命令
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// StopContinue ...
func StopContinue() {
	fmt.Printf("\n输入数据格式或数据范围不正确!\n回车键重新选择")
	var t string //Scanln的临时变量
	fmt.Scanln(&t)

}

// Input 基本输入
func Input(termString string, regexpText string) (inputText string) {
	var trigger = false
	for !trigger {
		fmt.Printf("请输入%s", termString)
		reader := bufio.NewReader(os.Stdin)
		inputBytes, _, err := reader.ReadLine()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Occur error when input", termString, ":", err)
			continue //??? os.Exit(0)
		}

		inputText = string(inputBytes)
		trigger, err = regexp.MatchString(regexpText, inputText)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Occur error when match", termString, "(", inputText, "):", err)
			continue //??? os.Exit(0)
		}

		if !trigger {
			StopContinue()
			Clear()
		}
	}
	return
}
