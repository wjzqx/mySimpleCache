// fieldType
package cache

import (
	"fmt"

	"github.com/mySimpleCache/util"
)

/**
 * 单条数据信息结构体
 * 字段数据集合对应着数据库中的一条记录
 * 格式如下： [b:aaa a:aaa]或{"a":"aaa","b":"aaa"}
 * 该方法集合主要处理数据的原子化操作：针对字段的增，删，改，查的操作
 */
 
type RowType struct {
	/** 一条记录的唯一ID标识*/
	FieldId int64
	/**创建用户ID*/
	CreateUid int
	/** 创建数据的时间*/
	CreateTime int64
	/** 修改数据的时间*/
	UpdateTime int64
	/**上一条记录的ID，首位数据为0，用于数据排序使用*/
	PrevFieldId int64
	/** 单条数据内容 map:key, val*/
	ContentMap map[string]interface{}
	/** 该条数据是否显示; true 显示， false 隐藏*/
	IsDisplay bool
}

/**
 * 创建一条数据
 * @param uid 用户ID
 * @param mapTamp 创建的数据
 * @return FieldType对象
 */
func CreateParamToRow(uid int, mapTemp map[string]interface{}) (_fieldType RowType, _err error) {
	var fieldId = util.RandInt64()

	fieldTemp := RowType{
		FieldId:     fieldId,
		CreateUid:   uid,
		CreateTime:  util.GetTimestamp(""),
		UpdateTime:  util.GetTimestamp(""),
		PrevFieldId: 0,
		ContentMap:  mapTemp,
	}
	return fieldTemp, _err

}

/**
 * 单条数据追加元素，
 * 如果添加的map中key 与原数据中key重复
 * 则覆盖原数据
 *
 * @return err,执行成功时 err = nil
 */
func (f *RowType) AddParamToRow(mapTemp map[string]interface{}) (err error) {
	contentMap := f.ContentMap

	if contentMap != nil && len(contentMap) > 0 {
		for k, v := range mapTemp {
			contentMap[k] = v
		}

	} else {
		contentMap = mapTemp
	}

	f.UpdateTime = util.GetTimestamp("")

	return err
}

/**
 * 根据key地段值删除对应的value
 * @param key
 *
 * @return  error
 */
func (f *RowType) DelParamToField(key string) (err error) {
	contentMap := f.ContentMap

	// 查找键值是否存在
	if _, ok := contentMap[key]; ok {
		delete(contentMap, key)
		f.UpdateTime = util.GetTimestamp("")
		return err

	} else {
		return ErrMapKeyNotFind
	}
}

func test() {
	fmt.Println("test is true")
}
