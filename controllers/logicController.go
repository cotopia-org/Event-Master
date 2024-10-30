package controllers

import (
	"net/http"
	"strconv"

	"github.com/cotopia-org/Event-Master/logic"
	"github.com/gin-gonic/gin"
)

// @BasePath /

// IntersectSegments godoc
// @Summary calculates intersection
// @Schemes list of LineSegments that represents a segment in 1D space with two endpoints (Start, End float64)
// @Description returns the intersection of many line segments
// @Tags logic
// @Accept json
// @Produce json
// @Success 201 {} []LineSegment
// @Router /logic/intersect [post]
func IntersectSegments(c *gin.Context) {
	var segments []logic.LineSegment
	if err := c.ShouldBindJSON(&segments); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := logic.IntersectAll(segments)
	c.JSON(http.StatusCreated, result)
}

// @BasePath /

// UnionSegments godoc
// @Summary calculates union
// @Schemes list of LineSegments that represents a segment in 1D space with two endpoints (Start, End float64)
// @Description returns the union of many line segments
// @Tags logic
// @Accept json
// @Produce json
// @Success 201 {} []LineSegment
// @Router /logic/union [post]
func UnionSegments(c *gin.Context) {
	var segments []logic.LineSegment
	if err := c.ShouldBindJSON(&segments); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := logic.UnionAll(segments)
	c.JSON(http.StatusCreated, result)
}

// @BasePath /

// ComplementSegments godoc
// @Summary calculates complement
// @Schemes list of LineSegments that represents a segment in 1D space with two endpoints (Start, End float64)
// @Description returns the complement of many line segments
// @Tags logic
// @Accept json
// @Produce json
// @Success 201 {} []LineSegment
// @Router /logic/complement/:minBound/:maxBound" [post]
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
