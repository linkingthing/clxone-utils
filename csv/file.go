package util

import (
	"github.com/linkingthing/cement/log"
	"net/http"
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func RegisterFileApi(router *gin.Engine, apiPath string) {
	router.StaticFS(path.Join(apiPath, FileResourceName), http.Dir(FileRootPath))
	router.POST(path.Join(apiPath, FileResourceName), UploadFiles)
}

func UploadFiles(ctx *gin.Context) {
	ctx.Request.Header.Set(IgnoreAuditLog, IgnoreAuditLog)
	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var directory string
	if len(form.Value[UploadDirectoryKey]) > 0 {
		directory = form.Value[UploadDirectoryKey][0]
	}
	files := form.File[UploadFileKey]
	var fileNames string
	for _, file := range files {
		if err := CreateUploadFolder(directory); err != nil {
			log.Warnf("create upload folder failed:%s", err.Error())
			continue
		}
		filename := path.Join(directory, filepath.Base(file.Filename))
		if err := ctx.SaveUploadedFile(file, path.Join(FileRootPath, filename)); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		fileNames = filename
	}

	ctx.JSON(http.StatusOK, gin.H{
		UploadFileName: fileNames,
	})
}
