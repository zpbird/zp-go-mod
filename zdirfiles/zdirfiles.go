// Package zdirfiles import
package zdirfiles

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"github.com/Chain-Zhang/pinyin"
	"github.com/zpbird/zp-go-mod/zstrings"
)

var KindList = [...]string{"all", "dir", "file"}

// 当前系统路径分隔符
// var sysSep = string(os.PathSeparator)

// DirFileExist 参数[path:文件或目录路径字符串默认当前目录 kind:(all、dir、file)指定查找的种类]
func DirFileExist(path string, kind string) (bool, error) {
	info, err := os.Stat(path)
	switch kind {
	case KindList[0]:
		if err == nil {
			return true, nil
		}
		if os.IsNotExist(err) {
			return false, nil
		}
	case KindList[1]:
		if err == nil && info.IsDir() {
			return true, nil
		}
		if os.IsNotExist(err) {
			return false, nil
		}
	case KindList[2]:
		if err == nil && !info.IsDir() {
			return true, nil
		}
		if os.IsNotExist(err) {
			return false, nil
		}
	default:
		fmt.Println("DirFileExist()参数错误，kind应该为[dir、file、all]之一")
		return false, nil
	}

	return false, err
}

// GetDirFileList 参数[path:文件或目录路径字符串默认当前目录 kind:(all、dir、file)指定查找的种类]
func GetDirFileList(path string, kind string) ([]string, error) {
	dfSlice := []string{}
	subList, err := ioutil.ReadDir(path)
	if err != nil {
		return dfSlice, err
	}

	// 获取文件夹列表map
	dfMap := map[string]string{}

	switch kind {
	case KindList[0]:
		for _, df := range subList {
			if !zstrings.IncludeZhCn(df.Name()) {
				dfMap[strings.ToUpper(df.Name())] = df.Name()
				continue
			}
			strPinyin, err := pinyin.New(df.Name()).Split("").Mode(pinyin.WithoutTone).Convert()
			if err != nil {
				return dfSlice, err
			} else {
				dfMap[strPinyin] = df.Name()
			}
		}
	case KindList[1]:
		for _, df := range subList {
			if df.IsDir() {
				if !zstrings.IncludeZhCn(df.Name()) {
					dfMap[strings.ToUpper(df.Name())] = df.Name()
					continue
				}
				strPinyin, err := pinyin.New(df.Name()).Split("").Mode(pinyin.WithoutTone).Convert()
				if err != nil {
					return dfSlice, err
				} else {
					dfMap[strPinyin] = df.Name()
				}
			}
		}
	case KindList[2]:
		for _, df := range subList {
			if !df.IsDir() {
				if !zstrings.IncludeZhCn(df.Name()) {
					dfMap[strings.ToUpper(df.Name())] = df.Name()
					continue
				}
				strPinyin, err := pinyin.New(df.Name()).Split("").Mode(pinyin.WithoutTone).Convert()
				if err != nil {
					return dfSlice, err
				} else {
					dfMap[strPinyin] = df.Name()
				}
			}
		}
	default:
		fmt.Println("GetDirFileList()参数错误，kind应该为[dir、file、all]之一")
		return dfSlice, nil
	}

	// 生产按拼音排序后的slice
	tempSlice := []string{}
	for key := range dfMap {
		tempSlice = append(tempSlice, key)
	}

	sort.Strings(tempSlice)

	for _, key := range tempSlice {
		dfSlice = append(dfSlice, dfMap[key])
	}

	return dfSlice, err
}
