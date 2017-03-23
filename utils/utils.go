// gphis/utils - 通用工具模块
package utils

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"time"
)

// 产生16进制MD5值
func MD5(data []byte) string {
	return fmt.Sprintf("%x", md5.Sum(data))
}

// 产生标准ID序列
func GetSerial() string {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	tmp := random.Intn(99)
	for {
		if tmp >= 10 {
			break
		} else {
			tmp = random.Intn(99)
		}
	}
	now := time.Now().Format("060102150405")
	return fmt.Sprintf("%s%d", now, tmp)
}
