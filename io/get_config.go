// @title		GetConfig
// @description	此函数的用途为，根据数据 "类型"，在配置文件中找出对应的 "写入行为"，并反馈相关数据到数据处理函数中。
// @auth		ryl				2022/3/17		10:00
// @param		targetResource	string			特型卡片类型（如 "诗词" 和 "车" 等）
// @return		itemSettings	[]ItemSetting	此键值下所有需要写入数据库的数据
// @return		err				error			错误值

package io

import (
	"dianasdog/path"
	"fmt"
	"io/ioutil"

	"github.com/tidwall/gjson"
)

// ItemSetting	保存 存入数据库的数据在文件中的路径 和 需要储入的数据库
type ItemSetting struct {
	ItemPath      string // 存入数据库的资料路径
	DumpDigest    bool   // 本字段是否需要 dump 摘要 (Redis)
	DumpInvertIdx bool   // 本字段是否需要 dump 倒排 (ES)
	DumpDict      bool   // 本字段是否需要 dump 词表 (Dict)
}

func GetConfig(targetResource string) ([]ItemSetting, error) {

	// 得到此文件的绝对路径
	abspath, _ := path.GetAbsPath()

	// 查找对应类型的 config 文档路径
	filepath := abspath + "config/" + targetResource + ".json"
	fmt.Println(filepath)
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	json := string(file)
	settings := gjson.Get(json, "write_setting|@pretty")

	var itemSettings = make([]ItemSetting, 0)

	// 在配置文件中查找可能的配置
	settings.ForEach(func(key, value gjson.Result) bool {

		var item ItemSetting
		item.ItemPath = key.String()

		// 读取此路径下的 dump 信息
		value.ForEach(func(key, value gjson.Result) bool {
			switch key.String() {
			case "dump_digest":
				item.DumpDigest = value.Bool()
			case "dump_dict":
				item.DumpDict = value.Bool()
			case "dump_invert_idx":
				item.DumpInvertIdx = value.Bool()
			}
			return true
		})

		// 更新配置数组
		itemSettings = append(itemSettings, item)
		return true
	})

	// 查找成功，返回数组
	return itemSettings, nil
}