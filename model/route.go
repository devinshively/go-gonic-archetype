package model

import (
    "github.com/gin-gonic/gin"
)

type Route struct {
	Method  string
	Path    string
	Handler func(c *gin.Context)
}
