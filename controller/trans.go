package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhangyiming748/ollamaTranslate/logic"
	"net/http"
)

type TransController struct{}

/*
curl --location --request GET 'http://127.0.0.1:8192/api/v1/s1/gethello?user=<user>' \
--header 'User-Agent: Apifox/1.0.0 (https://www.apifox.cn)'
*/
type HealthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"` // 可选的错误信息
}

func (t TransController) GetHealth(ctx *gin.Context) {
	health, err := logic.GetHealth()
	if err != nil {
		// 返回 500 错误，并包含错误信息
		response := HealthResponse{
			Status:  "error",
			Message: err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}
	var response HealthResponse
	if string(health) == "Ollama is running" {
		response = HealthResponse{
			Status: "ok",
		}
	} else {
		response = HealthResponse{
			Status:  "degraded",
			Message: "Service is not fully healthy.",
		}
	}
	ctx.JSON(http.StatusOK, response)
}

// 结构体必须大写 否则找不到
type RequestBody struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
type ResponseBody struct {
	Name string `json:"name"`
}

/*
 */
func (t TransController) PostChinese(ctx *gin.Context) {
	fmt.Println("get")
	var requestBody RequestBody
	if err := ctx.BindJSON(&requestBody); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	} else {
		fmt.Println(requestBody)
	}
	fmt.Println(requestBody.Name, requestBody.Age)
	var rep ResponseBody
	rep.Name = fmt.Sprintf("我已经%d年没见过%s了", requestBody.Age, requestBody.Name)
	ctx.JSON(200, rep)
}
