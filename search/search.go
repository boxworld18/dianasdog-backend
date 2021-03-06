// @title	Search
// @description	此函数的用途为搜索与句子相关的信息
// @auth	ryl		2022/4/27	16:00
// @param	query	string		句子
// @return	resList	[]map[string]interface{}	结果

package search

import (
	"dianasdog/database"
	"fmt"
)

func Search(query string) []string {

	// 意图识別
	intentList := IntentionRecognition(query)

	// Query 理解
	result := QueryUnderstanding(intentList, query)
	fmt.Println(result)
	// es搜索
	var docIdList []string

	for item := range result {
		resourceName := result[item].Resource
		var content string
		content = ""
		for x := range result[item].detail {
			if result[item].pattern[x] != "garbage" && result[item].pattern[x] != "intent" {
				content = result[item].detail[x]
				resList, _ := database.SearchFromEs(resourceName, database.EsClient, content)
				for k := range resList {
					docIdList = append(docIdList, resList[k].DocID)
				}
			}
		}

	}

	// 根据得到的 docid 列表向 redis 中查找
	// var resList []map[string]interface{}
	var resList []string

	for _, docid := range docIdList {
		// 从 redis 中查找
		res, _ := database.GetFromRedis(database.RedisClient, docid)

		// res = strings.Replace(res, "\n", "", -1)
		// res = strings.Replace(res, "\r", "", -1)

		fmt.Println(res)
		// 结果转化为 json
		// var result map[string]interface{}
		// json.Unmarshal([]byte(res), &result)
		result := res

		// 加入队尾
		resList = append(resList, result)
	}

	// 返回搜索结果
	return resList
}
