package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// SetupRouter prepare  the router to serve  requests
func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(gin.Logger())
	r.Delims("{{", "}}")

	r.Static("/css", "./assets/css")
	r.Static("/images", "./assets/images")
	r.Static("/js", "./assets/js")
	r.LoadHTMLGlob("templates/*")

	r.NoRoute(func(cont *gin.Context) {
		cont.HTML(
			http.StatusOK,
			"404.html",
			gin.H{},
		)
	})

	r.GET("/", func(cont *gin.Context) {
		cont.HTML(
			http.StatusOK,
			"index.html",
			gin.H{
				"home": true,
			},
		)
	})

	r.GET("/contact", func(cont *gin.Context) {
		cont.HTML(
			http.StatusOK,
			"contact.html",
			gin.H{
				"contact": true,
			},
		)
	})

	r.POST("/result", func(cont *gin.Context) {
		cont.HTML(
			http.StatusOK,
			"result.html",
			getNounConjugations(cont.PostForm("ukrainiannoun")),
		)
	})

	return r
}

//GetPort return the port to use in the coneccion
func GetPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	return port
}

func main() {
	port := GetPort()
	router := SetupRouter()
	router.Run(":" + port)
}
