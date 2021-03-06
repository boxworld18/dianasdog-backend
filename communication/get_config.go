// @title	GetConfig
// @description	后端发出写入行为之接口
// @auth	ryl		2022/4/26		17:30
// @param	context	*gin.Context

package communication

import (
	"dianasdog/database"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

// @Summary 取得写入行为描述
// @Tags Setting
// @Description 后端返回写入行为描述之接口
// @Produce json
// @Param resource query string true "特型卡名称 (如: car, poem 等)"
// @Success 200 {object} string "{"data": data}"
// @Failure 400 {object} string "Bad Request"
// @Router /setting [get]
func GetConfig(context *gin.Context) {

	// 检查收到信息的格式是否正确
	resource, ok := context.GetQuery("resource")

	// 若不是，则返回错误
	if !ok {
		context.JSON(400, gin.H{
			"err": "wrong param",
		})
		return
	}

	// 取得文件
	data, err := database.GetFile(database.ConfigClient, "file", resource)

	// 若不存在文件/对应特型卡类型，则返回错误
	if err != nil {
		context.JSON(400, gin.H{
			"err": err.Error(),
		})
		return
	}

	// 结果转化为 json
	var result map[string]interface{}
	json.Unmarshal(data, &result)

	// 正常返回结果
	context.JSON(200, gin.H{
		"data": result,
	})
}
