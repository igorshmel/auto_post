package port

import "github.com/gin-gonic/gin"

// Endpoint _
type Endpoint interface {
	Execute(ginContext *gin.Context)
}
