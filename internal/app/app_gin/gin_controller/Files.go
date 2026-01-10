package gin_controller

import (
	"datathink.top/Waigo/internal/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (gcl *GinController) UploadFile(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	if err != nil {
		fmt.Println("上传文件时报错：", err)
		ctx.JSON(http.StatusOK, gin.H{
			"state":   0,
			"msg":     "error",
			"content": form,
		})
	} else {
		localFilesDir := ctx.PostForm("local_files_dir")
		//if len(localFilesDir) < 1 { // 取软件本地的地址
		//	lf := services.LocalFiles{}
		//	localFilesDir = lf.GetLocalFilesDir()
		//}
		//token := ctx.PostForm("token")
		files := form.File["file"] // file与formData的name一致
		for i := 0; i < len(files); i++ {
			file := files[i]
			// 文件信息
			//filename := file.Filename
			//fileSize := file.Size
			//fileType := file.Header["Content-Type"][0]
			filePath := strings.ReplaceAll(file.Header["Content-Disposition"][0], `form-data; name="file"; filename="`, "") // 可以上传整个目录
			filePath = strings.ReplaceAll(filePath, `"`, "")
			filePath = common.ConvertedPath(filePath)
			// 保存文件的绝对路径
			savePath := localFilesDir + "/" + filePath
			//fmt.Println("单个文件信息：", fileType, savePath)
			// 保存到本地
			err := ctx.SaveUploadedFile(file, savePath)
			if err != nil {
				fmt.Println("保存文件时报错：", savePath, err)
			}
		}
		ctx.JSON(http.StatusOK, gin.H{
			"state":   1,
			"msg":     "YES",
			"content": len(files),
		})
	}

	return
}
