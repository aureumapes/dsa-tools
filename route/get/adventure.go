package get

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
		"adv":    adv,
		"Heroes": strings.Split(adv.Heroes, ", "),
	})
}
