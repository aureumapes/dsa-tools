package get

import (
	"fmt"
	"github.com/aureumapes/dsa-tools/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"slices"
)

func ReadDays(ctx *gin.Context) {
	var dates []database.Date
	var year, month = ctx.Params.ByName("year"), ctx.Params.ByName("month")
	database.Db.Order("year,month,day").Find(&dates, "year = ? AND month = ?", year, month)

	var days []string
	for _, date := range dates {
		dayString := fmt.Sprint(date.Day)
		if !slices.Contains(days, dayString) {
			days = append(days, dayString)
		}
	}
	ctx.HTML(http.StatusOK, "display_days.gohtml", gin.H{
		"MOON": month,
		"YEAR": year,
		"DAYS": days,
	})
}
