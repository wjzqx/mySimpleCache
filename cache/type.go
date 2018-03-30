// type
package cache

import (
	"errors"
	"fmt"
)

type CacheInfo struct {
	GroupName   string
	CloseTime   int
	Persistence string
}

// 声明常量
var (
	/** ErrMapKeyNotFind 查询key值不存 */
	ErrMapKeyNotFind = errors.New("key查询不存在")
	/** ErrMapKeyExistence key已存在 */
	ErrMapKeyExistence = errors.New("key已存在")
	/** ErrArrayIndexOutOfRange index超出数据的实际范围 */
	ErrArrayIndexOutOfRange = errors.New("index超出数据的实际范围")
)

func (c CacheInfo) CreateCache() {

	//myCache = make(map[string]string)
	fmt.Println("this is cacheInfo: %v\n", 123)
}
