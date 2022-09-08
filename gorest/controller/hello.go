package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/pilinux/gorest/lib/renderer"
)

func Greetings(c *gin.Context) {
	renderer.Render(c, gin.H{"msg": "Hello world!"}, http.StatusOK)
}
