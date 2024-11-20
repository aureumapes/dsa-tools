package route

import (
	"fmt"
	"github.com/aureumapes/dsa-tools/database"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"slices"
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
	os.Remove(path)
	ctx.Redirect(http.StatusPermanentRedirect, "/files/"+file.Filename)
}

func FileGet(ctx *gin.Context) {
	fileName := strings.TrimPrefix(ctx.Request.URL.Path, "/files/")
	if strings.Contains(fileName, ".") {
		var file database.File
		if database.Db.Find(&file, "name=?", strings.TrimSuffix(fileName, "/")).RowsAffected == 0 {
			ctx.Status(404)
			return
		}
		var content string
		database.Db.Raw("SELECT lo_get('" + file.Content + "')").Scan(&content)
		ctx.Writer.Write([]byte(content))
	} else {
		var files []database.File
		filesStr := []string{
			"..",
		}
		database.Db.Raw(fmt.Sprintf("SELECT * FROM files WHERE name LIKE '%s%s'", fileName, "%")).Scan(&files)
		for _, file := range files {
			fileN := strings.TrimPrefix(file.Name, fileName)
			if !slices.Contains(filesStr, strings.Split(fileN, "/")[0]+"/") {
				filesStr = append(filesStr, strings.Split(fileN, "/")[0]+"/")
			}
		}
		ctx.HTML(http.StatusOK, "files.gohtml", filesStr)
	}
}
