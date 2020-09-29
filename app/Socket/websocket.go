package Socket

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ServerWs(c *gin.Context) {
	c.HTML(http.StatusOK, "chat/index", gin.H{})
}
