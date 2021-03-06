// @title	SetTestData
// @description	将单元测试需要的数据入库
// @auth	ryl		2022/4/25	18:00
// @param	stage	int			数据模式
// @return	err		error		错误值

package testcase

import (
	"dianasdog/database"
	"dianasdog/path"
	"io/ioutil"
)

func SetTestData(stage int) error {

	// 得到此文件的绝对路径
	abspath, _ := path.GetAbsPath()
	abspath += "testcase/"

	// 数据加入数据库
	res := "testdata"
	filename := "testcase.xml"
	filepath := abspath + filename

	// 读入文件
	data, _ := ioutil.ReadFile(filepath)

	// 若无对应表则建表
	_ = database.InsertCategory(database.CategoryClient, "word", res)
	_ = database.CreateCategoryTable(database.CategoryClient, res)
	_ = database.CreateFileTable(database.DataClient, res)
	_ = database.CreateTableInDict(res)
	_ = database.InsertFile(database.DataClient, res, filename, data)

	// 数据加入数据库
	filename = "config.json"
	if stage == 1 {
		filename = "config2.json"
	}
	filepath = abspath + filename

	// 读入文件
	data, _ = ioutil.ReadFile(filepath)

	// 加入数据库中
	_ = database.InsertFile(database.ConfigClient, "file", res, data)

	// 数据加入数据库
	filename = "template.json"
	filepath = abspath + filename

	// 读入文件
	data, _ = ioutil.ReadFile(filepath)

	// 加入数据库中
	_ = database.InsertToPattern(res, string(data))

	return nil
}
