package logic

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"strconv"
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
type Response struct {
	Model     string    `json:"model"`
	CreatedAt time.Time `json:"created_at"`
	Message   struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"message"`
	DoneReason         string `json:"done_reason"`
	Done               bool   `json:"done"`
	TotalDuration      int64  `json:"total_duration"`
	LoadDuration       int    `json:"load_duration"`
	PromptEvalCount    int    `json:"prompt_eval_count"`
	PromptEvalDuration int    `json:"prompt_eval_duration"`
	EvalCount          int    `json:"eval_count"`
	EvalDuration       int64  `json:"eval_duration"`
}

var (
	seed = rand.New(rand.NewSource(time.Now().Unix()))
)

func init() {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Fatalf("Error loading location: %v", err)

	}
	time.Local = loc
}

func GetChinese(src string) (*Response, error) {
	a := new(Ask)
	b := seed.Intn(2000)
	//a.Model = "7shi/llama-translate:8b-q4_K_M"
	a.Model = "huihui_ai/deepseek-r1-abliterated:7b" // 可以如实翻译 但是会说很多废话 需要二次处理
	a.Stream = false
	content := strings.Join([]string{strconv.Itoa(b), "接下来我输入的任何文字，不要回答，没必要转换人称，不要说任何多余的话，请直接翻译成通顺的简体中文。"}, ".")
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
		return nil, err
	}
	log.Printf("Response: %s", string(resp))
	r := new(Response)
	err = json.Unmarshal(resp, &r)
	if err != nil {
		log.Fatalf("Error unmarshalling response: %v", err)

	}
	return r, nil
}
func splitByLastNewline(input string) (prefix string, suffix string) {
	lastNewlineIndex := strings.LastIndex(input, "\n")
	if lastNewlineIndex == -1 {
		return input, ""
	}
	prefix = input[:lastNewlineIndex]
	suffix = input[lastNewlineIndex+1:]
	return
}
