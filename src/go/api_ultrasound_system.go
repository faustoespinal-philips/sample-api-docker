/*
 * us-device-api
 *
 * System API for U/S device components
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetIdentity - 
func GetIdentity(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetIdentityCertificate - 
func GetIdentityCertificate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetSystemState - 
func GetSystemState(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetUltrasoundSummary - 
func GetUltrasoundSummary(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
