package controllers

import (
	"net/http"
	"strconv"

	"hammond/models"
	"hammond/service"

	"github.com/gin-gonic/gin"
)

func RegisteImportController(router *gin.RouterGroup) {
	router.POST("/import/fuelly", fuellyImport)
	router.POST("/import/drivvo", drivvoImport)
	router.POST("/import/generic", genericImport)
}

func fuellyImport(c *gin.Context) {
	bytes, err := getFileBytes(c, "file")
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}
	errors := service.FuellyImport(bytes, c.MustGet("userId").(string))
	if len(errors) > 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": errors})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func drivvoImport(c *gin.Context) {
	bytes, err := getFileBytes(c, "file")
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}
	vehicleId := c.PostForm("vehicleID")
	if vehicleId == "" {
		c.JSON(http.StatusUnprocessableEntity, "Missing Vehicle ID")
		return
	}
	importLocation, err := strconv.ParseBool(c.PostForm("importLocation"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Please include importLocation option.")
		return
	}

	errors := service.DrivvoImport(bytes, c.MustGet("userId").(string), vehicleId, importLocation)
	if len(errors) > 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": errors})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func genericImport(c *gin.Context) {
	var json models.ImportData
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if json.VehicleId == "" {
		c.JSON(http.StatusUnprocessableEntity, "Missing Vehicle ID")
		return
	}
	errors := service.GenericImport(json, c.MustGet("userId").(string))
	if len(errors) > 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": errors})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
