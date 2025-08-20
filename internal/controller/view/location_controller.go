package controller

import (
	"fmt"
	"net/http"
	"strconv"

	model "travel-web-backend/internal/entity"
	"travel-web-backend/internal/service"

	"github.com/gin-gonic/gin"
)

type LocationController struct {
    service service.LocationService
}

func NewLocationController(service service.LocationService) *LocationController {
    return &LocationController{service: service}
}

func (c *LocationController) GetAllLocation(ctx *gin.Context){
    location, err := c.service.GetAllLocation(ctx)

    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
    }

    ctx.JSON(http.StatusOK, gin.H{"data": location})
}

func (c *LocationController) CreateLocation(ctx *gin.Context){
    var location model.Location
    if err := ctx.ShouldBindJSON(&location); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    created, err := c.service.CreateLocation(ctx, location)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusCreated, gin.H{"data": created})
}

func(c *LocationController) UpdateLocation(ctx *gin.Context){
    id := ctx.Param("id")
    var location model.Location
    if err := ctx.ShouldBindJSON(&location); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    updated, err := c.service.UpdateLocation(ctx, id, location)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"data": updated})
}

func(c *LocationController) DeleteLocation(ctx *gin.Context){
    id := ctx.Param("id")
    err := c.service.DeleteLocation(ctx, id)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"message": "Location deleted successfully"})
}

func (c *LocationController) SearchLocationsByName(ctx *gin.Context) {
    name := ctx.Query("name")
    fmt.Println("Search param name:", name)

    page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
    limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

    results, err := c.service.SearchLocationsByName(ctx, name, page, limit)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "data":  results,
        "page":  page,
        "limit": limit,
        "count": len(results),
    })
}

func (c *LocationController) GetLocationByID(ctx *gin.Context) {
    id := ctx.Param("id")

    location, err := c.service.GetLocationByID(ctx, id)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{
            "error": err.Error(),
        })
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "data": location,
    })
}

