package get

import (
	"github.com/aureumapes/dsa-tools/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"slices"
	"strconv"
)

func ReadYears(ctx *gin.Context) {
	println(ctx.Params.ByName("year"), ctx.Params.ByName("month"), ctx.Params.ByName("day"))
	if ctx.Params.ByName("year") == "" {
		years := []string{}
		var dates []database.Date
		database.Db.Order("year").Find(&dates)
		for _, date := range dates {
			year := strconv.Itoa(int(date.Year))
			if !slices.Contains(years, year) {
				years = append(years, year)
				years = append(years)
			}
		}
		ctx.HTML(http.StatusOK, "read_years.gohtml", years)
		return
	}
}
