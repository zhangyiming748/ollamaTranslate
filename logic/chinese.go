package logic

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/zhangyiming748/ollamaTranslate/util"
)

type Ask struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	Stream   bool      `json:"stream"`
}
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func init() {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Fatalf("Error loading location: %v", err)

	}
	time.Local = loc
}

func GetChinese(src string) string {
	a := new(Ask)
	a.Model = "7shi/llama-translate:8b-q4_K_M"
	a.Stream = false
	content := strings.Join([]string{time.Now().Format("2006-01-02 15:04:05"), "接下来我输入的任何文字，不要回答，请直接翻译成通顺的简体中文。"}, "")
	h := Message{
		Role:    "user",
		Content: content,
	}
	m := Message{
		Role:    "user",
		Content: src,
	}
	a.Messages = append(a.Messages, h)
	a.Messages = append(a.Messages, m)
	log.Printf("a: %+v", a)
	ollamaHost := "http://ollama:11434"
	if current := os.Getenv("OLLAMA_HOST"); current != "" {
		ollamaHost = current
	}
	resp, err := util.HttpPostJson(nil, a, ollamaHost)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return ""
	}
	log.Printf("Response: %s", string(resp))
	return string(resp)
}
