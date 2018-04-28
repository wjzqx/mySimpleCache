// util
package cache

import (
	"encoding/json"
	"fmt"

	"log"
	"os"
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
func jsonToRow(val string) (_rc RowCache, err error) {
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

/*********************************文件工具类************************/
type MyFileObj struct {
	FileName    string
	FileContent []byte
}

type mldg struct {
	fn string
	fc []byte
}

/**
 * 文件操作接口
 */
type fileOpInters interface {
	/** 检查文件是否存在*/
	checkFileIsExist() bool
	/** 创建文件*/
	createFile(fileObject *os.File) bool
	/** 创建文件夹*/
	//createFolder(fileObject *os.File) bool
	/** 保存文件内容*/
	saveFile(fileObject *os.File) bool
	/** 删除文件*/
	//delFile(fileObject *os.File) bool

	// 打开文件
	openFile() (fileObject *os.File, _fileErr error)
}

func SaveFileOp(f fileOpInters) {
	var fileObject *os.File
	var b = true
	//var fileErr error

	if f.checkFileIsExist() {
		fileObject, _ = f.openFile()
	} else {
		// 文件不存在，则创建文件
		b = f.createFile(fileObject)
	}
	if b {
		f.saveFile(fileObject)
	}

}

/**
 * 判断文件是否存在，true：存在；false：不存在
 */
func (mf MyFileObj) checkFileIsExist() bool {

	log.Println("文件名称: ", mf.FileName)
	if _, err := os.Stat(mf.FileName); os.IsNotExist(err) {
		log.Println("文件不存在")
		return false
	}

	return true

}

func (mf MyFileObj) openFile() (fileObject *os.File, _fileErr error) {

	fileObject, fileErr := os.OpenFile(mf.FileName, os.O_APPEND, 0666)
	if fileErr != nil {
		log.Println("文件读取失败: ", fileErr)
		return fileObject, fileErr
	}

	log.Println("文件读取成功！")
	return fileObject, fileErr
}

func (mf MyFileObj) createFile(fileObject *os.File) bool {
	var fileErr error

	log.Println("创建文件: ", mf.FileName)

	fileObject, fileErr = os.Create(mf.FileName)

	if fileErr != nil {
		log.Println("文件创建失败: ", fileErr)
		return false
	}

	log.Println("创建文件成功！")
	return true
}

func (mf MyFileObj) saveFile(fileObject *os.File) bool {
	var saveFileFlag = true
	log.Println("文件写入的FileContent：", mf.FileContent)

	fielTemp, fileErr := fileObject.Write(mf.FileContent)

	if fileErr != nil {
		saveFileFlag = false
		log.Println("文件写入失败：", fileErr)
	} else {
		log.Printf("文件写入成功: 写入 %d 个字节n", fielTemp)
	}
	fileObject.Sync()
	// 关闭文件对象
	defer fileObject.Close()

	log.Println("关闭文件对象")

	return saveFileFlag

}
