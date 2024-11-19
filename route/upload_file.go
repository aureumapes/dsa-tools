package route

import (
	"fmt"
	"github.com/aureumapes/dsa-tools/database"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"strings"
)

func UploadFileGet(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "upload.html", nil)
}

func UploadFilePost(ctx *gin.Context) {
	file, _ := ctx.FormFile("file")

	src, _ := file.Open()
	defer src.Close()
	bytes, _ := io.ReadAll(src)
	if _, err := os.ReadDir("/tmp/dsa"); os.IsNotExist(err) {
		os.Mkdir("/tmp/dsa", 0777)
	}
	path := "/tmp/dsa/" + file.Filename
	os.WriteFile(path, bytes, 0777)
	var oid int
	database.Db.Raw("SELECT lo_import('" + path + "');").Scan(&oid)
	database.Db.Create(&database.File{
		Name:    file.Filename,
		Content: fmt.Sprintf("%d", oid),
	})
	ctx.Redirect(http.StatusPermanentRedirect, "/files/"+file.Filename)
}

func FileGet(ctx *gin.Context) {
	fileName := strings.TrimPrefix(ctx.Request.URL.Path, "/files/")
	var file database.File
	database.Db.Find(&file, "name=?", fileName)
	var content string
	database.Db.Raw("SELECT lo_get('" + file.Content + "')").Scan(&content)
	ctx.Writer.Write([]byte(content))
}
