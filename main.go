package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	customLintingRulesFile = "rules.yaml"
)

func main() {
	r := gin.Default()

	// health check
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.POST("/add-custom-linting-rules", addRules)
	r.POST("/update-custom-linting-rules", updateRules)
	r.GET("/get-custom-linting-rules", getRules)

	if err := r.Run(); err != nil {
		panic(err)
	}
}
