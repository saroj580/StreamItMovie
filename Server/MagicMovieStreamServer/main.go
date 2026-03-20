package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	controller "github.com/saroj580/MagicMovie/Server/MagicMovieStreamServer/controllers"
)

func main(){
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context){
		c.String(200, "Hello, Magic Strem Movie!")
	})

	router.GET("/movies", controller.GetMovies())
	router.GET("/movie/:imdb_id", controller.GetMovie())
	router.POST("/addmovie", controller.AddMovie())
	router.POST("/register", controller.RegisterUser())
	router.POST("/login", controller.LoginUser())

	if err:=router.Run(":8080");err!=nil{
		fmt.Println("Failed to start server", err)
	}
}