package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello 打招呼的函数
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("请输入一个名字吧，不然怎么打招呼呢")
	}

	message := fmt.Sprintf(randomFormat(), name)

	return message, nil
}

// Hellos 同时打招呼多人模式
func Hellos(names []string) (map[string]string, error) {

	// messages 创建了一个映射建与值字符串类型
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}

	return messages, nil
}

// randomFormat 返回一句随机的问h语
func randomFormat() string {
	formats := []string{
		"Hi, %v. 您好",
		"我知道你来了 %v！",
		"哇，%v帅哥?",
	}

	return formats[rand.Intn(len(formats))]
}
