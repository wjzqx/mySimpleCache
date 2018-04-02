/**
 * rowCache
 * 行数据集合，包含多个字段集合。
 * 该数据集合对应对应数据中的表数据机构。
 * 数据格式入下:[map[B:b C:c D:d E:e F:f A:a] map[A2:a2 B2:b2]]
 * json格式:[{"B":"b", "C":"c", "D":"d", "E":"e", "F":"f", "A":"a"},{"A2":"a2", "B2":"b2"}]
 */
package cache

//"fmt"

// 设置RowCache对象类型-->FieldCache型切片
type RowCache []FieldCache

/**
 * 添加一行字段数据集合
 * @param rc RowCache   行数据集合
 * @param fc FieldCache 字段数据集合
 *
 * @return _rc RowCache 行数据集合
 * 说明：按照官方文档来看，slice类型应该是引用类型，但是实际使用的时候
 *      出现方法内部修改对象，方法外的对象没有变化的问题。后期再试试其
 *      他方法，看能不能解决这个问题，目前暂时采用返回对象方式，给外部
 *      调用赋值。
 */
func AddRow(rc RowCache, fc FieldCache) (_rc RowCache) {

	// fc对象为空时，自动创建一个新FieldCache
	if fc == nil {
		fc = make(FieldCache)
	}

	rc = append(rc, fc)

	return rc
}

/**
 * 获取一行字段数据集合
 * @param rc RowCache    行数据集合
 * @param i  int         行数
 *
 * @return fc FieldCache 字段数据集合
 *         err error     index超出数据的实际范围
 */
func GetRowByIndex(rc RowCache, i int) (fc FieldCache, err error) {

	if i > len(rc) {
		return fc, ErrArrayIndexOutOfRange
	}

	return rc[i], err
}

/**
 * 删除一行字段数据集合
 * @param rc RowCache    行数据集合
 * @param i  int         行数
 *
 * @return fc FieldCache 字段数据集合
 *         err error     index超出数据的实际范围
 */
func DelRowByIndex(rc RowCache, i int) (_rc RowCache, err error) {
	if i > len(rc) {
		return _rc, ErrArrayIndexOutOfRange
	}

	//ss = append(rc[:i], rc[i+1:]...)
	rc = append(rc[:i], rc[i+1:]...)

	return rc, err
}

/**
 * 修改一行字段数据集合
 * @param rc RowCache    行数据集合
 * @param i  int         行数
 *
 * @return err error     index超出数据的实际范围
 */
func UpdataRowByIndex(rc RowCache, i int, fc FieldCache) (err error) {
	if i > len(rc) {
		return ErrArrayIndexOutOfRange
	}

	rc[i] = fc

	return err
}
