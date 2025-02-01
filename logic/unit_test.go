package logic

import "testing"

func TestGetChinese(t *testing.T) {
	ret := GetChinese("hello dear")
	t.Log(ret)
}
