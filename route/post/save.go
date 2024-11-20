package post

import (
	"github.com/aureumapes/dsa-tools/database"
	"github.com/gin-gonic/gin"
	"io"
	"strconv"
	"strings"
)

func Save(ctx *gin.Context) {
	bytes, _ := io.ReadAll(ctx.Request.Body)
	date := strings.Split(string(bytes), " ")
	day, _ := strconv.ParseInt(strings.Split(date[0], ".")[0], 10, 64)
	monthName := date[1]
	year, _ := strconv.ParseInt(strings.Split(date[2], "BF")[0], 10, 64)
	var exists bool
	database.Db.Model(&database.Date{}).Select("count(*) > 0").Where("day = ? AND month = ? AND year = ?", day, monthName, year).Find(&exists)
	if !exists {
		database.Db.Create(&database.Date{
			Day:   day,
			Month: monthName,
			Year:  year,
		})
	}
}
