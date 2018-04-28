// sortUtil
package cache

import (
	"fmt"
	//	"reflect"
	"sort"
	//"unsafe"
)

type PersonInt struct {
	Value int
	Key   int
}

type PersonStr struct {
	value int
	key   string
}

type PersonSlice []PersonInt
type PersonSliceDesc []PersonInt

func (s PersonSlice) Len() int           { return len(s) }
func (s PersonSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s PersonSlice) Less(i, j int) bool { return s[i].Key < s[j].Key }

func (s PersonSliceDesc) Len() int           { return len(s) }
func (s PersonSliceDesc) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s PersonSliceDesc) Less(i, j int) bool { return s[i].Key > s[j].Key }

/**
 * 根据MAP的key进行排序（正序）
 */
func SortAscMapKey(mapTemp map[int]int) (_p PersonSlice) {

	var pSlice = make(PersonSlice, 0, 100)
	for k, v := range mapTemp {
		pint := PersonInt{Value: v, Key: k}
		pSlice = append(pSlice, pint)
		fmt.Printf("k=%v, v=%v\n", k, mapTemp[k])
	}
	fmt.Printf("pSlice: %v\n", pSlice)
	sort.Stable(pSlice)
	fmt.Printf("pSlice: %v\n", pSlice)

	return pSlice
}

/**
 * 根据MAP的key进行排序（倒序）
 */
func SortDescMapKey(mapTemp map[int]int) (_p PersonSliceDesc) {

	var pSlice = make(PersonSliceDesc, 0, 100)
	for k, v := range mapTemp {
		pint := PersonInt{Value: v, Key: k}
		pSlice = append(pSlice, pint)
		fmt.Printf("k=%v, v=%v\n", k, mapTemp[k])
	}
	fmt.Printf("pSlice: %v\n", pSlice)
	sort.Stable(pSlice)
	fmt.Printf("pSlice: %v\n", pSlice)

	return pSlice
}
