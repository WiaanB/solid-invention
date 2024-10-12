package user

import (
	"cinnanym/database/surreal"
	"cinnanym/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Router(r *gin.Engine) {
	usrRouter := r.Group("/user")
	{
		// Get all Users
		usrRouter.GET("/all", func(c *gin.Context) {
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
				Table: "user",
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
		usrRouter.GET("/:id", func(c *gin.Context) {
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
		// Create new user
		usrRouter.POST("/create", func(c *gin.Context) {
			var user model.User
			err := c.BindJSON(&user)
			if err != nil {
				c.JSON(400, gin.H{
					"error": "failed to bind user",
				})
				return
			}

			newUser, err := Create(user.Username, user.Name, user.Password, user.Email, user.Role)
			if err != nil {
				c.JSON(400, gin.H{
					"error": err.Error(),
				})
				return
			}

			c.JSON(200, newUser)
		})
		// Update user
		usrRouter.PUT("/:id", func(c *gin.Context) {
			id := c.Param("id")
			var user model.User
			err := c.BindJSON(&user)
			if err != nil {
				c.JSON(400, gin.H{
					"error": "failed to bind user",
				})
				return
			}

			updatedUser, err := Update(id, user)
			if err != nil {
				c.JSON(400, gin.H{
					"error": err.Error(),
				})
				return
			}

			c.JSON(200, updatedUser)
		})
		// Delete user
		usrRouter.DELETE("/:id", func(c *gin.Context) {
			id := c.Param("id")
			err := Delete(id)
			if err != nil {
				c.JSON(400, gin.H{
					"error": err.Error(),
				})
				return
			}

			c.JSON(200, gin.H{
				"message": "user deleted",
			})
		})
	}
}
