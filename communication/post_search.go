// @title	PostSearch
// @description	后端搜索接口
// @auth	ryl		2022/4/27	16:30
// @param	context	*gin.Context

package communication

import (
	"dianasdog/search"

	"github.com/gin-gonic/gin"
)

type SearchBody struct {
	Query string `json:"query" binding:"required"`
}

// @Summary 搜索系统接口
// @Tags Search
// @Description 后端搜索系统接口
// @Accept json
// @Produce json
// @Param query query string true "要搜索的句子"
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Bad Request"
// @Router /search [post]
func PostSearch(context *gin.Context) {
	var body SearchBody

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
	query := body.Query

	// 开始搜索流程
	result := search.Search(query)

	// 返回结果
	context.JSON(200, gin.H{
		"content": result, //result,
	})
}
