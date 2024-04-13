package tools

import (
	"github.com/gin-gonic/gin"
)

// Error Response
// type Error struct {
// 	// Error code
// 	Code int

// 	// Error message
// 	Message string
// }

func WriteJSON(c *gin.Context, code int, data interface{}) {
	c.JSON(code, gin.H{
		"data": data,
	})
}

func WriteError(c *gin.Context, code int) {
	c.AbortWithStatus(code)
}

// func WriteError(w http.ResponseWriter, message string, code int) {
// 	resp := Error{
// 		Code:    code,
// 		Message: message,
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(code)

// 	json.NewEncoder(w).Encode(resp)
// }
