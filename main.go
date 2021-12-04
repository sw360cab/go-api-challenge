package main

import (
	"c3lx/db"
	"c3lx/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	// gin.SetMode(gin.ReleaseMode)

	// This handler will get all given Challenges
	router.GET("/challenges", func(c *gin.Context) {
		var challenges []models.Challenge
		db.Db.Where("available = ?", true).Find(&challenges)
		c.JSON(http.StatusOK, challenges)
	})

	router.POST("/challenge/:id/:action", func(c *gin.Context) {
		var user models.User
		var challenge models.Challenge
		/* if here authorization Middleware passed ... */

		// get username from auth
		username := "Bob"
		db.Db.Where("username = ?", username).First(&user)
		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("invalid id format for %q", c.Param("id")))
			return
		}

		switch action := c.Param("action"); action {
		case "accept":
			result := db.Db.Where("id = ?", id).Find(&challenge)
			if result.RowsAffected == 0 { // 404
				c.String(http.StatusNotFound, fmt.Sprintf("challenge %d not found", id))
			} else if !challenge.Available { // 403
				c.String(http.StatusForbidden, fmt.Sprintf("challenge %q not available", challenge.Name))
			} else { // create challenge and answer 200
				//nolint:errcheck
				db.Db.Model(&user).Association("Challenges").Append([]models.Challenge{challenge})
				c.String(http.StatusOK, fmt.Sprintf("challenge %q accepted!", challenge.Name))
			}
		default:
			c.String(http.StatusBadRequest, fmt.Sprintf("action %q not allowed", action))
		}
	})
	return router
}

func main() {
	r := setupRouter()
	//nolint:errcheck
	r.Run(":8080")
}
