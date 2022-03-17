package file

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/0014526/gin-demo/util/util"
)

// 上传文件
func UploadFileHandle(context *gin.Context) {
	headerList:=[]string{"image/png"}
	file, err := context.FormFile("file")
	if err != nil {
		log.Panicln("上传文件失败")
	}
	result := util.IsValueInList(file.Header.Get("Content-Type"), headerList)
	if !result {
		log.Printf("文件格式不是图片类型")
	}
	// 使用的时候要用文件路径加文件名
	err = context.SaveUploadedFile(file, "./util/file/"+file.Filename)
	if err != nil {
		log.Panicln("上传文件失败")
	}
}

// 批量上传文件
func UploadManyFile(context *gin.Context) {
	form, err := context.MultipartForm()
	if err != nil {
		context.String(http.StatusBadRequest, fmt.Sprintf("get err %s", err.Error()))
	}
	// 获取所有图片
	files:=form.File["files"]
	// 遍历所有图片
	for _, file := range files {
		// 逐个存
		if err := context.SaveUploadedFile(file, file.Filename); err != nil {
			context.String(http.StatusBadRequest, fmt.Sprintf("upload err %s", err.Error()))
			return
		}
	}
	context.String(http.StatusOK,fmt.Sprintf("upload ok %d files",len(files)))
}