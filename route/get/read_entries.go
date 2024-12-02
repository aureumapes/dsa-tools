package get

import (
	"github.com/aureumapes/dsa-tools/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ReadEntries(ctx *gin.Context) {
	var entries []struct {
		Day     int64
		Month   string `gorm:"type:months"`
		Year    int64
		Hour    int64
		Content string
	}

	database.Db.Raw("SELECT d.day,d.month,d.year,d.hour,e.content\nFROM entries AS e\nJOIN dates AS d on d.id = e.date_refer\nORDER BY d.year,d.month,d.day,d.hour;").Find(&entries)

	ctx.HTML(http.StatusOK, "display_entries.gohtml", entries)
}
