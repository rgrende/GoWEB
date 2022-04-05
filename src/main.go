//This is a simple GO Web Project that simulates a jukebox. It consists of three linked pages
//two stylesheets, and three HTML templates. The project address is http://localhost:8080/home
//you will be redirected to the home page.

package main

//GO imports
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//main function
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
	//GET 80s page
	router.GET("/eighties", func(c *gin.Context) {
		c.HTML(http.StatusOK, "eighties.tmpl", gin.H{
			"title":   "Jukebox-3000",
			"heading": "Best of the 80s",
			"link3":   "Home",
		})
	})
	//GET Easy Listening page
	router.GET("/easy", func(c *gin.Context) {
		c.HTML(http.StatusOK, "easy.tmpl", gin.H{
			"title":   "Jukebox-3000",
			"heading": "Easy Listening",
			"link3":   "Home",
		})
	})
	router.Run(":8080")
}
