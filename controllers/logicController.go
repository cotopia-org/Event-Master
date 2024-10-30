package controllers

import (
	"net/http"
	"strconv"

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

func ComplementSegments(c *gin.Context) {
	var segments []logic.LineSegment
	minBoundFloat := 0.0
	maxBoundFloat := 0.0
	minBound := c.Param("minBound")
	if s, err := strconv.ParseFloat(minBound, 64); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		minBoundFloat = s
	}

	maxBound := c.Param("maxBound")
	if s, err := strconv.ParseFloat(maxBound, 64); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		maxBoundFloat = s
	}

	if minBoundFloat > maxBoundFloat {
		c.JSON(http.StatusBadRequest, gin.H{"error": "minBound is greater than maxBound!"})
		return
	}

	if err := c.ShouldBindJSON(&segments); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := logic.ComplementAll(segments, minBoundFloat, maxBoundFloat)
	c.JSON(http.StatusCreated, result)
}
