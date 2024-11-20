package get

import (
	"fmt"
	"github.com/aureumapes/dsa-tools/database"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func ReadJubilees(ctx *gin.Context) {
	var jubilees []database.Jubilee
	jubileeNest := make(map[string]string)
	returnJubilee := ""
	database.Db.Order("year,month,day").Find(&jubilees)
	for _, jubilee := range jubilees {
		jubileeNest[fmt.Sprintf("%d-%s-%d", jubilee.Year, jubilee.Month, jubilee.Day)] = jubilee.Name
	}
	year := int64(1000)
	month := ""
	for _, jubilee := range jubilees {
		currYear := jubilee.Year
		currMonth := jubilee.Month
		currDay := jubilee.Day
		if currYear != year {
			returnJubilee += fmt.Sprintf("%d<br>", currYear)
			year = currYear
			month = ""
		}
		if currMonth != month {
			returnJubilee += "&nbsp;&nbsp;&nbsp;&nbsp;" + currMonth + "<br>"
			month = currMonth
		}
		returnJubilee += fmt.Sprintf("&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;%d - %s<br>", currDay, jubilee.Name)
	}
	ctx.HTML(http.StatusOK, "display_jubilees.gohtml", template.HTML(returnJubilee))
}
