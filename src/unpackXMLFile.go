// @Title  unpackfile
// @Description  Read XML files in folder and unpack them
// @Author  于沛楠
// @Update  2022/3/16
package unpackfile

import (
	"github.com/beevik/etree"
)

// @title   unpackXMLFile
// @description  unpack large XML file to single item
// @auth    于沛楠       2022/3/16
// @param	fileName     string             the name of XML File to unpack
//			resourceName string				the category name of special card
// @return  itemList     []*etree.Element   XML <item> array (the itemList uses etree from "github.com/beevik/etree")
//			itemCount    int                total count of <item></item>
//		    err          error              non-nil when fileName is wrong
func unpackXMLFile(fileName string, resourceName string) (itemList []*etree.Element, itemCount int, err error) {
	doc := etree.NewDocument()
	if err := doc.ReadFromFile(fileName); err != nil { //wrong fileName
		return nil, 0, err
	}
	root := doc.SelectElement("DOCUMENT")
	itemList = root.SelectElements("item")
	itemCount = len(root.SelectElements("item"))
	return itemList, itemCount, nil
}
