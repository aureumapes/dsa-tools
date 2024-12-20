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
	text := strings.Split(string(bytes), " - ")
	date := strings.Split(text[0], " ")
	hour, _ := strconv.ParseInt(strings.Split(text[1], ":")[0], 10, 64)
	day, _ := strconv.ParseInt(strings.Split(date[0], ".")[0], 10, 64)
	monthName := date[1]
	year, _ := strconv.ParseInt(strings.Split(date[2], "BF")[0], 10, 64)
	var exists bool
	database.Db.Model(&database.Date{}).Select("count(*) > 0").Where("day = ? AND month = ? AND year = ? AND hour = ?", day, monthName, year, hour).Find(&exists)
	if !exists {
		database.Db.Create(&database.Date{
			Day:   day,
			Month: monthName,
			Year:  year,
			Hour:  hour,
		})
	}
}
