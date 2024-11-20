package get

import (
	"fmt"
	"github.com/aureumapes/dsa-tools/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"slices"
	"strings"
)

func Files(ctx *gin.Context) {
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

func UploadFileGet(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "upload.html", nil)
}
