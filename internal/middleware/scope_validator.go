package middleware

import (
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"

	"github.com/quocanh1897/sample-gin-server/internal/constant"
	apperror "github.com/quocanh1897/sample-gin-server/internal/error"
)

func ScopeValidatorMiddlewareFactory() func([][]constant.Scope) gin.HandlerFunc {
	return middlewareFactoryImpl{}.ValidateScope
}

type middlewareFactoryImpl struct{}

// [[a,b], [c]]
// [[a] || [b]]
func (m middlewareFactoryImpl) ValidateScope(requiredScopes [][]constant.Scope) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		CheckScope(ctx, requiredScopes)
		ctx.Next()
	}
}

func CheckScope(ctx *gin.Context, requiredScopes [][]constant.Scope) {
	sessionScopes, ok := ctx.Value(constant.SessionScopes).([]constant.Scope)
	if !ok {
		ctx.AbortWithError(http.StatusInternalServerError, apperror.NewServerError("failed to get scopes from ctx"))
	}

	validatedSuccess := false
	for _, scopes := range requiredScopes {
		scopeCheck := true
		for _, scope := range scopes {
			scopeCheck = scopeCheck && slices.Contains(sessionScopes, scope)
		}
		validatedSuccess = validatedSuccess || scopeCheck
	}
	if !validatedSuccess {
		ctx.AbortWithError(http.StatusForbidden, apperror.NewForbiddenError("missing required scope"))
	}
}
