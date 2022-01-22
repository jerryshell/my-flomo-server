package handler

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/jerryshell/my-flomo-server/model"
	"github.com/jerryshell/my-flomo-server/result"
	"github.com/jerryshell/my-flomo-server/service"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var loc, _ = time.LoadLocation("Asia/Shanghai")

func Upload(c *gin.Context) {
	user := c.MustGet("user").(*model.User)

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusOK, result.ErrorWithMessage(err.Error()))
		return
	}

	uploadFileList := form.File["uploadFileList[]"]
	successSaveFilePathList := make([]string, 0)
	for _, file := range uploadFileList {
		src, err := file.Open()
		if err != nil {
			log.Println(err)
			continue
		}

		doc, err := goquery.NewDocumentFromReader(src)
		if err != nil {
			log.Println(err)
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

			_, err = service.MemoCreate(memoContent, user.ID)
			if err != nil {
				log.Println("MemoSave :: error", err)
				return
			}
		})
	}

	removeFileList(successSaveFilePathList)
	c.JSON(http.StatusOK, result.Success())
}

func removeFileList(filePathList []string) {
	for _, filePath := range filePathList {
		err := os.Remove(filePath)
		if err != nil {
			log.Println("remove "+filePath+" :: error", err)
			continue
		}
	}
}
