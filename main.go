package main

import (
	"github.com/aureumapes/dsa-tools/database"
	"github.com/aureumapes/dsa-tools/route"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	database.Db.AutoMigrate(&database.Date{})
	database.Db.AutoMigrate(&database.Entry{})
	database.Db.AutoMigrate(&database.Jubilee{})
	database.Db.AutoMigrate(&database.Adventure{})
	database.Db.AutoMigrate(&database.Char{})

	router.LoadHTMLGlob("resource/static/*html")

	router.PUT("/save", route.Save)
	router.PUT("/entry", route.Entry)
	router.PUT("/newchar", route.NewCharPut)

	router.GET("/chars", route.Chars)
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})
	router.GET("/time", route.Time)
	router.GET("/newchar", route.NewCharGet)
	router.GET("/favicon.ico", func(context *gin.Context) {
		context.Redirect(http.StatusPermanentRedirect, "/image/dsa.png")
	})
	router.GET("/char/:name", route.Char)
	router.GET("/dates", route.EntryRead)
	router.GET("/jubilee", route.ReadJubilees)
	router.GET("/dates/:year", route.ReadMonths)
	router.GET("/dates/:year/:month", route.ReadDays)
	router.GET("/adventure/", route.Adventures)
	router.GET("/adventure/:adv", route.Adventure)

	router.StaticFS("/image", http.Dir("resource/images"))
	router.StaticFS("/wasm", http.Dir("resource/static/wasm"))
	router.Run(":4921")
}
