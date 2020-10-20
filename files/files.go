package files

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Files struct {
	Id       int64
	User_Id  int64
	FileName string
	Type     string
	Name     string
	Link     string
	Created  time.Time
	Updated  time.Time
}

func FileUploadHandler() func(c *gin.Context) {
	return func(c *gin.Context) {

		var model Files
		err := c.ShouldBindJSON(&model)
		file, err := c.FormFile("file")
		if err != nil {
			fmt.Println(err)
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
			return
		}

		fileSrc := "../uploads/" + filename
		// FILE SAVED
		if err := c.SaveUploadedFile(file, fileSrc); err != nil {

			c.JSON(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		} else {
			err := model.Add
			if err != nil {
				c.JSON(http.StatusInternalServerError, fmt.Sprintf("upload file err: %s", err.Error()))
				return

			} else {
				c.JSON(200, gin.H{
					"success": "files added successfully",
				})
			}
		}
	}
}
