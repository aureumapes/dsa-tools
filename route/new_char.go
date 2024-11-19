package route

import (
	"encoding/json"
	"github.com/aureumapes/dsa-tools/database"
	"github.com/gin-gonic/gin"
	"io"
)

func NewCharGet(ctx *gin.Context) {
	ctx.HTML(200, "new_char.html", nil)
}

func NewCharPut(ctx *gin.Context) {
	body, _ := io.ReadAll(ctx.Request.Body)
	var char database.Char
	err := json.Unmarshal(body, &char)
	if err != nil {
		println(err.Error())
		return
	}
	database.Db.Create(&char)
}
