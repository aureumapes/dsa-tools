package get

import (
	"fmt"
	"github.com/aureumapes/dsa-tools/database"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"slices"
	"strconv"
	"strings"
)

func ReadDays(ctx *gin.Context) {
	var days []string
	var daysAndEntries []struct {
		Name  string
		Entry []template.HTML
	}
	var dates []database.Date
	database.Db.Find(&dates, "year = ? AND month = ?", ctx.Params.ByName("year"), ctx.Params.ByName("month"))
	for _, date := range dates {
		day := strconv.Itoa(int(date.Day))
		if !slices.Contains(days, day) {
			days = append(days, day)
			var entry database.Entry
			database.Db.Find(&entry, "date_refer = ?", date.ID)
			entries := strings.ReplaceAll(entry.Content, "Tsatag", "<img src=\"/image/tsa.png\" width=\"10px\" style=\"width: 25px;height: 25px; position: fixed;margin-top: -2px;margin-left: -30px;\" title=\"Tsatag\" alt=\"\">")
			entries = strings.ReplaceAll(entries, "Borontag", "<img src=\"/image/boron.png\" width=\"10px\" style=\"width: 25px;height: 25px; position: fixed;margin-top: -2px;margin-left: -30px;\" title=\"Boronstag\" alt=\"\">")
			var entriesHTML []template.HTML
			for _, entryString := range strings.Split(entries, ", ") {
				entriesHTML = append(entriesHTML, template.HTML(entryString))
			}
			dayStruct := struct {
				Name  string
				Entry []template.HTML
			}{
				fmt.Sprintf("%d. %s %dBF", date.Day, date.Month, date.Year),
				entriesHTML,
			}
			daysAndEntries = append(daysAndEntries, dayStruct)
		}
	}
	ctx.HTML(http.StatusOK, "read_days.gohtml", gin.H{
		"DAYS": daysAndEntries,
		"YEAR": ctx.Params.ByName("year"),
		"MOON": ctx.Params.ByName("month"),
	})
}
