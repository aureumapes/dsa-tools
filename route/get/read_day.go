package get

import (
	"fmt"
	"github.com/aureumapes/dsa-tools/database"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"strings"
)

func ReadDay(ctx *gin.Context) {
	var days []database.Date
	daysHtml := make([]template.HTML, 24)
	var year, month, day = ctx.Params.ByName("year"), ctx.Params.ByName("month"), ctx.Params.ByName("day")

	database.Db.Order("year,month,day").Find(&days, "year = ? AND month = ? AND day = ?", year, month, day)

	for _, date := range days {
		var entry database.Entry
		database.Db.Find(&entry, "date_refer = ?", date.ID)
		entries := strings.Split(entry.Content, ", ")
		for i, entryText := range entries {
			entries[i] = fmt.Sprintf("<span style=\"line-height: 1.5;\">%s</span>", entryText)
		}
		entry.Content = strings.Join(entries, "<br>\n")
		entry.Content = strings.ReplaceAll(entry.Content, "Borontag", "<img src=\"/image/boron.png\" title=\"Boronstag\" alt=\"\">")
		entry.Content = strings.ReplaceAll(entry.Content, "Tsatag", "<img src=\"/image/tsa.png\" title=\"Tsatag\" alt=\"\">\n")
		daysHtml[date.Hour] = template.HTML("<div>" + entry.Content + "</div>")
	}
	for i := 0; i < 24; i++ {
		if daysHtml[i] == "" {
			daysHtml[i] = "---"
		}

	}
	ctx.HTML(http.StatusOK, "display_day.gohtml", gin.H{
		"DAY":     day,
		"MOON":    month,
		"YEAR":    year,
		"Entries": daysHtml,
	})
	return
}
