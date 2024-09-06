package server

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gotcha/config"
)

func Serve() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/config", func(c *gin.Context) {
		if config.Get().App.Environment == "dev" {
			jsonConfig, err := json.Marshal(config.Get())
			if err != nil {
				c.JSON(500, gin.H{
					"error": "Failed to marshal config",
				})
				return
			}

			c.Data(200, "application/json", jsonConfig)
		} else {
			c.JSON(200, gin.H{
				"message": "config is only available in dev environment, sorry bud :p",
			})
		}
	})

	r.Run(fmt.Sprintf(":%d", config.Get().Api.Port))
}
