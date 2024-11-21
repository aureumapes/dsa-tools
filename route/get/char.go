package get

import (
	"github.com/aureumapes/dsa-tools/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Char(ctx *gin.Context) {
	name := ctx.Params.ByName("name")
	var char database.Char
	database.Db.Find(&char, "name = ?", name)
	ctx.HTML(http.StatusOK, "display_char.gohtml", char)
}

func Chars(ctx *gin.Context) {
	var chars []database.Char
	database.Db.Find(&chars)
	ctx.HTML(http.StatusOK, "list_chars.gohtml", chars)
}

func NewCharGet(ctx *gin.Context) {
	ctx.HTML(200, "new_char.html", nil)
}
