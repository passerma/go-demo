package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// {功能} 生成随机验证码
// {参数} 无
// {返回} 验证码
func GenRandCode() (code string) {
	code = fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	return
}
