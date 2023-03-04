package middleware

import (
	"go-shop-api/configs"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserData struct {
	UUID uuid.UUID
}

var identityKey = "uuid"

// the jwt middleware
func JWTMiddleware() (*jwt.GinJWTMiddleware, error) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     1 * time.Minute,
		MaxRefresh:  1 * time.Minute,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*UserData); ok {
				return jwt.MapClaims{
					identityKey: v.UUID,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &UserData{
				UUID: uuid.FromStringOrNil(claims[identityKey].(string)), // converting to uuid
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals Login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			username := loginVals.Username
			password := loginVals.Password

			isAuth, uuid, err := configs.AuthUser(configs.DB, username, password)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"message": err})
				return nil, err
			}

			if isAuth {
				return &UserData{
					UUID: uuid,
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication

		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if _, ok := data.(*UserData); ok {
				return true
			}
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
		SendCookie:    true,
		// SecureCookie:   false, //non HTTPS dev environments
		CookieHTTPOnly: true, // JS can't modify
		CookieDomain:   "localhost:8080",
		// CookieName:     "token", // default jwt
		// TokenLookup:    "cookie:token",
		CookieSameSite: http.SameSiteDefaultMode,
	})

	if err != nil {
		return nil, err
	}

	return authMiddleware, nil
}
