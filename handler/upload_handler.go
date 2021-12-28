package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jerryshell/my-flomo-server/config"
	"log"
	"os"
)

func Upload(c *gin.Context) {
	form, _ := c.MultipartForm()
	uploadFileList := form.File["uploadFileList[]"]
	successSaveFilePathList := make([]string, 0)
	for _, file := range uploadFileList {
		filename := file.Filename
		filePath := config.Data.FileUploadDir + filename

		err := c.SaveUploadedFile(file, filePath)
		if err != nil {
			removeFileList(successSaveFilePathList)
			c.JSON(500, gin.H{
				"success": false,
				"message": "file: [" + filename + "] :: " + err.Error(),
			})
			return
		}
		successSaveFilePathList = append(successSaveFilePathList, filePath)
	}

	for _, filePath := range successSaveFilePathList {
		log.Println("handle filePath: ", filePath)
	}

	removeFileList(successSaveFilePathList)
	c.JSON(200, gin.H{
		"success": true,
		"message": "ok",
	})
}

func removeFileList(filePathList []string) {
	for _, filePath := range filePathList {
		err := os.Remove(filePath)
		if err != nil {
			log.Println("remove " + filePath + " :: " + err.Error())
			continue
		}
	}
}
