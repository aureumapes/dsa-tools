package main

import (
	"fmt"
	"github.com/aureumapes/dsa-tools/database"
	"github.com/aureumapes/dsa-tools/route/get"
	"github.com/aureumapes/dsa-tools/route/post"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	logfile, _ := os.Create("dsa.log")
	defer logfile.Close()
	router.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		Output: logfile,
		Formatter: func(param gin.LogFormatterParams) string {
			return fmt.Sprintf("[GIN] %s 	|%s\u001B[1;30m %3d %s	|%s %s %s	| %s\n",
				param.ClientIP,
				param.StatusCodeColor(), param.StatusCode, param.ResetColor(),
				param.MethodColor(), param.Method, param.ResetColor(),
				param.Path,
			)
		},
	}))
	router.Use(gin.Recovery())

	database.Db.AutoMigrate(&database.Date{})
	database.Db.AutoMigrate(&database.Entry{})
	database.Db.AutoMigrate(&database.Jubilee{})
	database.Db.AutoMigrate(&database.Adventure{})
	database.Db.AutoMigrate(&database.Char{})
	database.Db.AutoMigrate(&database.File{})

	router.LoadHTMLGlob("resource/static/*/*html")
	router.NoRoute(func(context *gin.Context) { context.HTML(http.StatusNotFound, "404.html", gin.H{}) })

	router.POST("/upload", post.UploadFilePost)

	router.PUT("/save", post.Save)
	router.PUT("/entry", post.Entry)
	router.PUT("/newchar", post.NewCharPut)

	router.GET("/upload", post.UploadFileGet)
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})
	router.GET("/time", post.Time)
	router.GET("/newchar", post.NewCharGet)
	router.GET("/favicon.ico", func(context *gin.Context) { context.Redirect(http.StatusPermanentRedirect, "/image/dsa.png") })
	router.GET("/dates", get.ReadYears)
	router.GET("/jubilee", get.ReadJubilees)

	router.Group("/adventure/").
		GET("/", get.Adventures).
		GET("/:adv", get.Adventure)
	router.Group("/dates").
		GET("/", get.ReadYears).
		GET("/:year", get.ReadMonths).
		GET("/:year/:month", get.ReadDays)
	router.Group("/chars").
		GET("/", get.Chars).
		GET("/:name", get.Char)
	router.Group("/files").
		Any("/*file", post.FileGet)

	router.StaticFS("/image", http.Dir("resource/static/images"))
	router.StaticFS("/wasm", http.Dir("resource/static/wasm"))

	router.Run(":4921")
}
