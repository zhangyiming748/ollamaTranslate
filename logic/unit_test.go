package logic

import (
	"os"
	"testing"
)

func init() {
	os.Setenv("OLLAMA_HOST", "http://127.0.0.1:11434/api/chat")
}

func TestGetChinese(t *testing.T) {
	ret := GetChinese("hello dear")
	t.Log(ret)
}
func TestGetHealth(t *testing.T) {
	os.Setenv("OLLAMA_HOST", "http://127.0.0.1:11434")
	ret, err := GetHealth()
	if err != nil {
		t.Error(err)
	}
	t.Log(string(ret))
}
