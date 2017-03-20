// gphis/idauth - Identity Authentication 身份认证模块
//
// 身份认证模块定义了接口IAuth, 所有实现了该接口的结构均可用于进行身份认证,
// Authenticate函数返回的两个参数是: bool类型的认证是否成功, string类型的用户ID.
//
// 定义了统一登录消息响应格式AuthResp, 其中, Success字段表示当前登录是否成功,
// Response字段表示登录的返回信息. 在Login中, Response正常情况下会返回登录token.
//
// 登录token是用于状态管理的一个字符串, 产生方式是md5(时间+id).
//
// Check函数用户检测token合法性, 成功则返回校验的bool值以及用户id.
package idauth

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"gphis/utils"
)

// 定义响应消息结构
type AuthResp struct {
	Success  bool   `json:"success"`
	Response string `json:"response"`
}

// 全局登录状态信息存储数据结构, key存储token, value存储用户信息(用户ID)
var tokenList map[string]string

func init() {
	tokenList = make(map[string]string)
	initDir()
	readFile()
}

type IAuth interface {
	Authenticate() (bool, string)
}

// 登录函数
func Login(ia IAuth) (string, error) {
	ok, id := ia.Authenticate()
	if ok == true {
		tmp := id2Token(id)
		if tmp != "" {
			delete(tokenList, tmp)
		}

		token := getToken(id)
		rstmp, err := json.Marshal(&AuthResp{Success: true, Response: token})
		if err != nil {
			return "", fmt.Errorf("Login error: %v", err)
		}

		tokenList[token] = id
		writeFile()
		return string(rstmp), nil
	} else {
		rstmp, err := json.Marshal(&AuthResp{Success: false, Response: "Login failed."})
		if err != nil {
			return "", fmt.Errorf("Login error: %v", err)
		}

		return string(rstmp), nil
	}
}

// 退出函数
func Logout(token string) (string, error) {
	_, ok := tokenList[token]
	if ok == false {
		rstmp, err := json.Marshal(&AuthResp{Success: false, Response: "不存在的token"})
		if err != nil {
			return "", fmt.Errorf("Logout error: %v", err)
		}

		return string(rstmp), nil
	} else {
		delete(tokenList, token)
		writeFile()
		rst, err := json.Marshal(&AuthResp{Success: true, Response: "退出成功"})
		if err != nil {
			return "", fmt.Errorf("Logout error: %v", err)
		}

		return string(rst), nil
	}
}

// 校验token是否合法
func Check(token string) (bool, string) {
	rst, ok := tokenList[token]

	if ok == true {
		return ok, rst
	} else {
		return ok, ""
	}
}

// 获取状态token
func getToken(id string) string {
	return utils.MD5([]byte(time.Now().Format("20060102150405") + id))
}

// 本地文件存储目录初始化
func initDir() {
	_, err := os.Stat("/gphis/idauth")
	if err != nil {
		// 目录不存在, 创建目录
		if os.IsNotExist(err) {
			er := os.MkdirAll("/gphis/idauth", 0755)
			if er != nil {
				log.Fatalf("初始化本地存储目录失败: %v", er)
			}
		} else {
			log.Fatalf("初始化本地存储目录失败: %v", err)
		}
	}
}

// 内存中的存储持久化到本地磁盘
func writeFile() {
	file, err := os.OpenFile("/gphis/idauth/auths", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	defer file.Close()

	if err != nil {
		log.Printf("(低错误级别)写入文件失败: %v", err)
	} else {
		for token, id := range tokenList {
			file.WriteString(token + ":" + id + "\n")
		}
	}
}

// 从本地存储读取程序宕掉之前的token
func readFile() {
	file, err := os.OpenFile("/gphis/idauth/auths", os.O_RDONLY, 0644)
	defer file.Close()

	if err != nil {
		log.Printf("(低错误级别)读取文件错误: %v", err)
	} else {
		reader := bufio.NewReader(file)
		for {
			line, _, err := reader.ReadLine()
			if err != nil {
				if err == io.EOF {
					break
				}
				fmt.Printf("(低错误级别)读取文件错误: %v", err)
				break
			} else {
				arr := strings.Split(string(line), ":")
				tokenList[arr[0]] = arr[1]
			}
		}
	}
}

// 由用户ID反向检索token
func id2Token(idIn string) string {
	for token, id := range tokenList {
		if id == idIn {
			return token
		}
	}
	return ""
}
