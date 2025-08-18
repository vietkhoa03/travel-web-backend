package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"travel-web-backend/internal/service"

	"github.com/gin-gonic/gin"
)

type LocationController struct {
    service service.LocationService
}

func NewLocationController(service service.LocationService) *LocationController {
    return &LocationController{service: service}
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

