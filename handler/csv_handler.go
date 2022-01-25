package handler

import (
	"bytes"
	"encoding/csv"
	"github.com/gin-gonic/gin"
	"github.com/jerryshell/my-flomo-server/result"
	"github.com/jerryshell/my-flomo-server/service"
	"github.com/jerryshell/my-flomo-server/util"
	"log"
	"net/http"
)

func CsvExport(c *gin.Context) {
	token := c.Param("token")
	if token == "" {
		c.Abort()
		c.JSON(http.StatusOK, result.TokenError())
		return
	}

	username, err := util.GetUsernameFromJWT(token)
	if err != nil {
		log.Println("util.GetUsernameFromJWT :: err", err)
		c.Abort()
		c.JSON(http.StatusOK, result.TokenErrorWithMessage(err.Error()))
		return
	}

	user, err := service.UserGetByUsername(username)
	if err != nil {
		log.Println("service.UserGetByUsername :: err", err)
		c.Abort()
		c.JSON(http.StatusOK, result.TokenErrorWithMessage(err.Error()))
		return
	}
	if user.ID == "" {
		c.Abort()
		c.JSON(http.StatusOK, result.TokenErrorWithMessage("用户不存在"))
		return
	}

	memoList, err := service.MemoListByUserID(user.ID)
	if err != nil {
		log.Println("service.MemoListByUserID :: err", err)
		c.JSON(http.StatusOK, result.ErrorWithMessage(err.Error()))
		return
	}

	// csv buffer
	bytesBuffer := &bytes.Buffer{}

	// csv writer
	csvWriter := csv.NewWriter(bytesBuffer)

	// csv header
	err = csvWriter.Write([]string{"ID", "CreatedAt", "UpdatedAt", "Content"})
	if err != nil {
		log.Println("csvWriter.Write :: err", err)
		c.JSON(http.StatusOK, result.ErrorWithMessage(err.Error()))
		return
	}

	// csv data
	for _, memo := range memoList {
		err = csvWriter.Write([]string{
			memo.ID,
			memo.CreatedAt.Format("2006-01-02 15:04:05"),
			memo.UpdatedAt.Format("2006-01-02 15:04:05"),
			memo.Content,
		})
		if err != nil {
			log.Println("csvWriter.Write :: err", err)
			continue
		}
	}

	csvWriter.Flush()

	c.Writer.Header().Set("Content-Disposition", "attachment; filename=memo.csv")
	c.Data(http.StatusOK, "text/csv", bytesBuffer.Bytes())
}
