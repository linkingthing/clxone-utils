package excel

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/linkingthing/cement/log"
)

const (
	UploadDirectoryKey = "directory"
	UploadFileKey      = "path"
	UploadFileName     = "filename"
	FileResourceName   = "files"
	IgnoreAuditLog     = "ignoreAuditLog"
)

var FileRootPath = "/opt/files"

func RegisterFileApi(router *gin.Engine, apiPath string) {
	router.StaticFS(path.Join(apiPath, FileResourceName), http.Dir(FileRootPath))
	router.POST(path.Join(apiPath, FileResourceName), UploadFiles)
}

func UploadFiles(ctx *gin.Context) {
	ctx.Writer.Header().Add(IgnoreAuditLog, IgnoreAuditLog)
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

func RemoveFile(fileName string) error {
	if fileName != "" {
		if f, _ := os.Stat(path.Join(FileRootPath, fileName)); f != nil {
			return os.Remove(path.Join(FileRootPath, fileName))
		}
	}

	return nil
}

func CreateUploadFolder(folderName string) error {
	if _, err := os.Stat(path.Join(FileRootPath, folderName)); os.IsNotExist(err) {
		if err = os.Mkdir(path.Join(FileRootPath, folderName), 0777); err != nil {
			return fmt.Errorf("createFolder %s failed:%s ", folderName, err.Error())
		}
	}

	return nil
}

func NormalizeFilepath(fileName string) (string, error) {
	if len(fileName) == 0 {
		return "", fmt.Errorf("empty file")
	}

	if strings.Contains(fileName, "../") {
		return "", fmt.Errorf("file name invalid with path traversal attacks")
	}

	return path.Join(FileRootPath, fileName), nil
}
