package route

import (
	"github.com/aureumapes/dsa-tools/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Adventure(ctx *gin.Context) {
	advNumber := ctx.Params.ByName("adv")
	var adv database.Adventure
	database.Db.Find(&adv, "number = ?", advNumber)
	ctx.HTML(http.StatusOK, "adventure.gohtml", gin.H{
		"NUMBER": adv.Number,
		"NAME":   adv.Name,
		"HEROS":  strings.Split(adv.Heroes, ", "),
		"TIME":   adv.Time,
	})
}
