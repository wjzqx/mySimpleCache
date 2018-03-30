/**
 * 字段数据集合对应着数据库中的一条记录
 * 格式如下： [b:aaa a:aaa]或{"a":"aaa","b":"aaa"}
 * 该方法集合主要处理数据的原子化操作：针对字段的增，删，改，查的操作
 */
package cache

import (
	"encoding/json"
	//	"errors"
	"fmt"
)

// 设置字段缓存类型
type FieldCache map[string]string

/** 初始化缓存对象*/
//func init() {

//	// myCache为nil是， make初始化map
//	if FieldCache == nil {
//		FieldCache = make(map[string]string)
//	}
//}

/**
 * 向字段缓存数据中添加字段数据
 * @param key
 * @param value
 *
 * @return error
 */
func AddSouce(fc FieldCache, key string, value string) (err error) {

	// 查找键值是否存在
	if _, ok := fc[key]; ok {
		return ErrMapKeyExistence
	} else {

		fc[key] = value
		return err
	}
}

/**
 * 根据key字段值获取对应的value值
 * @param key
 *
 * @return string, error
 */
func GetSouceByKey(fc FieldCache, key string) (s string, err error) {

	// 查找键值是否存在
	if _, ok := fc[key]; ok {
		_souce := fc[key]

		b, err := json.Marshal(_souce)
		if err != nil {
			return key, err
		}
		fmt.Printf("value: %v/n", string(b))
		return string(b), err

	} else {
		return key, ErrMapKeyExistence
	}

}

/**
 * 根据key地段值删除对应的value
 * @param key
 *
 * @return string, error
 */
func DelSouceByKey(fc FieldCache, key string) (err error) {

	// 查找键值是否存在
	if _, ok := fc[key]; ok {
		delete(fc, key)
		return err

	} else {
		return ErrMapKeyNotFind
	}

}

/**
 * 根据key值修改value
 * @param key
 *
 * @return string, error
 */
func UpdataSouceByKey(fc FieldCache, key string, value string) (err error) {

	// 查找键值是否存在
	if _, ok := fc[key]; ok {
		fc[key] = value
		return err
	} else {
		return ErrMapKeyNotFind
	}
}
