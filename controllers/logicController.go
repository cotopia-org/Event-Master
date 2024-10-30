package controllers

import (
	"net/http"

	"github.com/cotopia-org/Event-Master/logic"
	"github.com/gin-gonic/gin"
)

func IntersectSegments(c *gin.Context) {
	var segments []logic.LineSegment
	if err := c.ShouldBindJSON(&segments); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := logic.IntersectAll(segments)
	c.JSON(http.StatusCreated, result)
}

func UnionSegments(c *gin.Context) {
	var segments []logic.LineSegment
	if err := c.ShouldBindJSON(&segments); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := logic.UnionAll(segments)
	c.JSON(http.StatusCreated, result)
}


// func ComplementSegments(c *gin.Context) {
// 	var segments []logic.LineSegment
// 	minBound := c.Param("minBound")
// 	maxBound := c.Param("maxBound")
// 	if err := c.ShouldBindJSON(&segments); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	result := logic.ComplementAll(segments, minBound, maxBound)
// 	c.JSON(http.StatusCreated, result)
// }

