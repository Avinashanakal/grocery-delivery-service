package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const SecretKeyHeader = "X-Secret-Key"

func SecretAuthenticationMiddleware(ctx *gin.Context) {
	// Get the secret key from the request header
	secretKey := ctx.GetHeader(SecretKeyHeader)

	// Check if the secret key is valid
	validSecret := ValidateSecretKey(secretKey)

	if !validSecret {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Continue processing if the secret key is valid
	ctx.Next()
}

// validateSecretKey is a placeholder for your actual secret key validation logic
func ValidateSecretKey(secretKey string) bool {

	return secretKey == "2sdGzJ6r-2sdGzJ6r"
}
