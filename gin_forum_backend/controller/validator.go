package controller

import (
	"strings"
)

// removeTopStruct 去除提示信息中的结构体名称
func RemoveTopStruct(fields string)  string {
	index:=strings.Index(fields,"Error:")
	res:=fields[index+6:]
	return res
}
