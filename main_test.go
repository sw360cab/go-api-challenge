package main

import (
	"c3lx/db"
	"c3lx/models"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type ChallengeAPI struct {
	Name        string
	Description string
}

var router *gin.Engine

func init() {
	router = setupRouter()

	// create static records
	// Delete permanently
	// cmp. https://gorm.io/docs/delete.html#Delete-permanently
	db.Db.Exec("DELETE FROM user_challenges")
	db.Db.Unscoped().Where("1 = 1").Delete(&models.User{})
	db.Db.Unscoped().Where("1 = 1").Delete(&models.Challenge{})

	var users = []models.User{{Username: "Bob"}, {Username: "Alice"}}
	db.Db.Create(&users)

	var challenges = []models.Challenge{{
		Name:        "Get to Steppin'",
		Description: "Get at least 10,000 steps a day for a month",
	}, {
		Name:        "You are what you eat",
		Description: "Eat two healthy meals a day for two weeks",
	}, {
		Name:        "Rip Van Winkle",
		Description: "Average seven hours of sleep for five nights in a row",
	}}
	db.Db.Create(challenges)
	// unavailable challenge
	db.Db.Model(models.Challenge{}).Create(
		map[string]interface{}{
			"Name":        "Not available challenge",
			"Description": "A challenge that is not available",
			"Available":   false})
}

func TestChallenges(t *testing.T) {
	var challenges []ChallengeAPI

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/challenges", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	json.NewDecoder(w.Body).Decode(&challenges)
	assert.Equal(t, 3, len(challenges))
}

func TestAcceptChallenge(t *testing.T) {
	var availableChallenge, unavailableChallenge models.Challenge
	db.Db.Where("available = ?", true).First(&availableChallenge)
	db.Db.Where("available = ?", false).First(&unavailableChallenge)

	var tests = []struct {
		id     string
		action string
		want   int
	}{
		{"4444", "accept", 404},
		{"ppp", "accept", 400},
		{"224", "refuse", 400},
		{fmt.Sprintf("%d", unavailableChallenge.ID), "accept", 403},
		{fmt.Sprintf("%d", availableChallenge.ID), "accept", 200},
	}
	for _, test := range tests {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/challenge/"+test.id+"/"+test.action, nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, test.want, w.Code)
	}
}
