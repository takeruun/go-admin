package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	router.Static("/plugins", "./templates/plugins")
	router.Static("/dist", "./templates/dist")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index3.html", gin.H{})
	})

	router.Run(":3000")
}
