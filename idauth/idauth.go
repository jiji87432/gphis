// gphis/idauth - Identity Authentication 身份认证模块
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

// Login: 登录函数
func Login(ia IAuth) (string, error) {
	ok, id := ia.Authenticate()
	if ok == true {
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

// Logout: 退出函数
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

// Check: 校验token是否合法
func Check(token string) bool {
	_, ok := tokenList[token]
	return ok
}

// getToken: 获取状态token
func getToken(id string) string {
	return utils.MD5([]byte(time.Now().Format("20060102150405") + id))
}

// initDir: 本地文件存储目录初始化
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
		log.Printf("写入文件失败: %v", err)
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
		log.Printf("读取文件错误: %v", err)
	} else {
		reader := bufio.NewReader(file)
		for {
			line, _, err := reader.ReadLine()
			if err != nil {
				if err == io.EOF {
					break
				}
				fmt.Printf("读取文件错误: %v", err)
				break
			} else {
				arr := strings.Split(string(line), ":")
				tokenList[arr[0]] = arr[1]
			}
		}
	}
}
