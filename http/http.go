package http

import (
	"io/ioutil"
	"net/http"
	"time"
)

func HttpGet(url string) (body []byte, err error) {
	// https://stackoverflow.com/questions/45751608/why-is-http-client-prefixed-with
	// http.Client{} 是一个复合字面量，它创建结构类型 http.Client 的值
	//前面加 & 获取存储此结构值的匿名变量的地址：
	//获取复合文字的地址会生成一个指针，该指针指向使用该文字的值初始化的唯一变量。
	// 如果 client 是一个指针，你可以自由地将它传递给其他函数，只会复制指针值，而不是复制http.Client结构， 因此结构本身(http.Client值)将会被重用。
	// 如果你不使用指针， 如果你将它传递给其他函数，结构体本身将被复制而不是重用
	client := &http.Client{Timeout: 1 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
