package restaurant

import (
	"cinnanym/database/surreal"
	"cinnanym/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Router(r *gin.Engine) {
	restaurantRouter := r.Group("/restaurant")
	{
		// Get all Restaurants
		restaurantRouter.GET("/all", func(c *gin.Context) {
			var size int
			var page int
			sizeParam := c.Query("size")
			pageParam := c.Query("page")

			var err error
			size, err = strconv.Atoi(sizeParam)
			if err != nil {
				size = 10
			}

			page, err = strconv.Atoi(pageParam)
			if err != nil {
				page = 0
			}

			resp, err := surreal.FindAll(surreal.FindAllPayload{
				Table: "restaurant",
				Size:  size,
				Page:  page,
			})
			if err != nil {
				c.JSON(400, gin.H{
					"error": err.Error(),
				})
				return
			}

			c.JSON(200, resp["result"])
		})
		// Get by ID
		restaurantRouter.GET("/:id", func(c *gin.Context) {
			id := c.Param("id")
			user, err := surreal.FindOne(id)
			if err != nil {
				c.JSON(400, gin.H{
					"error": err.Error(),
				})
				return
			}

			c.JSON(200, user)
		})
		// Create new restaurant
		restaurantRouter.POST("/create", func(c *gin.Context) {
			var restaurant model.Restaurant
			err := c.BindJSON(&restaurant)
			if err != nil {
				c.JSON(400, gin.H{
					"error": "failed to bind restaurant",
				})
				return
			}

			err = surreal.Create[*model.Restaurant](surreal.CreatePayload{
				Table:      "restaurant",
				Identifier: "",
				Data:       restaurant.ToMap(),
			}, &restaurant)

			if err != nil {
				c.JSON(400, gin.H{
					"error": err.Error(),
				})
				return
			}

			c.JSON(200, restaurant)
		})
		// Update restaurant
		restaurantRouter.PUT("/:id", func(c *gin.Context) {
			id := c.Param("id")
			var restaurant model.Restaurant
			err := c.BindJSON(&restaurant)
			if err != nil {
				c.JSON(400, gin.H{
					"error": "failed to bind restaurant",
				})
				return
			}

			err = surreal.Update(surreal.UpdatePayload{
				ID:   id,
				Data: restaurant.ToMap(),
			})

			if err != nil {
				c.JSON(400, gin.H{
					"error": err.Error(),
				})
				return
			}

			c.JSON(200, restaurant)
		})
		// Delete restaurant
		restaurantRouter.DELETE("/:id", func(c *gin.Context) {
			id := c.Param("id")
			err := surreal.Delete(id)
			if err != nil {
				c.JSON(400, gin.H{
					"error": err.Error(),
				})
				return
			}

			c.JSON(200, gin.H{
				"message": "restaurant deleted",
			})
		})
	}
}
