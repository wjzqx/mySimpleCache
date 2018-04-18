/**
 * tableCache
 * 表数据集合对应着数据库中的一张表数据记录
 * 格式如下： [myTable:[map[a:aaa b:aaa] map[c:aaa] map[d:aaa e:aaa f:aaa] map[d:aaa] map[d:aaa] map[d:aaa]]]
 *          或{"myTable":[{"a":"aaa","b":"aaa"},{"c":"aaa"},{"d":"aaa","e":"aaa","f":"aaa"},{"d":"aaa"},{"d":"aaa"},{"d":"aaa"}]}
 * 该方法集合主要处理表数据的操作：针对表的增，删，改，查的操作
 */
package cache

//"errors"

//	"fmt"

type tableCache map[string]RowCache

func addTable(tc tableCache, key string, rc RowCache) (_tc tableCache, err error) {

	if rc == nil {
		return _tc, ErrRowIsNull
	}

	tc[key] = rc
	return tc, err

}

func getTableByKey(tc tableCache, key string) (_rc RowCache, err error) {

	if tc == nil {
		return _rc, ErrTableIsNull
	}

	// 查找键值是否存在
	if _, ok := tc[key]; ok {
		return tc[key], err
	} else {
		return _rc, ErrMapKeyNotFind
	}
}

func updateTableBykey(tc tableCache, key string, rc RowCache) (_tc tableCache, err error) {
	if tc == nil {
		err = ErrTableIsNull
	} else if rc == nil {
		err = ErrRowIsNull
	} else if _, ok := tc[key]; ok {
		// 查找键值是否存在
		tc[key] = rc
		_tc = tc

	} else {
		err = ErrMapKeyNotFind
	}
	return _tc, err
}

func delTableByKey(tc tableCache, key string) (_tc tableCache, err error) {
	if tc == nil {
		err = ErrTableIsNull
	} else if _, ok := tc[key]; ok {
		delete(tc, key)
	} else {
		// 查找键值是否存在
		err = ErrMapKeyNotFind
	}

	return tc, err
}
