package route

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func UploadFileGet(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "upload.html", nil)
}

func UploadFilePost(ctx *gin.Context) {
	file, _ := ctx.FormFile("file")

	src, _ := file.Open()
	defer src.Close()
	bytes, _ := io.ReadAll(src)
	os.WriteFile("./resource/files/"+file.Filename, bytes, 744)
}
