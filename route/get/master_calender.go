package get

import (
	"fmt"
	"github.com/aureumapes/dsa-tools/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MastersCalender(ctx *gin.Context) {
	var dateStruct database.Date
	database.Db.Last(&dateStruct)
	date := fmt.Sprintf("%d. %s %dBF", dateStruct.Day, dateStruct.Month, dateStruct.Year)
	time := fmt.Sprintf("%02d:00", dateStruct.Hour)
	var jubileesList []database.Jubilee
	database.Db.Find(&jubileesList)
	var jubilees []string
	for _, jubilee := range jubileesList {
		if jubilee.Day == dateStruct.Day && jubilee.Month == dateStruct.Month {
			year := dateStruct.Year - jubilee.Year
			jubilees = append(jubilees, fmt.Sprintf("%d. %s", year, jubilee.Name))
		}
	}
	ctx.HTML(http.StatusOK, "masters_calender.gohtml", gin.H{"DATE": date, "JUBILEES": jubilees, "HOUR": time})
}
