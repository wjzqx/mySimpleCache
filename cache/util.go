// util
package cache

import (
	"encoding/json"
	"fmt"
	"strings"
)

/**
 * 校验对象是否为空
 * @param args... 需要校验的对象，可以为多个。
 * @return err 校验的错误信息，没有错误为nil
 */
func validParam(args ...interface{}) (err error) {
	fmt.Println("validParam")
	for _, param := range args {
		switch vtype := param.(type) {
		case FieldCache:
			if len(vtype) == 0 {
				return ErrFieldIsNull
			}
		case RowCache:
			if len(vtype) == 0 {
				return ErrRowIsNull
			}
		}
	}

	return err
}

/**
 * 数据转换 json字符串 -> RowCache类型
 */
func JsonToRow(val string) (_rc RowCache, err error) {
	var rc RowCache
	val = FormatStrHeadOrTail(val, "[", "]")
	var _val = []byte(val)
	err = json.Unmarshal(_val, &rc)
	fmt.Printf("JsonToRow %v\n", rc)
	return rc, err
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
