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
	// Setze den Gin-Modus auf Release
	gin.SetMode(gin.ReleaseMode)

	// Erstelle einen neuen Router
	router := gin.New()

	// Erstelle das Logfile
	logfile, _ := os.Create("dsa.log")
	defer logfile.Close()

	// Füge Logger-Middleware hinzu
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

	// Füge Recovery-Middleware hinzu (bei Panik wird automatisch eine Fehlerseite angezeigt)
	router.Use(gin.Recovery())

	// Automatische Migration der Datenbank-Modelle
	database.Db.AutoMigrate(&database.Date{})
	database.Db.AutoMigrate(&database.Entry{})
	database.Db.AutoMigrate(&database.Jubilee{})
	database.Db.AutoMigrate(&database.Adventure{})
	database.Db.AutoMigrate(&database.Char{})
	database.Db.AutoMigrate(&database.File{})

	// Lade HTML-Dateien
	router.LoadHTMLGlob("resource/static/*/*html")

	// Definiere eine NoRoute-Funktion für 404-Fehler
	router.NoRoute(func(context *gin.Context) {
		context.HTML(http.StatusNotFound, "404.html", gin.H{})
	})

	// Definiere POST-Routen
	router.POST("/upload", post.UploadFilePost)

	// Definiere PUT-Routen
	router.PUT("/save", post.Save)
	router.PUT("/entry", post.Entry)
	router.PUT("/newchar", post.NewCharPut)

	// Definiere GET-Routen
	router.GET("/entries", get.ReadEntries)
	router.GET("/upload", get.UploadFileGet)
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})
	router.GET("/time", get.MastersCalender)
	router.GET("/newchar", get.NewCharGet)
	router.GET("/favicon.ico", func(context *gin.Context) {
		context.Redirect(http.StatusPermanentRedirect, "/image/dsa.png")
	})
	router.GET("/dates", get.ReadYears)

	// Definiere Gruppenrouten
	router.Group("/adventure/").
		GET("/", get.Adventures).
		GET("/:adv", get.Adventure)

	router.Group("/dates").
		GET("/", get.ReadYears).
		GET("/:year", get.ReadMonths).
		GET("/:year/:month", get.ReadDays).
		GET("/:year/:month/:day", get.ReadDay)

	router.Group("/chars").
		GET("/", get.Chars).
		GET("/:name", get.Char)

	router.Group("/files").
		Any("/*file", get.Files)

	// Definiere statische Routen
	router.StaticFS("/image", http.Dir("resource/static/images"))
	router.StaticFS("/wasm", http.Dir("resource/static/wasm"))

	// Starte den Server auf Port 4921
	router.Run(":4921")
}
