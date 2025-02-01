package logic

import (
	"github.com/zhangyiming748/ollamaTranslate/util"
	"os"
)

func GetHealth() ([]byte, error) {
	ollamaHost := "http://ollama:11434"
	if current := os.Getenv("OLLAMA_HOST"); current != "" {
		ollamaHost = current
	}
	return util.HttpGet(nil, nil, ollamaHost, false)
}
