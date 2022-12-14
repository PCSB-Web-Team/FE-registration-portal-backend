package utils

import (
	// "encoding/json"
	// "fmt"

	"github.com/gin-gonic/gin"
)

func ErrorResponse(err error) gin.H {
	// errJSON, _ := json.Marshal(err)
	// fmt.Println(json.(errJSON))
	return gin.H{"error": err.Error()}
}
