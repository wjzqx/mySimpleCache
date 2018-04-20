// server
package cache

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var table tableCache

/**
 *  创建一张表空间
 *
 */
func CreateTable(talbleName string) (tc tableCache, err error) {

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
	log.Println("AddParam start...")
	// 先查询该表下是否存在数据
	oldRc, _ := getTableByKey(table, talbleName)
	log.Println("oldRc-->", oldRc)
	newRc, _err := jsonToRow(val)
	log.Println("newRc:-->", newRc)
	// 如果该表下存在数据，则把新增数拼接到数据后面
	if oldRc != nil && len(oldRc) > 0 {

		oldRc = append(oldRc, newRc...)
		table, _ = addTable(table, talbleName, oldRc)
	} else {
		table, _ = addTable(table, talbleName, newRc)
	}
	log.Println("add after table data:-->", table)
	log.Println("AddParam end...")
	return table, _err
}

/**
* 根据key和val查询数据集合
* talbleName 表名
* key 查询条件字段名称
* val 查询条件字段值
* return
 */
func QueryParamByCondition(talbleName string, key string, val interface{}) (_val string, _err error) {
	rc, err1 := getTableByKey(table, talbleName)

	var rcTemp = make(RowCache, 0, 100)

	for _, v := range rc {
		vTemp, err2 := GetSouceByKey(v, key)
		//valTemp, _ := ToJsonStr(val)
		log.Println("valTemp-->", val)
		log.Println("vTemp-->", vTemp)
		log.Println("err2-->", err2)
		log.Println("vTemp==valTemp", (vTemp == val))
		if err2 == nil && vTemp == val {
			rcTemp = AddRow(rcTemp, v)
		}
	}

	if len(rcTemp) > 0 {
		jsonStr, err3 := ToJsonStr(rcTemp)
		return jsonStr, err3
	}

	return "", err1
}

func Run() {
	consoleInput()
}

/**
 * 运行控制台
 * 命令输入的方式使用缓存数据库
 */
func consoleInput() {
	running := true
	reader := bufio.NewReader(os.Stdin)

	for running {
		data, _, _ := reader.ReadLine()
		command := string(data)
		if command == "stop" {
			running = false
			return
		}

		s := strings.Split(command, " ")
		switch s[0] {
		case "caretTable":
			CreateTable(s[1])
			log.Println("CreateTable success", command)
		case "insertRow":
			tc, _ := AddParam(s[1], s[2])
			log.Println("insertRow success", tc)
		case "query":
			str, err := QueryParamByCondition(s[1], s[2], s[3])
			if err != nil {
				log.Println("query field", err)
			} else {
				log.Println("query success", str)
			}

		}

		log.Println("command", s[0])
	}

}
