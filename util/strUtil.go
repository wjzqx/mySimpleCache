// strUtil
package util

import (
	//"fmt"

	"encoding/json"

	"strings"
	"time"
	//"github.com/mySimpleCache/cache"
)

/**
 * 数据转换 json字符串 -> RowCache类型
 */
func JsonToRow(val string, v interface{}) (_err error) {
	//var rc RowCache
	val = FormatStrHeadOrTail(val, "[", "]")
	var _val = []byte(val)
	err := json.Unmarshal(_val, v)
	//fmt.Printf("JsonToRow %v\n", rc)
	return err
}

/**
 *
 *
 */
func FormatStrHeadOrTail(val string, headStr string, tailStr string) string {
	if !strings.HasPrefix(val, headStr) {
		val = headStr + val
	}
	if !strings.HasSuffix(val, tailStr) {
		val = val + tailStr
	}

	return val
}

/**
 * 对象转json字符串
 * @param arg 需要转换的对象
 * @return key json字符串
 *         err 转换时错误信息
 */
func ToJsonStr(arg interface{}) (key string, _err error) {
	b, err := json.Marshal(arg)
	if err != nil {
		return key, err
	}
	return string(b), err
}

/**
 * 获取时间戳
 * return 毫秒
 */
func GetTimestamp(dataStr string) (_timeStamp int64) {
	var timeStamp int64
	if dataStr == "" {
		cur := time.Now()
		//timeStamp = cur.UTC().UnixNano()
		timeStamp = cur.UnixNano() / 1000000 //UnitNano获取的是纳秒，除以1000000获取毫秒级的时间戳
	} else {
		timeLayout := "2006-01-02 15:04:05"
		loc, _ := time.LoadLocation("Local")                         //重要：获取时区
		theTime, _ := time.ParseInLocation(timeLayout, dataStr, loc) //使用模板在对应时区转化为time.time类型
		timeStamp = theTime.UnixNano() / 1000000
	}

	return timeStamp
}
