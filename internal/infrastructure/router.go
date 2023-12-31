package infrastructure

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const listenPort = ":9999"
const tokenCookieName = "access_token"

type Router struct {
	router *gin.Engine
}

func NewRouter() *Router {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			os.Getenv("WEB_SERVER_ORIGIN"),
			os.Getenv("WEB_SERVER_ORIGIN_LOCAL"),
		},
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
		},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	public := router.Group("/api")

	public.POST("/register", func(c *gin.Context) { authCtrl.Register(c) })
	public.POST("/login", func(c *gin.Context) { authCtrl.GetToken(c, tokenCookieName) })

	protected := router.Group("/api/admin")
	protected.Use(isValidToken())

	protected.POST("/logout", func(c *gin.Context) { authCtrl.UnSetToken(c, tokenCookieName) })
	protected.GET("/user", func(c *gin.Context) {
		token, err := c.Cookie(tokenCookieName)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is required"})
			c.Abort()
			return
		}
		authCtrl.GetUserInfo(c, token)
	})
	protected.POST("/enable", func(c *gin.Context) {
		authCtrl.Enable(c)
	})
	protected.POST("/disable", func(c *gin.Context) {
		authCtrl.Disable(c)
	})

	return &Router{router: router}
}

func (r *Router) Run() {
	err := r.router.Run(listenPort)
	if err != nil {
		panic("Error while running server")
	}
}

func isValidToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie(tokenCookieName)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is required"})
			c.Abort()
			return
		}

		userId, err := tkn.ExtractUserId(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		foundUser, err := orm.FindUserByID(userId)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		if foundUser.Disabled == true {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Target user is disabled"})
			c.Abort()
			return
		}

		c.Next()
	}
}
