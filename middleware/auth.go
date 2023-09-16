package middleware

import (
	"fmt"
	"ginSample/config"
	custom "ginSample/handler/err"
	"ginSample/utils/jwt"
	"github.com/gin-gonic/gin"
)

func Authorization(roles []string, toml config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("#Authorization Middleware# jwt token parsing... start")
		fmt.Println(" # 기존 Access Token : ", ctx.GetHeader("Authorization"))

		claims, res := jwt.TokenParsing(ctx.Request, toml)
		if res != nil {
			ctx.AbortWithStatusJSON(res.StatusCode, res)
			return
		}

		isAuthorized := false
		for _, role := range roles {
			if claims.Role == role {
				isAuthorized = true
				break
			}
		}

		if !isAuthorized {
			ctx.AbortWithStatusJSON(401, custom.NewErrRes(custom.ERR_UNAUTHORIZED))
			return
		}

		ctx.Set("email", claims.Email)
		ctx.Set("name", claims.Name)
		ctx.Set("role", claims.Role)
		ctx.Next()
	}
}
