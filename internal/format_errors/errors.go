package format_errors

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RecordNotFound(c *gin.Context, err error, errMessage ...string) {
	errorMessage := "Record not found"
	if len(errMessage) > 0 {
		errorMessage = errMessage[0]
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"message": errorMessage})
		return
	}

	InternalServerError(c)

}

func InternalServerError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
	return
}
