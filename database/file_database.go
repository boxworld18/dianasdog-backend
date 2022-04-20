// @title	file_database
// @description	此函数的用途是配置后端文件对应的数据库
// @auth	ryl		2022/4/20	11:30
// @param	t		*testing.T	testing 用参数

package database

import "database/sql"

// 文件数据库接口
var CategoryClient *sql.DB
var DataClient *sql.DB
var ConfigClient *sql.DB
var TemplateClient *sql.DB

func init() {
	// 开启数据库
	CategoryClient, _ = sql.Open("mysql", GenUrl("category"))
	DataClient, _ = sql.Open("mysql", GenUrl("data"))
	ConfigClient, _ = sql.Open("mysql", GenUrl("config"))
	TemplateClient, _ = sql.Open("mysql", GenUrl("template"))

	inittask := "SET NAMES utf8 "

	// 生成特型卡词典
	CategoryClient.Exec(inittask)
	CreateTableFromDict(CategoryClient, "word", []string{"id"})

	// 生成源数据数据库（每个特型卡有多个对应文件）
	DataClient.Exec(inittask)

	// 生成写入行为配置数据库（每个特型卡只有一个对应文件）
	ConfigClient.Exec(inittask)
	CreateFileTable(ConfigClient, "file")

	// 生成模板配置数据库（每个特型卡只有一个对应文件）
	TemplateClient.Exec(inittask)
	CreateFileTable(TemplateClient, "file")
}

type FileStruct struct {
	Filename string `json:filename`
	Data     []byte `json:data`
}

// 新建文件表格（含文件名和内容）
func CreateFileTable(db *sql.DB, tableName string) error {
	task := "CREATE TABLE IF NOT EXISTS " + tableName + " (filename VARCHAR(64) PRIMARY KEY NULL, data LONGBLOB NULL) DEFAULT CHARSET=utf8;"
	_, err := db.Exec(task)
	return err
}

// func InsertFile(db *sql.DB, tableName string, filename string, data []byte) error {
// 	err := db.Exec("INSERT IGNORE INTO " + tableName + " VALUES(?)" + filename + ", " + data)
// 	return err
// }