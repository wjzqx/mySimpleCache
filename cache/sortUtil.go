// sortUtil
package cache

import (
	"fmt"
	"sort"
)

type personInt struct {
	Value int
	Key   int
}

type personStr struct {
	value int
	key   string
}

type personSlice []personInt
type personSliceDesc []personInt

func (s personSlice) Len() int           { return len(s) }
func (s personSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s personSlice) Less(i, j int) bool { return s[i].Key < s[j].Key }

func (s personSliceDesc) Len() int           { return len(s) }
func (s personSliceDesc) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s personSliceDesc) Less(i, j int) bool { return s[i].Key > s[j].Key }

/**
 * 根据MAP的可以进行排序
 */
func SortAscMapKey(mapTemp map[int]int) (_p personSlice) {

	var pSlice = make(personSlice, 0, 100)
	for k, v := range mapTemp {
		pint := personInt{Value: v, Key: k}
		pSlice = append(pSlice, pint)
		fmt.Printf("k=%v, v=%v\n", k, mapTemp[k])
	}
	fmt.Printf("pSlice: %v\n", pSlice)
	sort.Stable(pSlice)
	fmt.Printf("pSlice: %v\n", pSlice)

	return pSlice
}

func SortDescMapKey(mapTemp map[int]int) (_p personSliceDesc) {

	var pSlice = make(personSliceDesc, 0, 100)
	for k, v := range mapTemp {
		pint := personInt{Value: v, Key: k}
		pSlice = append(pSlice, pint)
		fmt.Printf("k=%v, v=%v\n", k, mapTemp[k])
	}
	fmt.Printf("pSlice: %v\n", pSlice)
	sort.Stable(pSlice)
	fmt.Printf("pSlice: %v\n", pSlice)

	return pSlice
}
