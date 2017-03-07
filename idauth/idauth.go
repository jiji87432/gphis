// gphis/idauth - Identity Authentication 身份认证模块
//
package idauth

import (
	"fmt"
)

// 全局登录状态信息存储数据结构
var tokenList []string

func init() {

}

type IAuth interface {
	Authenticate() (bool, error)
}

// Auth: 身份验证函数
func Auth(ia IAuth) (bool, error) {
	return ia.Authenticate()
}

// Login: 登录函数
func Login(ia IAuth) (string, error) {

}

// Logout: 退出函数
func Logout(ia IAuth) (bool, error) {

}

// Check: 校验token是否合法
func Check(token string) bool {
	if len(tokenList) != 0 {
		for tmp := range tokenList {
			if tmp == token {
				return true
			}
		}
	}
	return false
}
