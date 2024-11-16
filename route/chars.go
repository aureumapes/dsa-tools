package route

import (
	"fmt"
	"github.com/aureumapes/dsa-tools/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Chars(ctx *gin.Context) {
	var chars []database.Char
	database.Db.Find(&chars)
	fmt.Printf("%v", chars)
	ctx.HTML(http.StatusOK, "chars.gohtml", chars)
}
