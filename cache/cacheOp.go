// cacheOp
package cache

var tableMap map[string]TableEntity
var tableTemp TableEntity

func CreateTableSer(uid int, tableMapName string, tableName string) (t TableEntity) {
	table := CreateTable(uid, tableName)

	return table
}

/**
 * 重置缓存
 *
 */
func ResetCache() {

}

/**
 * 根据key匹配对应的数据
 */
func QueryRowTableListByKey(key string) (table map[string]TableEntity, err error) {

	return table, err
}

/**
 * 根据key和value组成条件，匹配对应的数据
 */
func QueryRowListByCondition(key string, val string) (rows []RowType, err error) {
	return rows, err
}
