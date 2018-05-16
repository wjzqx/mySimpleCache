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
