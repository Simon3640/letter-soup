package middlewares

import (
	"context"
	"lettersoup/src/application/modules/auth"
	app_context "lettersoup/src/application/shared/context"
	"lettersoup/src/application/shared/locales"
	"lettersoup/src/domain/models"
	"lettersoup/src/infrastructure/api"
	database "lettersoup/src/infrastructure/database/lettersoup"
	"lettersoup/src/infrastructure/providers"
	"lettersoup/src/infrastructure/repositories"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		uc_result := auth.NewAuthUserUseCase(
			providers.Logger,
			repositories.NewUserRepository(database.DB, providers.Logger),
			providers.JWTProviderInstance,
		).Execute(c, locales.EN_US, token)

		if uc_result.HasError() {
			headers := map[api.HTTPHeaderTypeEnum]string{
				api.CONTENT_TYPE: string(api.APPLICATION_JSON),
			}
			content, statusCode := api.NewRequestResolver[models.UserWithRole]().ResolveDTO(c, uc_result, headers)
			c.JSON(statusCode, content)
			c.Abort()
			return
		}

		user := uc_result.GetData()

		ctx := context.Background()
		ctx = context.WithValue(ctx, app_context.UserKey, *user)

		c.Request = c.Request.WithContext(ctx)

		c.Next()

	}
}
