// server
package cache

import (
	"fmt"
)

var table tableCache

/**
 *  创建一张表空间
 *
 */
func CreateTable(talbleName string) (tc tableCache, err error) {
	fmt.Println("CreateTable")
	if table == nil {
		table = make(tableCache)
	}

	// 查找键值是否存在
	if _, ok := table[talbleName]; ok {
		return tc, ErrMapKeyExistence
	} else {

		table[talbleName] = nil
		fmt.Printf("table: %v\n", table)
		return table, err
	}
}

/**
 *  添加一条数据
 *  val json字符串
 *  return tableCache
 */
func AddParam(talbleName string, val string) (tc tableCache, err error) {

	oldRc, _ := getTableByKey(table, talbleName)

	newRc, _err := JsonToRow(val)

	if oldRc != nil && len(oldRc) > 0 {

		fmt.Printf("oldRc: %v\n", oldRc)
		oldRc = append(oldRc, newRc...)
		table, _ = addTable(table, talbleName, oldRc)
	} else {
		table, _ = addTable(table, talbleName, newRc)
	}

	return table, _err
}

/**
* 查询一条数据
* talbleName 表名
* key 查询条件
  val 查询条件字段
* return
*/
func QueryParam(talbleName string, key string, val interface{}) (_val string, _err error) {
	rc, err1 := getTableByKey(table, talbleName)

	var rcTemp = make(RowCache, 0, 100)

	for _, v := range rc {
		vTemp, err2 := GetSouceByKey(v, key)
		fmt.Printf("vTemp: %v\n", vTemp)
		fmt.Printf("val: %v\n", val)
		fmt.Printf("err2: %v\n", err2)

		valTemp, _ := ToJsonStr(val)
		fmt.Printf("vTemp vs val: %v\n", (vTemp == valTemp))
		if err2 == nil && vTemp == valTemp {

			rcTemp = AddRow(rcTemp, v)

		}
	}

	if len(rcTemp) > 0 {
		jsonStr, err3 := ToJsonStr(rcTemp)
		return jsonStr, err3
	}

	return "", err1
}
