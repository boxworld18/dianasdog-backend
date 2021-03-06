// @title		SetupRouter
// @description	后端与前端交互之接口
// @auth		ryl		2022/4/20	23:30

package communication

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())

	// 搜索句子
	router.POST("/search", PostSearch)
	// 登入较验
	router.POST("/login", PostLogin)
	// 数据上传
	router.POST("/data", PostData)
	// 模板/垃圾词/意图词上传
	router.POST("/pattern", PostTemplate)
	// 写入配置上传
	router.POST("/setting", PostConfig)
	// es测试接口
	router.POST("/testes", PostEsTest)
	// 删除用户
	router.POST("/userdelete", DeleteUser)
	// 添加用户
	router.POST("/useradd", AddUser)

	// 数据下载
	router.GET("/data", GetData)
	// 返回某一特型卡下的所有文件名
	router.GET("/dataname", GetDataName)
	// 模板/垃圾词/意图词下载
	router.GET("/pattern", GetTemplate)
	// 写入配置下载
	router.GET("/setting", GetConfig)
	// 取得某一条数据
	router.GET("/item", GetItem)
	// 取得所有特型卡名字
	router.GET("/category", GetCategory)
	// 取得某个特型卡对应 item 中的所有键值
	router.GET("/key", GetKey)
	// 取得所有用户
	router.GET("/alluser", AllUser)

	return router
}
