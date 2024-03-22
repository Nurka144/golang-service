package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/example/test/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
)

type KeycloakResponse struct {
	Sub                string `json:"sub"`
	Email_verified     bool   `json:"email_verified"`
	ProviderId         string `json:"providerId"`
	Name               string `json:"name"`
	Preferred_username string `json:"preferred_username"`
	Given_name         string `json:"given_name"`
	Family_name        string `json:"family_name"`
}

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")

		if token == "" {
			ctx.JSON(http.StatusUnauthorized, &gin.H{"status": "error", "message": "Токен аутентификации отсутствует"})
			ctx.Abort()
			return
		}

		client := resty.New()

		resp, _ := client.R().
			SetHeader("Accept", "application/json").
			SetHeader("Authorization", token).
			Get(viper.GetString("keycloak.url") + "/realms/" + viper.GetString("keycloak.realm") + "/protocol/openid-connect/userinfo")

		statusCode := resp.StatusCode()
		respBodyString := string(resp.Body())
		var responseKeycloakBody KeycloakResponse

		if statusCode != 200 {
			ctx.JSON(statusCode, &gin.H{"status": "error", "message": "Получена ошибка с Keycloak : " + respBodyString})
			ctx.Abort()
			return
		}

		if errResKeycloak := json.Unmarshal(resp.Body(), &responseKeycloakBody); errResKeycloak != nil {
			ctx.JSON(statusCode, &gin.H{"status": "error", "message": "Получена ошибка при вычитывание тела запроса : " + errResKeycloak.Error()})
			ctx.Abort()
			return
		}

		ctx.Set("AuthData", &models.AuthMiddlewareData{Username: responseKeycloakBody.Preferred_username, UserId: -1, AuthType: "Keycloak"})

		ctx.Next()
	}
}
