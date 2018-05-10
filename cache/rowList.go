// rowList
package cache

import (
	"fmt"

	"github.com/mySimpleCache/util"
)

/**
 * 表数据集合
 * rowList
 * 行数据集合，包含多个字段集合。
 * 该数据集合对应对应数据中的表数据机构。
 * 数据格式入下:[map[B:b C:c D:d E:e F:f A:a] map[A2:a2 B2:b2]]
 * json格式:[{"B":"b", "C":"c", "D":"d", "E":"e", "F":"f", "A":"a"},{"A2":"a2", "B2":"b2"}]
 */
type rowList struct {
	rowsId     int64
	rows       []RowType
	rowsName   string
	createTime int64
	updataTime int64
}

/**
 * 创建数据集合
 * @param _rowsName row数据集合名称
 * @param _rows     row数据集合
 * @retrun _rowList rowList对象
 *         err      错误信息
 */
func CreateRowList(_rowsName string, _rows []RowType) (_rowList rowList, err error) {

	rows := rowList{
		rowsId:     util.RandInt64(),
		rows:       _rows,
		rowsName:   _rowsName,
		createTime: util.GetTimestamp(""),
		updataTime: util.GetTimestamp(""),
	}
	return rows, err
}

/**
 * 添加数据
 */
func (r *rowList) AddParamRowList(_rows []RowType) (err error) {
	rows := r.rows
	rows = append(rows, _rows...)
	r.updataTime = util.GetTimestamp("")
	return err
}

/**
 * 修改一条数据
 */
func (r *rowList) updataParamRowList(_rowType RowType, index int) (err error) {
	row := r.rows

	if index > len(row) {
		return ErrArrayIndexOutOfRange
	}

	row[index] = _rowType
	return err
}

/**
 * 删除一条数据信息
 */
func (r *rowList) DelParamRow(index int) (err error) {
	rc := r.rows
	if index > len(rc) {
		return ErrArrayIndexOutOfRange
	}

	rc = append(rc[:index], rc[index+1:]...)
	r.updataTime = util.GetTimestamp("")
	return err
}

func test1() {
	fmt.Println("test is true")
}
