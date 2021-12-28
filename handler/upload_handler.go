package handler

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/jerryshell/my-flomo-server/config"
	"github.com/jerryshell/my-flomo-server/model"
	"github.com/jerryshell/my-flomo-server/service"
	"github.com/jerryshell/my-flomo-server/util"
	"log"
	"os"
	"strings"
	"time"
)

var loc, _ = time.LoadLocation("Asia/Shanghai")

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
		file, err := os.Open(filePath)
		if err != nil {
			log.Println("open file"+filePath+" :: ", err)
			continue
		}
		doc, err := goquery.NewDocumentFromReader(file)
		if err != nil {
			log.Println("NewDocumentFromReader file"+filePath+" :: ", err)
			continue
		}
		doc.Find(".memo").Each(func(i int, memoElement *goquery.Selection) {
			var memoTime time.Time
			memoElement.Find(".time").Each(func(i int, timeElement *goquery.Selection) {
				timeStr := strings.TrimSpace(timeElement.Text())
				memoTime, err = time.ParseInLocation("2006-01-02 15:04:05", timeStr, loc)
				if err != nil {
					return
				}
			})

			var memoContent string
			memoElement.Find(".content p").Each(func(i int, p *goquery.Selection) {
				memoContent += strings.TrimSpace(p.Text()) + "\n"
			})
			memoContent = strings.TrimSpace(memoContent)

			id, err := util.NextIDStr()
			if err != nil {
				log.Println("NextIDStr :: ", err)
				return
			}

			var memo = model.Memo{
				BaseModel: model.BaseModel{
					ID:        id,
					CreatedAt: memoTime,
					UpdatedAt: memoTime,
				},
				Content: memoContent,
			}
			log.Println(memo)
			err = service.MemoSave(&memo)
			if err != nil {
				log.Println("MemoSave :: ", err)
				return
			}
		})
		_ = file.Close()
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
