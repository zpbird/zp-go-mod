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
	RegAll = ""
	// numbers
	RegNum = ""

	// time
	RegYear = "^[0-9]{1}[0-9]{3}$"
	RegMon  = "^([1-9]{1}|0{1}[0-9]{1}|1{1}[0-2]{1})$"
	RegDay  = ""
	RegHms  = ""

	// phoneNumber
	RegPhone = ""
	// email
	RegEmail = ""
)

// --------------------------------------------------------------------------------
// 工具函数------------------------------------------------------------------------
// --------------------------------------------------------------------------------

// GetNum ...
func GetNum() {

}

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

// Input 基本输入
func Input(termName string, regexpText string) (inputText string) {
	Clear()
	var trigger = false
	for !trigger {
		fmt.Printf("请输入 %s : ", termName)
		reader := bufio.NewReader(os.Stdin)
		inputBytes, _, err := reader.ReadLine()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Occur error when input", termName, ":", err)
			continue
		}
		inputText = string(inputBytes)
		trigger, err = regexp.MatchString(regexpText, inputText)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Occur error when match", termName, "(", inputText, "):", err)
			continue
		}
		if !trigger {
			Clear()
			fmt.Fprintln(os.Stdout, termName, "(", inputText, ") -> 输入数据格式或数据范围不正确!")
		}
	}
	return
}
