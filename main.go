// configServer project main.go
package main

import (
	"encoding/json"
	"fmt"
	//	"os"

	"github.com/mySimpleCache/cache"
	"github.com/mySimpleCache/util"
)

func main() {

	// test()
	var map1 = make(map[string]interface{})
	map1["a1"] = "a1"
	map1["b1"] = 123

	var map2 = make(map[string]interface{})
	map2["a2"] = "a2"
	map2["b2"] = 456

	var map3 = make(map[string]interface{})
	map3["a3"] = "a3"
	map3["b3"] = 789

	f1, _ := cache.CreateParamToRow(123, map1)
	fmt.Printf("map1: %v\n", f1)
	f1.AddParamToRow(map2)
	fmt.Printf("map1: %v\n", f1)
	f2, _ := cache.CreateParamToRow(123, map2)
	fmt.Printf("map2: %v\n", f2)

	f3, _ := cache.CreateParamToRow(123, map3)
	fmt.Printf("map3: %v\n", f3)

	var fArray = make([]cache.RowType, 0, 1)
	fArray = append(fArray, f1)
	//fArray = append(fArray, f2)

	fos, _ := cache.CreateRowList("test", fArray)

	fmt.Printf("CreateRowList is :%+v\n", fos)
	var fArray1 = make([]cache.RowType, 0, 1)
	fArray1 = append(fArray1, f2)
	fmt.Printf("fArray1: %v\n", fArray1)
	fos.AddParamRowList(fArray1)
	fmt.Printf("AddParamRowList is :%+v\n", fos)
	fmt.Printf("f3: %v\n", f3)
	fos.UpdataParamRowList(f3, 0)
	fmt.Printf("UpdataParamRowList is :%+v\n", fos)
	fos.DelParamRow(1)
	fmt.Printf("DelParamRow is :%+v\n", fos)
}

func test1() {
	cache.CreateTable("testTable")

	var a = `{"a":15,"b":"aaa"},{"a":12}`
	var b = `[{"c":"aaa"}`

	tc, _ := cache.AddParam("testTable", a)
	fmt.Printf("AddParam: %v\n", tc)

	tc, _ = cache.AddParam("testTable", b)
	fmt.Printf("AddParam: %v\n", tc)

	ascstr, _ := cache.QueryParamByTableSort("testTable", "a", "asc")
	fmt.Printf("QueryParamByTableSort: %v\n", ascstr)
	descstr, _ := cache.QueryParamByTableSort("testTable", "a", "desc")
	fmt.Printf("QueryParamByTableSort: %v\n", descstr)
	jsonStr, _ := json.Marshal(descstr)

	testF := util.MyFileObj{FileName: "./output1.txt", FileContent: jsonStr}
	fmt.Printf("MyFileObj: %v\n", testF)
	//util.SaveFileOp(testF)
	var mun = util.RandInt64()
	fmt.Printf("mun: %v\n", mun)
	var timestamp = util.GetTimestamp("")
	fmt.Printf("time: %v\n", timestamp)
}
