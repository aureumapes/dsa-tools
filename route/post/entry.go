package post

import (
	"github.com/aureumapes/dsa-tools/database"
	"github.com/gin-gonic/gin"
	"io"
)

func Entry(ctx *gin.Context) {
	bytes, _ := io.ReadAll(ctx.Request.Body)
	var date database.Date
	database.Db.Last(&date)
	database.Db.Create(&database.Entry{Content: string(bytes), DateRefer: date.ID})
	if ctx.Request.URL.Query().Get("jubilee") == "true" {
		database.Db.Create(&database.Jubilee{Day: date.Day, Month: date.Month, Year: date.Year, Name: string(bytes)})
	}
}
