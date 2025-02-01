package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

func HttpPostJson(addHeaders map[string]string, data interface{}, urlPath string) (body []byte, err error) {
	bytesData, err := json.Marshal(data)
	if err != nil {
		return
	}
	reader := bytes.NewReader(bytesData)
	req, err := http.NewRequest("POST", urlPath, reader)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	for headerKey, headerVal := range addHeaders {
		req.Header.Set(headerKey, headerVal)
	}
	client := http.Client{
		Timeout: 100 * time.Second, // 设置超时时间为 10 秒
	}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()
	body, err = io.ReadAll(resp.Body)
	return
}
func HttpGet(addHeaders map[string]string, data map[string]string, urlPath string, proxy bool) (body []byte, err error) {
	//fmt.Printf("发送请求%v\n")
	params := url.Values{}
	urlInfo, err := url.Parse(urlPath)
	if err != nil {
		panic(err.Error())

	}
	for dataKey, dataVal := range data {
		params.Set(dataKey, dataVal)
	}
	urlInfo.RawQuery = params.Encode()
	fullUrl := urlInfo.String()
	log.Printf("发送的完整请求%v\n", fullUrl)
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return
	}
	for headerKey, headerVal := range addHeaders {
		req.Header.Set(headerKey, headerVal)
	}

	client := http.Client{}
	if proxy {
		proxyURL, err := url.Parse("http://127.0.0.1:8889")
		if err != nil {
			panic(err.Error())
		}
		// 设置代理的传输对象
		transport := &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}
		// 创建带有代理设置的客户端
		client = http.Client{
			Transport: transport,
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()

	body, err = io.ReadAll(resp.Body)
	statusCode := resp.StatusCode
	if statusCode == 403 {
		err = fmt.Errorf("403")
	}
	fmt.Printf("最终返回%v,%v\n", string(body), err)
	return
}
