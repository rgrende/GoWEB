package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/home")
	})
	router.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.tmpl", gin.H{
			"title":   "Jukebox-3000",
			"heading": "Search for albums in the links below!",
			"link1":   "Best of the 80s",
			"link2":   "Easy Listening",
			"image":   "/static/img/jukebox.png",
		})
	})
	router.GET("/eighties", func(c *gin.Context) {
		c.HTML(http.StatusOK, "eighties.tmpl", gin.H{
			"title":   "Jukebox-3000",
			"heading": "Best of the 80s",
			"link3":   "Home",
		})
	})
	router.GET("/easy", func(c *gin.Context) {
		c.HTML(http.StatusOK, "easy.tmpl", gin.H{
			"title":   "Jukebox-3000",
			"heading": "Easy Listening",
			"link3":   "Home",
		})
	})

	router.Run(":8080")
}
