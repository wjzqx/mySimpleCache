// configServer project main.go
package main

import (
	"encoding/json"
	"fmt"
	//	"os"

	"github.com/mySimpleCache/cache"
)

func main() {

	//	m2 := make(map[int]int)s
	//	m2[3] = 10
	//	m2[2] = 11
	//	m2[9] = 12
	//	m2[1] = 13
	//	m2[4] = 14
	//	m2[8] = 15
	//	m2[5] = 16
	//	m2[7] = 17
	//	m2[6] = 18

	//	sortM := cache.SortMapKey(m2)

	//	for _, v := range sortM {
	//		fmt.Printf("k=%v, v=%v\n", v.Key, v.Value)

	//	}
	//	fmt.Printf("sortM: %v\n", sortM)

	cache.CreateTable("testTable")
	//test1()

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

	testF := cache.MyFileObj{FileName: "./output1.txt", FileContent: jsonStr}
	fmt.Printf("MyFileObj: %v\n", testF)
	cache.SaveFileOp(testF)
	//cache.Run()
	//running := true
	//	fmt.Scanln(&opName, &queryParam, &targetFlag, &tableName, &termFlag, &termParam)
	//	fmt.Printf("Hi %s %s %s %s %s %s!\n", opName, queryParam, targetFlag, tableName, termFlag, termParam)
	//	fmt.Printf("opName %s\n", opName)
	//	fmt.Printf("queryParam %v\n", queryParam)
	//	fmt.Printf("targetFlag %s\n", targetFlag)
	//	fmt.Printf("tableName %s\n", tableName)
	//	fmt.Printf("termFlag %s\n", termFlag)
	//	fmt.Printf("termParam %s\n", termParam)
	//	log.Println("command", opName)
	//	for running {
	//		fmt.Printf("Hi %s %s %s %s %s %s!\n", opName, queryParam, targetFlag, tableName, termFlag, termParam)
	//		if opName == "stop" {
	//			running = false
	//		}

	//		log.Println("command", opName)
	//	}

	//fmt.Printf("GetSouceAll: %v\n", cache.FormatStrHeadOrTail(a, "[", "]"))
	//fmt.Printf("GetSouceAll: %v\n", cache.FormatStrHeadOrTail(b, "[", "]"))

}

func test1() {
	myCache := make(cache.FieldCache)
	cache.AddSouce(myCache, "A", "a")
	cache.AddSouce(myCache, "B", 1)
	cache.AddSouce(myCache, "C", 23)
	cache.AddSouce(myCache, "D", 12.1)
	cache.AddSouce(myCache, "E", "e")
	cache.AddSouce(myCache, "F", "f")

	myRow := make(cache.RowCache, 0, 5)
	myRow = cache.AddRow(myRow, myCache)

	fmt.Printf("GetSouceAll: %v\n", myRow)
	//i := cache.GetSouce()
	//fmt.Printf("cache is: %v\n", cache)

	c, e := cache.GetRowByIndex(myRow, 0)

	if e != nil {
		fmt.Printf("err: %v\n", e)
	}

	fmt.Printf("c: %v\n", c)

	fmt.Printf("GetSouceAll: %v\n", myRow)

	d, _ := json.Marshal(myRow)

	jsonStr := string(d)
	fmt.Printf("GetSouceAll to jsonStr: %v\n json:%v\n", jsonStr, d)

	myCache1 := make(cache.FieldCache)
	cache.AddSouce(myCache1, "A1", "a1")
	cache.AddSouce(myCache1, "B1", "b1")
	myRow = cache.AddRow(myRow, myCache1)

	fmt.Printf("GetSouceAll: %v\n", myRow)

	myCache2 := make(cache.FieldCache)
	cache.AddSouce(myCache2, "A2", "a2")
	cache.AddSouce(myCache2, "B2", "b2")

	e1 := cache.UpdataRowByIndex(myRow, 1, myCache2)

	if e1 != nil {
		fmt.Printf("err: %v\n", e1)
	}

	fmt.Printf("UpdataRowByIndex: %v\n", myRow)

	myRow, e2 := cache.DelRowByIndex(myRow, 0)
	if e2 != nil {
		fmt.Printf("err: %v\n", e2)
	}

	fmt.Printf("DelRowByIndex: %v\n", myRow)

	fmt.Println("=====================================================================================================")

	cacheInfoTemp := cache.CacheInfo{GroupName: "gN", CloseTime: 20, Persistence: "123"}
	cacheInfoTemp.CreateCache()
	fmt.Printf("GetSouceAll: %v\n", cacheInfoTemp)

	type field map[string]string

	aaaaaa := make(field)
	aaaaaa["a"] = "aaa"
	aaaaaa["b"] = "aaa"
	bbbbbb := make(field)
	bbbbbb["c"] = "aaa"
	cccccc := make(field)
	cccccc["d"] = "aaa"
	cccccc["e"] = "aaa"
	cccccc["f"] = "aaa"

	eeeeee := make(field)
	eeeeee["d"] = "aaa"
	ffffff := make(field)
	ffffff["d"] = "aaa"
	gggggg := make(field)
	gggggg["d"] = "aaa"

	type row []field
	var table = make(map[string]row)
	slice2 := make(row, 0, 5)
	slice2 = append(slice2, aaaaaa, bbbbbb, cccccc, eeeeee, ffffff, gggggg)

	table["myTable"] = slice2
	fmt.Printf("table: %v\n", table)
	d1, _ := json.Marshal(table)

	jsonStr1 := string(d1)

	fmt.Printf("jsonStr: %v\n", jsonStr1)
}
