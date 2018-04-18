// configServer project main.go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/mySimpleCache/cache"
)

func main() {
	cache.CreateTable("testTable")
	//test1()

	var a = `{"a":"aaa","b":"aaa"},{"a":12}`
	var b = `[{"c":"aaa"}`

	tc, _ := cache.AddParam("testTable", a)
	fmt.Printf("GetSouceAll: %v\n", tc)

	tc, _ = cache.AddParam("testTable", b)
	fmt.Printf("GetSouceAll: %v\n", tc)

	str, _ := cache.QueryParam("testTable", "a", 12)
	fmt.Printf("GetSouceAll: %v\n", str)

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
