package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kiruiaaron/GO-REACT-CALORIE-TRACKER/middlewares"
	"github.com/kiruiaaron/GO-REACT-CALORIE-TRACKER/routes"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(cors.Default())

	router.Use(middlewares.CORSMiddleware())

	router.POST("/entry/create", routes.AddEntry)
	router.GET("/entries", routes.GetEntries)
	router.GET("entry/:id", routes.GetEntryById)
	router.GET("/ingredient/:ingredient", routes.GetEntryByIngredient)
	router.PUT("/entry/update/:id", routes.UpdateEntry)
	router.PUT("/ingredient/update/:id", routes.UpdateIngredient)
	router.DELETE("/entry/delete/:id", routes.DeleteEntry)
	router.Run(":" + port)
}
