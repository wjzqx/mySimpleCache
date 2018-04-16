// tableCache
package cache

//"errors"

//	"fmt"

type tableCache map[string]RowCache

func addTable(tc tableCache, rc RowCache) (_tc tableCache, err error) {

	if rc == nil {
		return _tc, ErrRowIsNull
	}

	return _tc, err

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
