// @title	GoSetting
// @description	后端接收写入行为之接口
// @auth	ryl		2022/4/13		17:30
// @param	context	*gin.Context

package communication

import (
	"dianasdog/io"
	"encoding/json"
	"io/ioutil"
	"mime/multipart"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SettingBody struct {
	Resource string                `form:"resource" binding:"required"`
	Data     string                `form:"data"`
	File     *multipart.FileHeader `form:"file"`
}

type SettingJson struct {
	Resource string                 `json:"resource" binding:"required"`
	Setting  map[string]interface{} `json:"write_setting" binding:"required"`
}

func PostSetting(context *gin.Context) {
	var body SettingBody
	var err error

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
		str, _ := strconv.Unquote(content)

		// 检查数据内容是否正确
		var jsonContent SettingJson
		err = json.Unmarshal([]byte(str), &jsonContent)

		// 若不正确，则返回错误
		if err != nil {
			context.JSON(400, gin.H{
				"err": err.Error(),
			})
			return
		}

		data, _ = json.Marshal(jsonContent)

	} else if body.File != nil {
		// 若使用文件传输
		fileContent, _ := body.File.Open()
		data, err = ioutil.ReadAll(fileContent)

	} else {
		// 若没有传输数据，则错误
		context.JSON(401, gin.H{
			"err": "wrong parameters!",
		})
		return
	}

	// 若过程中出现错误
	if err != nil {
		context.JSON(402, gin.H{
			"err": err.Error(),
		})
		return
	}

	// 否则调用函数写入文件
	io.SetConfig(res, data)

	// 返回对应值
	context.JSON(200, gin.H{
		"message": "successful!", //result,
	})
}
