// gphis/permissions - 权限控制模块
//
// permissions包含有公开函数GetPermit, 传入2个参数, 第一个字符串参数为指定用户
// 权限串, 第二个整形参数为本模块权限代码, 返回响应为是否具有访问权限
package permissions

import (
	"fmt"
	"strconv"
	"strings"
)

// 基本响应结构
type BaseResp struct {
	Success bool   `json:"success"` // 是否具有模块访问权限
	Reason  string `json:"reason"`  // 原因
}

// 查询是否具有模块访问权限, ps是权限字符串, 格式为1:2:3:4, module是模块权限代码
func GetPermit(ps string, module int) *BaseResp {
	arr, err := permit2Array(ps)
	if err != nil {
		return &BaseResp{Success: false, Reason: err.Error()}
	}

	rst := findPermit(module, arr)
	if rst == true {
		return &BaseResp{Success: true, Reason: "具有此模块访问权限"}
	} else {
		return &BaseResp{Success: false, Reason: "不具有此模块访问权限, 详情联系系统管理员"}
	}
}

func findPermit(module int, arr []int) bool {
	for _, tmp := range arr {
		if module == tmp {
			return true
		}
	}
	return false
}

func permit2Array(ps string) ([]int, error) {
	var arrRet []int
	strArr := strings.Split(ps, ":")
	for _, str := range strArr {
		tmp, err := strconv.Atoi(str)
		if err != nil {
			return arrRet, fmt.Errorf("Sth. went wrong while converting permit str: %v", err)
		}
		arrRet = append(arrRet, tmp)
	}
	return arrRet, nil
}
