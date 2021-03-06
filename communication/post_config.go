// @title	PostConfig
// @description	后端接收写入行为之接口
// @auth	ryl		2022/4/13		17:30
// @param	context	*gin.Context

package communication

import (
	"dianasdog/setter"
	"encoding/json"
	"io/ioutil"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

type ConfigBody struct {
	Resource string                `form:"resource" binding:"required"`
	Data     string                `form:"data"`
	File     *multipart.FileHeader `form:"file"`
}

type ConfigJson struct {
	Resource string                 `json:"resource" binding:"required"`
	Setting  map[string]interface{} `json:"write_setting" binding:"required"`
}

// @Summary 发送写入行为描述
// @Tags Setting
// @Description 后端接收写入行为描述之接口
// @Accept mpfd
// @Produce json
// @Param resource formData string true "特型卡名称 (如: car, poem 等)"
// @Param file formData file false "文件形式上传之数据"
// @Param data formData string false "数据内容 (JSON 格式)"
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Bad Request"
// @Router /setting [post]
func PostConfig(context *gin.Context) {
	var body ConfigBody
	var err error
	var msg string

	// 检查收到信息的格式是否正确
	err = context.ShouldBind(&body)

	// 若不是，则返回错误
	if err != nil {
		context.JSON(400, gin.H{
			"err": err.Error(),
		})
		return
	}

	// 取得特型卡 ID
	res := body.Resource
	var data []byte

	// 取得对应的数据内容
	if len(body.Data) > 0 {
		// 若使用 json 格式的数据
		content := body.Data

		// 去掉前端多余的引号
		// content, _ := strconv.Unquote(content)

		// 检查数据内容是否正确
		var jsonContent ConfigJson
		err = json.Unmarshal([]byte(content), &jsonContent)

		// 若不正确，则返回错误

		if err != nil {
			msg = err.Error()
		}

		if jsonContent.Setting == nil {
			msg = "json data error: wrong parameters!" + content
		}

		data, _ = json.Marshal(jsonContent)

	} else if body.File != nil {
		// 若使用文件传输
		fileContent, _ := body.File.Open()
		data, err = ioutil.ReadAll(fileContent)

		if err != nil {
			msg = err.Error()
		}

	} else {
		// 若没有传输数据，则错误
		msg = "form data error: wrong parameters!"
	}

	// 若过程中出现错误
	if len(msg) > 0 {
		context.JSON(400, gin.H{
			"err": msg,
		})
		return
	}

	// 否则调用函数写入文件
	err = setter.SetConfig(res, data)

	if err != nil {
		context.JSON(400, gin.H{
			"err": err.Error(),
		})
		return
	}

	// 返回对应值
	context.JSON(200, gin.H{
		"message": "successful!", //result,
	})
}
