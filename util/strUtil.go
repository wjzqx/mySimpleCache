// strUtil
package util

import (
	//"fmt"

	"encoding/json"
	"strings"
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

func ToJsonStr(arg interface{}) (key string, _err error) {
	b, err := json.Marshal(arg)
	if err != nil {
		return key, err
	}
	return string(b), err
}
