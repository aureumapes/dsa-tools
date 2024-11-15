package route

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
