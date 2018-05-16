package cache

import (
	"github.com/mySimpleCache/util"
)

/**
 * tableType
 * 表数据集合对应着数据库中的一张表数据记录
 * 格式如下： [myTable:[map[a:aaa b:aaa] map[c:aaa] map[d:aaa e:aaa f:aaa] map[d:aaa] map[d:aaa] map[d:aaa]]]
 *          或{"myTable":[{"a":"aaa","b":"aaa"},{"c":"aaa"},{"d":"aaa","e":"aaa","f":"aaa"},{"d":"aaa"},{"d":"aaa"},{"d":"aaa"}]}
 * 该方法集合主要处理表数据的操作：针对表的增，删，改，查的操作
 */
type TableEntity struct {
	TableID    int64
	RLMap      map[string]RowList
	CreateTime int64
	UpdataTime int64
	UID        int
}

func CreateTable(uId int, name string) (t TableEntity) {

	tableMap := make(map[string]RowList)
	//tableMap[name] = rows
	tableTemp := TableEntity{
		TableID:    util.RandInt64(),
		RLMap:      tableMap,
		CreateTime: util.GetTimestamp(""),
		UpdataTime: util.GetTimestamp(""),
		UID:        uId,
	}

	return tableTemp
}

func (t *TableEntity) QueryTableByName(name string) (_rows RowList, err error) {

	// 查找键值是否存在
	if _, ok := t.RLMap[name]; ok {
		return t.RLMap[name], err
	}
	return _rows, ErrMapKeyNotFind

}

func (t *TableEntity) AddTable(key string, rows RowList) (err error) {

	if _, ok := t.RLMap[key]; ok {
		return ErrMapKeyExistence
	}
	t.RLMap[key] = rows
	return err
}

func (t *TableEntity) UpdataTable(key string, rows RowList) (err error) {
	if _, ok := t.RLMap[key]; ok {
		t.RLMap[key] = rows
		return err
	}

	return ErrMapKeyNotFind
}

func (t *TableEntity) DelTable(key string) (err error) {
	if _, ok := t.RLMap[key]; ok {
		delete(t.RLMap, key)
		return err
	}
	return ErrMapKeyNotFind
}
