package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func GetSubmitById(c *gin.Context) {

}

const resourceDirPath = "resources"

func AddSubmit(c *gin.Context) {
	newSubmit := &Submit{}

	if err := c.ShouldBind(newSubmit); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	if f, err := c.FormFile("beer-pic"); err == nil {
		dst := resourceDirPath + "only-res" + filepath.Ext(f.Filename)
		if err := c.SaveUploadedFile(f, dst); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		newSubmit.DrinkPic = dst
		hasher := sha256.New()
		csf, ferr := os.Open(dst)
		if ferr != nil {
			c.String(http.StatusBadRequest, ferr.Error())
			return
		}
		defer csf.Close()
		if _, err := io.Copy(hasher, csf); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		newSubmit.DrinkPicCheckSum = fmt.Sprintf("%x", hasher.Sum(nil))
	}

	storage = append(storage, newSubmit)
	c.HTML(http.StatusOK, "resp-template.tmpl", newSubmit)
}
