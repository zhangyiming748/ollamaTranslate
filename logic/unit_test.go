package logic

import (
	"os"
	"testing"
)

func init() {
	os.Setenv("OLLAMA_HOST", "http://127.0.0.1:11434/api/chat")
}

func TestGetChinese(t *testing.T) {
	ret, _ := GetChinese("Imagine that you're teasing those lips.")
	ans := ret.Message.Content
	prefix, suffix := splitByLastNewline(ans)
	t.Logf("prefix = %v\tsuffix = %v\n", prefix, suffix)
}
func TestGetHealth(t *testing.T) {
	os.Setenv("OLLAMA_HOST", "http://127.0.0.1:11434")
	ret, err := GetHealth()
	if err != nil {
		t.Error(err)
	}
	t.Log(string(ret))
}
