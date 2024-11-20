package get

import (
	"github.com/aureumapes/dsa-tools/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"slices"
)

func ReadMonths(ctx *gin.Context) {
	months := []string{}
	var dates []database.Date
	database.Db.Order("month, day").Find(&dates, "year = ?", ctx.Params.ByName("year"))
	for _, date := range dates {
		if !slices.Contains(months, date.Month) {
			months = append(months, date.Month)
			months = append(months)
		}
	}
	ctx.HTML(http.StatusOK, "display_godrun.gohtml", gin.H{"MOONS": months, "YEAR": ctx.Params.ByName("year")})
}
