package route

import (
	"github.com/aureumapes/dsa-tools/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Adventures(ctx *gin.Context) {
	var adventures []database.Adventure
	var adventuresStrings []struct {
		Name   string
		Number string
	}
	database.Db.Find(&adventures)
	for _, adventure := range adventures {
		adventuresStrings = append(adventuresStrings, struct {
			Name   string
			Number string
		}{Name: adventure.Name, Number: adventure.Number})
	}
	ctx.HTML(http.StatusOK, "adventures.gohtml", gin.H{
		"ADVENTURES": adventuresStrings,
	})
}
