package user

import (
	"cinnanym/database/surreal"
	"cinnanym/model"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	usrRouter := r.Group("/user")
	{
		// Get by ID
		usrRouter.GET("/find/:id", func(c *gin.Context) {
			id := c.Param("id")
			user, err := surreal.FindOne[model.User](id)
			if err != nil {
				c.JSON(400, gin.H{
					"error": err.Error(),
				})
				return
			}

			c.JSON(200, gin.H{
				"user": user,
			})
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

			c.JSON(200, gin.H{
				"user": newUser,
			})
		})
	}
}
