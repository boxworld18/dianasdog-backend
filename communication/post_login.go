// @title	PostLogin
// @description	后端密码接口
// @auth	ryl		2022/4/21	10:30
// @param	context	*gin.Context

package communication

import (
	"github.com/gin-gonic/gin"
)

type LoginBody struct {
	Username string `json:"username" binding:"required"`
}

func PostLogin(context *gin.Context) {
	var body LoginBody

	// 检查收到信息的格式是否正确
	err := context.ShouldBindJSON(&body)

	// 若不是，则返回错误
	if err != nil {
		context.JSON(400, gin.H{
			"err": err.Error(),
		})
		return
	}

	// 取得 query 字段
	username := body.Username
	username = username

	// 开始搜索流程
	// result := search.IntentRecognition(query)

	// 返回结果
	context.JSON(200, gin.H{
		"password": "",
	})
}
