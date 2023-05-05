package port

import "github.com/gin-gonic/gin"

// CreateRecordEndpoint --
type CreateRecordEndpoint interface {
	CreateRecordExecute(ginContext *gin.Context)
}

// DownloadEndpoint --
type DownloadEndpoint interface {
	DownloadExecute(ginContext *gin.Context)
}
