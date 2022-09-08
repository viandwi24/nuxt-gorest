// Package renderer uses template engine to
// render and serve HTML pages
package renderer

// github.com/pilinux/gorest
// The MIT License (MIT)
// Copyright (c) 2022 pilinux

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pilinux/structs"
)

// Render - render response in JSON format
// or use templating engine to serve HTML pages
func Render(c *gin.Context, data interface{}, statusCode int, htmlTpl ...string) {
	if len(htmlTpl) > 0 {
		reqType := c.Request.Header.Get("Accept")

		// Server HTML
		if strings.Contains(reqType, "text/html") {
			c.Set("template", htmlTpl[0])
			c.Set("data", structs.Map(data))
			return
		}
	}

	if statusCode >= 400 {
		c.AbortWithStatusJSON(statusCode, data)
		return
	}

	// Respond with JSON
	c.SecureJSON(statusCode, data)
}
