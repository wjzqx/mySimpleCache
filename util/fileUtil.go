// fileUtil
package util

import (
	"log"
	"os"
	//"strings"
)

type MyFileObj struct {
	FileName    string
	FileContent []byte
}

/**
 * 文件操作接口
 */
type fileOpInters interface {
	/** 检查文件是否存在*/
	checkFileIsExist() bool
	/** 创建文件*/
	createFile() (fileObject *os.File, _fileErr error)
	// 打开文件
	openFile() (fileObject *os.File, _fileErr error)
	/** 创建文件夹*/
	//createFolder(fileObject *os.File) bool
	/** 保存文件内容*/
	saveFile(fileObject *os.File) bool
	/** 删除文件*/
	//delFile(fileObject *os.File) bool
	// 读取文件内容

}

func SaveFileOp(f fileOpInters) {
	var fileObject *os.File
	var fileErr error
	//
	if f.checkFileIsExist() {
		fileObject, fileErr = f.openFile()
	} else {
		// 文件不存在，则创建文件
		fileObject, fileErr = f.createFile()
	}
	if fileErr == nil {
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

func (mf MyFileObj) createFile() (fileObject *os.File, _fileErr error) {

	log.Println("创建文件: ", mf.FileName)

	fileObject, fileErr := os.Create(mf.FileName)

	if fileErr != nil {
		log.Println("文件创建失败: ", fileErr)
		return fileObject, fileErr
	}

	log.Println("创建文件成功！")
	return fileObject, fileErr
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
