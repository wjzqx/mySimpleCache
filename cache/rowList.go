package cache

import (
	"github.com/mySimpleCache/util"
)

/**
 * ROWLIST:表数据集合
 *
 * 行数据集合，包含多个字段集合。
 * 该数据集合对应对应数据中的表数据机构。
 * 数据格式入下:[map[B:b C:c D:d E:e F:f A:a] map[A2:a2 B2:b2]]
 * json格式:[{"B":"b", "C":"c", "D":"d", "E":"e", "F":"f", "A":"a"},{"A2":"a2", "B2":"b2"}]
 */
type RowList struct {
	RowsID     int64
	Rows       []RowType
	RowsName   string
	CreateTime int64
	UpdataTime int64
}

/**
 * 创建数据集合
 * @param _rowsName row数据集合名称
 * @param _rows     row数据集合
 * @retrun _rowList rowList对象
 *         err      错误信息
 */
func CreateRowList(_rowsName string, _rows []RowType) (_rowList RowList, err error) {

	rows := RowList{
		RowsID:     util.RandInt64(),
		Rows:       _rows,
		RowsName:   _rowsName,
		CreateTime: util.GetTimestamp(""),
		UpdataTime: util.GetTimestamp(""),
	}
	return rows, err
}

/**
 * 添加数据
 */
func (r *RowList) AddParamRowList(_rows []RowType) (err error) {
	//log.Println("AddParamRowList start...")
	//log.Println(" [_rows]:--->", _rows)
	rows := r.Rows
	//log.Println("[rows]:--->", rows)
	rows = append(rows, _rows...)
	r.Rows = rows
	r.UpdataTime = util.GetTimestamp("")
	//	log.Println("[RowList]:--->", r)
	return err
}

/**
 * 修改一条数据
 */
func (r *RowList) UpdataParamRowList(_rowType RowType, index int) (err error) {
	//log.Println("UpdataParamRowList start...")
	//log.Println(" [_rowType]:--->", _rowType)
	rows := r.Rows

	if index >= len(rows) {
		return ErrArrayIndexOutOfRange
	}
	//log.Println("[rows]:--->", rows)
	rows[index] = _rowType
	r.Rows = rows
	//log.Println("[RowList]:--->", r)
	return err
}

/**
 * 删除一条数据信息
 */
func (r *RowList) DelParamRow(index int) (err error) {
	//log.Println("DelParamRow start...")
	rows := r.Rows
	if index >= len(rows) {
		return ErrArrayIndexOutOfRange
	}
	//log.Println("[rows]:--->", rows)
	rows = append(rows[:index], rows[index+1:]...)
	//log.Println("[rows]:--->", rows)
	r.UpdataTime = util.GetTimestamp("")
	r.Rows = rows
	//log.Println("[RowList]:--->", r)
	return err
}
